// Package lstree reads the output of calls to `git ls-tree`.
package lstree

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// Separators that appear in the output of `git ls-tree`
const space rune = ' '
const tab rune = '\t'
const nul rune = '\000'

// A ParseError is returned for parsing errors.
// Lines and columns are both 1-indexed.
type ParseError struct {
	Line   int   // Line where the error occurred
	Column int   // Column where the error occurred
	Err    error // The actual error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error on line %d, column %d: %v", e.Line, e.Column, e.Err)
}

func (e *ParseError) Unwrap() error { return e.Err }

// These are the errors that can be returned in ParseError.Err.
var (
	ErrInvalidObjectMode = errors.New("object mode is not valid")
	ErrInvalidObjectType = errors.New("object type is not valid")
	ErrInvalidObjectName = errors.New("object name is not valid")
	ErrInvalidPath       = errors.New("path is not valid")
	ErrFieldCount        = errors.New("too many fields")
)

// A Reader reads records from `git ls-tree`'s output`. The Reader converts
// all \r\n sequences in its input to plain \n.
type Reader struct {
	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	ReuseRecord bool

	reader *bufio.Reader

	// numLine is the current line being read in the CSV file.
	numLine int

	// rawBuffer is a line buffer only used by the readLine method.
	rawBuffer []byte

	// recordBuffer holds the unescaped fields, one after another.
	// The fields can be accessed by using the indexes in fieldIndexes.
	// E.g., For the row `a,"b","c""d",e`, recordBuffer will contain `abc"de`
	// and fieldIndexes will contain the indexes [1, 2, 5, 6].
	recordBuffer []byte

	// fieldIndexes is an index of fields inside recordBuffer.
	// The i'th field ends at offset fieldIndexes[i] in recordBuffer.
	fieldIndexes []int

	// fieldPositions is an index of field positions for the
	// last record returned by Read.
	fieldPositions []position

	// lastRecord is a record cache and only used when ReuseRecord == true.
	lastRecord []string
}

// NewReader returns a new Reader that reads from r.
func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader: bufio.NewReader(reader),
	}
}

// Read reads one record from `reader`.
// Read always returns either a non-nil record or a non-nil error,
// but not both.
//
// If there is no data left to be read, Read returns nil, io.EOF.
//
// If ReuseRecord is true, the returned slice may be shared
// between multiple calls to Read.
func (r *Reader) Read() (record []string, err error) {
	if r.ReuseRecord {
		record, err = r.readRecord(r.lastRecord)
		r.lastRecord = record
	} else {
		record, err = r.readRecord(nil)
	}
	return record, err
}

// FieldPos returns the line and column corresponding to
// the start of the field with the given index in the slice most recently
// returned by Read. Numbering of lines and columns starts at 1;
// columns are counted in bytes, not runes.
//
// If this is called with an out-of-bounds index, it panics.
func (r *Reader) FieldPos(field int) (line, column int) {
	if field < 0 || field >= len(r.fieldPositions) {
		panic("out of range index passed to FieldPos")
	}
	p := &r.fieldPositions[field]
	return p.line, p.col
}

// pos holds the position of a field in the current line.
type position struct {
	line, col int
}

// ReadAll reads all the records from reader until EOF.
//
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func (r *Reader) ReadAll() (records [][]string, err error) {
	for {
		record, err := r.readRecord(nil)
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}

// readEntry reads the next entry (with the trailing NUL).
// If EOF is hit without a trailing NUL, it will be omitted.
// If some bytes were read then the error is never io.EOF.
// The result is only valid until the next call to readEntry.
func (r *Reader) readEntry() ([]byte, error) {
	entry, err := r.reader.ReadSlice('\000')
	if err == bufio.ErrBufferFull {
		r.rawBuffer = append(r.rawBuffer[:0], entry...)
		for err == bufio.ErrBufferFull {
			entry, err = r.reader.ReadSlice('\000')
			r.rawBuffer = append(r.rawBuffer, entry...)
		}
		entry = r.rawBuffer
	}
	if len(entry) > 0 && err == io.EOF {
		entry = append(entry, '\000')
		err = nil
	}
	r.numLine++

	return entry, err
}

func checkValid(fieldNumber int, value *[]byte) error {
	switch fieldNumber {
	case 0:
		return checkObjectMode(value)
	case 1:
		return checkObjectType(value)
	case 2:
		return checkObjectName(value)
	case 3:
		return nil
	default:
		return ErrFieldCount
	}
}

func checkObjectMode(value *[]byte) error {
	if len(*value) != 6 {
		return ErrInvalidObjectMode
	}

	// 0-7 are 0x30 - 0x37
	for _, currentByte := range *value {
		if (currentByte ^ 0x30) > 7 {
			return ErrInvalidObjectMode
		}
	}

	// length of 6, 0-7
	return nil
}

var allowed = []byte("blob tree commit ")

func checkObjectType(value *[]byte) error {
	// Because of the space separator, there is no way to pass in a space.
	index := bytes.Index(allowed, *value)
	if index != -1 && allowed[index+len(*value)] != byte(space) {
		return ErrInvalidObjectType
	}
	return nil
}

func checkObjectName(value *[]byte) error {
	if len(*value) != 40 {
		return ErrInvalidObjectName
	}

	// 0-9 are 0x30 - 0x39
	// a-f are 0x61 - 0x66
	for _, currentByte := range *value {
		isNumber := (currentByte ^ 0x30) < 10
		isAlpha := (currentByte ^ 0x60 - 1) < 6
		if !(isNumber || isAlpha) {
			return ErrInvalidObjectName
		}
	}

	// length of 40, hex
	return nil
}

func checkPath(value *[]byte) error {
	// Exists at all.
	if len(*value) == 0 {
		return ErrInvalidPath
	}
	return nil
}

var separators = []rune{space, space, tab, nul}

// readRecord reads a single record.
func (r *Reader) readRecord(dst []string) ([]string, error) {
	entry, errRead := r.readEntry()
	if errRead == io.EOF {
		return nil, errRead
	}

	// Parse each field in the record.
	r.recordBuffer = r.recordBuffer[:0]
	r.fieldIndexes = r.fieldIndexes[:0]
	r.fieldPositions = r.fieldPositions[:0]
	pos := position{line: r.numLine, col: 1}

	for fieldNumber, separator := range separators {
		length := bytes.IndexRune(entry, separator)
		field := entry[:length]

		fieldError := checkValid(fieldNumber, &field)
		if fieldError != nil {
			return nil, fieldError
		}

		offset := length + 1
		entry = entry[offset:]
		r.recordBuffer = append(r.recordBuffer, field...)
		r.fieldIndexes = append(r.fieldIndexes, len(r.recordBuffer))
		r.fieldPositions = append(r.fieldPositions, pos)
		pos.col += offset
	}

	// Create a single string and create slices out of it.
	// This pins the memory of the fields together, but allocates once.
	str := string(r.recordBuffer) // Convert to string once to batch allocations
	dst = dst[:0]
	if cap(dst) < len(r.fieldIndexes) {
		dst = make([]string, len(r.fieldIndexes))
	}
	dst = dst[:len(r.fieldIndexes)]
	var preIdx int
	for i, idx := range r.fieldIndexes {
		dst[i] = str[preIdx:idx]
		preIdx = idx
	}

	return dst, nil
}
