// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lstree

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

type readTest struct {
	Name      string
	Input     string
	Output    [][]string
	Positions [][][2]int
	Errors    []error

	// These fields are copied into the Reader
	ReuseRecord bool
}

// In these tests, the § and ∑ characters in readTest.Input are used to denote
// the start of a field and the position of an error respectively.
// They are removed before parsing and are used to verify the position
// information reported by FieldPos.

var readTests = []readTest{
	{
		Name:   "simple",
		Input:  "§100644 §blob §e69de29bb2d1d6434b8b29ae775ad8c2e48c5391\t§package.json\000",
		Output: [][]string{{"100644", "blob", "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", "package.json"}},
	},
	{
		Name:   "no trailing nul",
		Input:  "§100644 §blob §e69de29bb2d1d6434b8b29ae775ad8c2e48c5391\t§package.json",
		Output: [][]string{{"100644", "blob", "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", "package.json"}},
	},
	{
		Name:  "weird file names",
		Input: "§100644 §blob §e69de29bb2d1d6434b8b29ae775ad8c2e48c5391\t§\t\000§100644 §blob §e69de29bb2d1d6434b8b29ae775ad8c2e48c5391\t§\"\000§100644 §blob §5b999efa470b056e329b4c23a73904e0794bdc2f\t§.eslintrc.js\000§100644 §blob §f44f57fff95196c5f7139dfa0b96875f1e9650a9\t§.gitignore\000§100644 §blob §33dbaf21275ca2a5f460249d941cbc27d5da3121\t§README.md\000§040000 §tree §7360f2d292aec95907cebdcbb412a6bf2bd10f8a\t§apps\000§100644 §blob §9ec2879b24ce2c817296eebe2cb3846f8e4751ea\t§package.json\000§040000 §tree §5759aadaea2cde55468a61e7104eb0a9d86c1d30\t§packages\000§100644 §blob §33d0621ee2f4da4a2f6f6bdd51a42618d181e337\t§turbo.json\000",
		Output: [][]string{
			{"100644", "blob", "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", "\t"},
			{"100644", "blob", "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391", "\""},
			{"100644", "blob", "5b999efa470b056e329b4c23a73904e0794bdc2f", ".eslintrc.js"},
			{"100644", "blob", "f44f57fff95196c5f7139dfa0b96875f1e9650a9", ".gitignore"},
			{"100644", "blob", "33dbaf21275ca2a5f460249d941cbc27d5da3121", "README.md"},
			{"040000", "tree", "7360f2d292aec95907cebdcbb412a6bf2bd10f8a", "apps"},
			{"100644", "blob", "9ec2879b24ce2c817296eebe2cb3846f8e4751ea", "package.json"},
			{"040000", "tree", "5759aadaea2cde55468a61e7104eb0a9d86c1d30", "packages"},
			{"100644", "blob", "33d0621ee2f4da4a2f6f6bdd51a42618d181e337", "turbo.json"},
		},
	},
}

func TestRead(t *testing.T) {
	newReader := func(tt readTest) (*Reader, [][][2]int, map[int][2]int) {
		positions, errPositions, input := makePositions(tt.Input)
		r := NewReader(strings.NewReader(input))

		r.ReuseRecord = tt.ReuseRecord
		return r, positions, errPositions
	}

	for _, tt := range readTests {
		t.Run(tt.Name, func(t *testing.T) {
			r, positions, errPositions := newReader(tt)
			out, err := r.ReadAll()
			if wantErr := firstError(tt.Errors, positions, errPositions); wantErr != nil {
				if !reflect.DeepEqual(err, wantErr) {
					t.Fatalf("ReadAll() error mismatch:\ngot  %v (%#v)\nwant %v (%#v)", err, err, wantErr, wantErr)
				}
				if out != nil {
					t.Fatalf("ReadAll() output:\ngot  %q\nwant nil", out)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected Readall() error: %v", err)
				}
				if !reflect.DeepEqual(out, tt.Output) {
					t.Fatalf("ReadAll() output:\ngot  %q\nwant %q", out, tt.Output)
				}
			}

			// Check field and error positions.
			r, _, _ = newReader(tt)
			for recNum := 0; ; recNum++ {
				rec, err := r.Read()
				var wantErr error
				if recNum < len(tt.Errors) && tt.Errors[recNum] != nil {
					wantErr = errorWithPosition(tt.Errors[recNum], recNum, positions, errPositions)
				} else if recNum >= len(tt.Output) {
					wantErr = io.EOF
				}
				if !reflect.DeepEqual(err, wantErr) {
					t.Fatalf("Read() error at record %d:\ngot %v (%#v)\nwant %v (%#v)", recNum, err, err, wantErr, wantErr)
				}
				// ErrFieldCount is explicitly non-fatal.
				if err != nil && !errors.Is(err, ErrFieldCount) {
					if recNum < len(tt.Output) {
						t.Fatalf("need more records; got %d want %d", recNum, len(tt.Output))
					}
					break
				}
				if got, want := rec, tt.Output[recNum]; !reflect.DeepEqual(got, want) {
					t.Errorf("Read vs ReadAll mismatch;\ngot %q\nwant %q", got, want)
				}
				pos := positions[recNum]
				if len(pos) != len(rec) {
					t.Fatalf("mismatched position length at record %d", recNum)
				}
				for i := range rec {
					line, col := r.FieldPos(i)
					if got, want := [2]int{line, col}, pos[i]; got != want {
						t.Errorf("position mismatch at record %d, field %d;\ngot %v\nwant %v", recNum, i, got, want)
					}
				}
			}
		})
	}
}

// firstError returns the first non-nil error in errs,
// with the position adjusted according to the error's
// index inside positions.
func firstError(errs []error, positions [][][2]int, errPositions map[int][2]int) error {
	for i, err := range errs {
		if err != nil {
			return errorWithPosition(err, i, positions, errPositions)
		}
	}
	return nil
}

func errorWithPosition(err error, recNum int, positions [][][2]int, errPositions map[int][2]int) error {
	parseErr, ok := err.(*ParseError)
	if !ok {
		return err
	}
	if recNum >= len(positions) {
		panic(fmt.Errorf("no positions found for error at record %d", recNum))
	}
	errPos, ok := errPositions[recNum]
	if !ok {
		panic(fmt.Errorf("no error position found for error at record %d", recNum))
	}
	parseErr1 := *parseErr
	parseErr1.Line = errPos[0]
	parseErr1.Column = errPos[1]
	return &parseErr1
}

// makePositions returns the expected field positions of all
// the fields in text, the positions of any errors, and the text with the position markers
// removed.
//
// The start of each field is marked with a § symbol;
// CSV lines are separated by ¶ symbols;
// Error positions are marked with ∑ symbols.
func makePositions(text string) ([][][2]int, map[int][2]int, string) {
	buf := make([]byte, 0, len(text))
	var positions [][][2]int
	errPositions := make(map[int][2]int)
	line, col := 1, 1
	recNum := 0

	for len(text) > 0 {
		r, size := utf8.DecodeRuneInString(text)
		switch r {
		case '\000':
			line++
			col = 1
			buf = append(buf, '\000')
			positions = append(positions, [][2]int{})
			recNum++
		case '§':
			if len(positions) == 0 {
				positions = append(positions, [][2]int{})
			}
			positions[len(positions)-1] = append(positions[len(positions)-1], [2]int{line, col})
		case '∑':
			errPositions[recNum] = [2]int{line, col}
		default:
			buf = append(buf, text[:size]...)
			col += size
		}
		text = text[size:]
	}
	return positions, errPositions, string(buf)
}

// nTimes is an io.Reader which yields the string s n times.
type nTimes struct {
	s   string
	n   int
	off int
}

func (r *nTimes) Read(p []byte) (n int, err error) {
	for {
		if r.n <= 0 || r.s == "" {
			return n, io.EOF
		}
		n0 := copy(p, r.s[r.off:])
		p = p[n0:]
		n += n0
		r.off += n0
		if r.off == len(r.s) {
			r.off = 0
			r.n--
		}
		if len(p) == 0 {
			return
		}
	}
}

// benchmarkRead measures reading the provided ls-tree data.
// initReader, if non-nil, modifies the Reader before it's used.
func benchmarkRead(b *testing.B, initReader func(*Reader), rows string) {
	b.ReportAllocs()
	r := NewReader(&nTimes{s: rows, n: b.N})
	if initReader != nil {
		initReader(r)
	}
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			b.Fatal(err)
		}
	}
}

const benchmarkLSTreeData = `100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391		\000100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391	"\000100644 blob 5b999efa470b056e329b4c23a73904e0794bdc2f	.eslintrc.js\000100644 blob f44f57fff95196c5f7139dfa0b96875f1e9650a9	.gitignore\000100644 blob 33dbaf21275ca2a5f460249d941cbc27d5da3121	README.md\000040000 tree 7360f2d292aec95907cebdcbb412a6bf2bd10f8a	apps\000100644 blob 9ec2879b24ce2c817296eebe2cb3846f8e4751ea	package.json\000040000 tree 5759aadaea2cde55468a61e7104eb0a9d86c1d30	packages\000100644 blob 33d0621ee2f4da4a2f6f6bdd51a42618d181e337	turbo.json\000`

func BenchmarkRead(b *testing.B) {
	benchmarkRead(b, nil, benchmarkLSTreeData)
}

func BenchmarkReadReuseRecord(b *testing.B) {
	benchmarkRead(b, func(r *Reader) { r.ReuseRecord = true }, benchmarkLSTreeData)
}
