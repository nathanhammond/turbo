package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
)

type AbsoluteSystemPath string
type RelativeSystemPath string
type AbsoluteUnixPath string
type RelativeUnixPath string

type AbsoluteUnixPathInterface interface {
	filePathStamp()
	absolutePathStamp()
	unixPathStamp()

	ToString() string
}

type RelativeUnixPathInterface interface {
	filePathStamp()
	relativePathStamp()
	unixPathStamp()

	ToString() string
}

type AbsoluteSystemPathInterface interface {
	filePathStamp()
	absolutePathStamp()
	systemPathStamp()

	ToString() string
}

type RelativeSystemPathInterface interface {
	filePathStamp()
	relativePathStamp()
	systemPathStamp()

	ToString() string
}

type AbsolutePathInterface interface {
	filePathStamp()
	absolutePathStamp()

	ToString() string
}

type RelativePathInterface interface {
	filePathStamp()
	relativePathStamp()

	ToString() string
}

// UnixPath represents a path whose separators are '/'.
type UnixPathInterface interface {
	filePathStamp()
	unixPathStamp()

	Rel(UnixPathInterface) (RelativeUnixPath, error)
	ToSystemPath() SystemPathInterface
	ToUnixPath() UnixPathInterface
	ToString() string
}

// SystemPath represents a path whose separators are system-dependent.
type SystemPathInterface interface {
	filePathStamp()
	systemPathStamp()

	Rel(SystemPathInterface) (RelativeSystemPath, error)
	ToSystemPath() SystemPathInterface
	ToUnixPath() UnixPathInterface
	ToString() string
}

type FilePathInterface interface {
	filePathStamp()

	ToSystemPath() SystemPathInterface
	ToUnixPath() UnixPathInterface
	ToString() string
}

// For interface reasons, we need a way to distinguish between
// Absolute/Relative/System/Unix/File paths so we stamp them.
func (_ AbsoluteUnixPath) absolutePathStamp()   {}
func (_ AbsoluteSystemPath) absolutePathStamp() {}

func (_ RelativeUnixPath) relativePathStamp()   {}
func (_ RelativeSystemPath) relativePathStamp() {}

func (_ AbsoluteUnixPath) unixPathStamp() {}
func (_ RelativeUnixPath) unixPathStamp() {}

func (_ AbsoluteSystemPath) systemPathStamp() {}
func (_ RelativeSystemPath) systemPathStamp() {}

func (_ AbsoluteUnixPath) filePathStamp()   {}
func (_ RelativeUnixPath) filePathStamp()   {}
func (_ AbsoluteSystemPath) filePathStamp() {}
func (_ RelativeSystemPath) filePathStamp() {}

func StringToUnixPath(path string) UnixPathInterface {
	if filepath.IsAbs(path) {
		return AbsoluteUnixPath(filepath.ToSlash(path))
	} else {
		return RelativeUnixPath(filepath.ToSlash(path))
	}
}

func StringToSystemPath(path string) SystemPathInterface {
	if filepath.IsAbs(path) {
		return AbsoluteSystemPath(filepath.FromSlash(path))
	} else {
		return RelativeSystemPath(filepath.FromSlash(path))
	}
}

func (p AbsoluteUnixPath) ToSystemPath() SystemPathInterface {
	return AbsoluteSystemPath(filepath.FromSlash(p.ToString()))
}

func (p RelativeUnixPath) ToSystemPath() SystemPathInterface {
	return RelativeSystemPath(filepath.FromSlash(p.ToString()))
}

func (p AbsoluteSystemPath) ToSystemPath() SystemPathInterface {
	return p
}

func (p RelativeSystemPath) ToSystemPath() SystemPathInterface {
	return p
}

func (p AbsoluteUnixPath) ToUnixPath() UnixPathInterface {
	return p
}

func (p RelativeUnixPath) ToUnixPath() UnixPathInterface {
	return p
}

func (p AbsoluteSystemPath) ToUnixPath() UnixPathInterface {
	return AbsoluteUnixPath(filepath.ToSlash(p.ToString()))
}

func (p RelativeSystemPath) ToUnixPath() UnixPathInterface {
	return RelativeUnixPath(filepath.ToSlash(p.ToString()))
}

func (p AbsoluteUnixPath) Rel(basePath UnixPathInterface) (RelativeUnixPath, error) {
	processed, err := filepath.Rel(basePath.ToString(), p.ToString())
	return RelativeUnixPath(processed), err
}

func (p RelativeUnixPath) Rel(basePath UnixPathInterface) (RelativeUnixPath, error) {
	processed, err := filepath.Rel(basePath.ToString(), p.ToString())
	return RelativeUnixPath(processed), err
}

func (p AbsoluteSystemPath) Rel(basePath SystemPathInterface) (RelativeSystemPath, error) {
	processed, err := filepath.Rel(basePath.ToString(), p.ToString())
	return RelativeSystemPath(processed), err
}

func (p RelativeSystemPath) Rel(basePath SystemPathInterface) (RelativeSystemPath, error) {
	processed, err := filepath.Rel(basePath.ToString(), p.ToString())
	return RelativeSystemPath(processed), err
}

// ToString returns a string represenation of this Path.
// Used for interfacing with APIs that require a string.
func (p AbsoluteUnixPath) ToString() string {
	return string(p)
}

// ToString returns a string represenation of this Path.
// Used for interfacing with APIs that require a string.
func (p RelativeUnixPath) ToString() string {
	return string(p)
}

// ToString returns a string represenation of this Path.
// Used for interfacing with APIs that require a string.
func (p AbsoluteSystemPath) ToString() string {
	return string(p)
}

// ToString returns a string represenation of this Path.
// Used for interfacing with APIs that require a string.
func (p RelativeSystemPath) ToString() string {
	return string(p)
}

func (p RelativeSystemPath) ToRelativeUnixPath() RelativeUnixPath {
	return p.ToUnixPath().(RelativeUnixPath)
}

// UnsafeToRelativeUnixPath ingests an arbitrary string and treats it as
// a RepoRelativeUnixPath.
func UnsafeToRelativeUnixPath(s string) RelativeUnixPath {
	return RelativeUnixPath(s)
}

// AbsolutePath represents a platform-dependent absolute path on the filesystem,
// and is used to enfore correct path manipulation
type AbsolutePath string

func CheckedToAbsolutePath(s string) (AbsolutePath, error) {
	if filepath.IsAbs(s) {
		return AbsolutePath(s), nil
	}
	return "", fmt.Errorf("%v is not an absolute path", s)
}

// ResolveUnknownPath returns unknown if it is an absolute path, otherwise, it
// assumes unknown is a path relative to the given root.
func ResolveUnknownPath(root AbsolutePath, unknown string) AbsolutePath {
	if filepath.IsAbs(unknown) {
		return AbsolutePath(unknown)
	}
	return root.Join(unknown)
}

func UnsafeToAbsolutePath(s string) AbsolutePath {
	return AbsolutePath(s)
}

func GetCwd() (AbsolutePath, error) {
	cwdRaw, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("invalid working directory: %w", err)
	}
	// We evaluate symlinks here because the package managers
	// we support do the same.
	cwdRaw, err = filepath.EvalSymlinks(cwdRaw)
	if err != nil {
		return "", fmt.Errorf("evaluating symlinks in cwd: %w", err)
	}
	cwd, err := CheckedToAbsolutePath(cwdRaw)
	if err != nil {
		return "", fmt.Errorf("cwd is not an absolute path %v: %v", cwdRaw, err)
	}
	return cwd, nil
}

func (ap AbsolutePath) ToStringDuringMigration() string {
	return ap.asString()
}

func (ap AbsolutePath) Join(args ...string) AbsolutePath {
	return AbsolutePath(filepath.Join(ap.asString(), filepath.Join(args...)))
}
func (ap AbsolutePath) asString() string {
	return string(ap)
}
func (ap AbsolutePath) Dir() AbsolutePath {
	return AbsolutePath(filepath.Dir(ap.asString()))
}

// MkdirAll is the AbsolutePath wrapper for os.MkdirAll
func (ap AbsolutePath) MkdirAll() error {
	return os.MkdirAll(ap.asString(), DirPermissions|0644)
}
func (ap AbsolutePath) Open() (*os.File, error) {
	return os.Open(ap.asString())
}

func (ap AbsolutePath) FileExists() bool {
	return FileExists(ap.asString())
}

// ReadFile reads the contents of the specified file
func (ap AbsolutePath) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(ap.asString())
}

// WriteFile writes the contents of the specified file
func (ap AbsolutePath) WriteFile(contents []byte, mode os.FileMode) error {
	return ioutil.WriteFile(ap.asString(), contents, mode)
}

// EnsureDir ensures that the directory containing this file exists
func (ap AbsolutePath) EnsureDir() error {
	return EnsureDir(ap.asString())
}

// Create is the AbsolutePath wrapper for os.Create
func (ap AbsolutePath) Create() (*os.File, error) {
	return os.Create(ap.asString())
}

// Ext implements filepath.Ext for an absolute path
func (ap AbsolutePath) Ext() string {
	return filepath.Ext(ap.asString())
}

// ToString returns the string representation of this absolute path. Used for
// interfacing with APIs that require a string
func (ap AbsolutePath) ToString() string {
	return ap.asString()
}

// RelativePathString returns the relative path from this AbsolutePath to another absolute path in string form as a string
func (ap AbsolutePath) RelativePathString(path string) (string, error) {
	return filepath.Rel(ap.asString(), path)
}

// EnsureDirFS ensures that the directory containing the given filename is created
func EnsureDirFS(fs afero.Fs, filename AbsolutePath) error {
	dir := filename.Dir()
	err := fs.MkdirAll(dir.asString(), DirPermissions)
	if errors.Is(err, syscall.ENOTDIR) {
		err = fs.Remove(dir.asString())
		if err != nil {
			return errors.Wrapf(err, "removing existing file at %v before creating directories", dir)
		}
		err = fs.MkdirAll(dir.asString(), DirPermissions)
		if err != nil {
			return err
		}
	} else if err != nil {
		return errors.Wrapf(err, "creating directories at %v", dir)
	}
	return nil
}

// WriteFile writes the given bytes to the specified file
func WriteFile(fs afero.Fs, filename AbsolutePath, toWrite []byte, mode os.FileMode) error {
	return afero.WriteFile(fs, filename.asString(), toWrite, mode)
}

// ReadFile reads the contents of the specified file
func ReadFile(fs afero.Fs, filename AbsolutePath) ([]byte, error) {
	return afero.ReadFile(fs, filename.asString())
}

// RemoveFile removes the file at the given path
func RemoveFile(fs afero.Fs, filename AbsolutePath) error {
	return fs.Remove(filename.asString())
}

type pathValue struct {
	base     AbsolutePath
	current  *AbsolutePath
	defValue string
}

func (pv *pathValue) String() string {
	if *pv.current == "" {
		return ResolveUnknownPath(pv.base, pv.defValue).ToString()
	}
	return pv.current.ToString()
}

func (pv *pathValue) Set(value string) error {
	*pv.current = ResolveUnknownPath(pv.base, value)
	return nil
}

func (pv *pathValue) Type() string {
	return "path"
}

var _ pflag.Value = &pathValue{}

// AbsolutePathVar adds a flag interpreted as an absolute path to the given FlagSet.
// It currently requires a root because relative paths are interpreted relative to the
// given root.
func AbsolutePathVar(flags *pflag.FlagSet, target *AbsolutePath, name string, root AbsolutePath, usage string, defValue string) {
	value := &pathValue{
		base:     root,
		current:  target,
		defValue: defValue,
	}
	flags.Var(value, name, usage)
}
