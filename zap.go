package gzap

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
//
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
// them. To minimize surprises, []byte values are treated as binary blobs, byte
// values are treated as uint8, and runes are always treated as integers.
func Any(key string, value interface{}) zapcore.Field {
	return zap.Any(key, value)
}

// Array constructs a field with the given key and ArrayMarshaler. It provides
// a flexible, but still type-safe and efficient, way to add array-like types
// to the logging context. The struct's MarshalLogArray method is called lazily.
func Array(key string, val zapcore.ArrayMarshaler) zapcore.Field {
	return zap.Array(key, val)
}

// Binary constructs a field that carries an opaque binary blob.
//
// Binary data is serialized in an encoding-appropriate format. For example,
// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
// use ByteString.
func Binary(key string, val []byte) zapcore.Field {
	return zap.Binary(key, val)
}

// Bool constructs a field that carries a bool.
func Bool(key string, val bool) zapcore.Field {
	return zap.Bool(key, val)
}

// Bools constructs a field that carries a slice of bools.
func Bools(key string, bs []bool) zapcore.Field {
	return zap.Bools(key, bs)
}

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
// Binary.
func ByteString(key string, val []byte) zapcore.Field {
	return zap.ByteString(key, val)
}

// ByteStrings constructs a field that carries a slice of []byte, each of which
// must be UTF-8 encoded text.
func ByteStrings(key string, bss [][]byte) zapcore.Field {
	return zap.ByteStrings(key, bss)
}

// Complex128 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex128 to
// interface{}).
func Complex128(key string, val complex128) zapcore.Field {
	return zap.Complex128(key, val)
}

// Complex128s constructs a field that carries a slice of complex numbers.
func Complex128s(key string, nums []complex128) zapcore.Field {
	return zap.Complex128s(key, nums)
}

// Complex64 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex64 to
// interface{}).
func Complex64(key string, val complex64) zapcore.Field {
	return zap.Complex64(key, val)
}

// Complex64s constructs a field that carries a slice of complex numbers.
func Complex64s(key string, nums []complex64) zapcore.Field {
	return zap.Complex64s(key, nums)
}

// Duration constructs a field with the given key and value. The encoder
// controls how the duration is serialized.
func Duration(key string, val time.Duration) zapcore.Field {
	return zap.Duration(key, val)
}

// Durations constructs a field that carries a slice of time.Durations.
func Durations(key string, ds []time.Duration) zapcore.Field {
	return zap.Durations(key, ds)
}

// Error is shorthand for the common idiom NamedError("error", err).
func Error(err error) zapcore.Field {
	return zap.Error(err)
}

// Errors constructs a field that carries a slice of errors.
func Errors(key string, errs []error) zapcore.Field {
	return zap.Errors(key, errs)
}

// Float32 constructs a field that carries a float32. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
func Float32(key string, val float32) zapcore.Field {
	return zap.Float32(key, val)
}

// Float32s constructs a field that carries a slice of floats.
func Float32s(key string, nums []float32) zapcore.Field {
	return zap.Float32s(key, nums)
}

// Float64 constructs a field that carries a float64. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
func Float64(key string, val float64) zapcore.Field {
	return zap.Float64(key, val)
}

// Float64s constructs a field that carries a slice of floats.
func Float64s(key string, nums []float64) zapcore.Field {
	return zap.Float64s(key, nums)
}

// Int constructs a field with the given key and value.
func Int(key string, val int) zapcore.Field {
	return zap.Int(key, val)
}

// Int16 constructs a field with the given key and value.
func Int16(key string, val int16) zapcore.Field {
	return zap.Int16(key, val)
}

// Int16s constructs a field that carries a slice of integers.
func Int16s(key string, nums []int16) zapcore.Field {
	return zap.Int16s(key, nums)
}

// Int32 constructs a field with the given key and value.
func Int32(key string, val int32) zapcore.Field {
	return zap.Int32(key, val)
}

// Int32s constructs a field that carries a slice of integers.
func Int32s(key string, nums []int32) zapcore.Field {
	return zap.Int32s(key, nums)
}

// Int64 constructs a field with the given key and value.
func Int64(key string, val int64) zapcore.Field {
	return zap.Int64(key, val)
}

// Int64s constructs a field that carries a slice of integers.
func Int64s(key string, nums []int64) zapcore.Field {
	return zap.Int64s(key, nums)
}

// Int8 constructs a field with the given key and value.
func Int8(key string, val int8) zapcore.Field {
	return zap.Int8(key, val)
}

// Int8s constructs a field that carries a slice of integers.
func Int8s(key string, nums []int8) zapcore.Field {
	return zap.Int8s(key, nums)
}

// Ints constructs a field that carries a slice of integers.
func Ints(key string, nums []int) zapcore.Field {
	return zap.Ints(key, nums)
}

// NamedError constructs a field that lazily stores err.Error() under the
// provided key. Errors which also implement fmt.Formatter (like those produced
// by github.com/pkg/errors) will also have their verbose representation stored
// under key+"Verbose". If passed a nil error, the field is a no-op.
//
// For the common case in which the key is simply "error", the Error function
// is shorter and less repetitive.
func NamedError(key string, err error) zapcore.Field {
	return zap.NamedError(key, err)
}

// Namespace creates a named, isolated scope within the logger's context. All
// subsequent fields will be added to the new namespace.
//
// This helps prevent key collisions when injecting loggers into sub-components
// or third-party libraries.
func Namespace(key string) zapcore.Field {
	return zap.Namespace(key)
}

// Object constructs a field with the given key and ObjectMarshaler. It
// provides a flexible, but still type-safe and efficient, way to add map- or
// struct-like user-defined types to the logging context. The struct's
// MarshalLogObject method is called lazily.
func Object(key string, val zapcore.ObjectMarshaler) zapcore.Field {
	return zap.Object(key, val)
}

// Reflect constructs a field with the given key and an arbitrary object. It uses
// an encoding-appropriate, reflection-based function to lazily serialize nearly
// any object into the logging context, but it's relatively slow and
// allocation-heavy. Outside tests, Any is always a better choice.
//
// If encoding fails (e.g., trying to serialize a map[int]string to JSON), Reflect
// includes the error message in the final log output.
func Reflect(key string, val interface{}) zapcore.Field {
	return zap.Reflect(key, val)
}

// Skip constructs a no-op field, which is often useful when handling invalid
// inputs in other Field constructors.
func Skip() zapcore.Field {
	return zap.Skip()
}

// Stack constructs a field that stores a stacktrace of the current goroutine
// under provided key. Keep in mind that taking a stacktrace is eager and
// expensive (relatively speaking); this function both makes an allocation and
// takes about two microseconds.
func Stack(key string) zapcore.Field {
	return zap.Stack(key)
}

// String constructs a field with the given key and value.
func String(key string, val string) zapcore.Field {
	return zap.String(key, val)
}

// Stringer constructs a field with the given key and the output of the value's
// String method. The Stringer's String method is called lazily.
func Stringer(key string, val fmt.Stringer) zapcore.Field {
	return zap.Stringer(key, val)
}

// Strings constructs a field that carries a slice of strings.
func Strings(key string, ss []string) zapcore.Field {
	return zap.Strings(key, ss)
}

// Time constructs a zapcore.Field with the given key and value. The encoder
// controls how the time is serialized.
func Time(key string, val time.Time) zapcore.Field {
	return zap.Time(key, val)
}

// Times constructs a field that carries a slice of time.Times.
func Times(key string, ts []time.Time) zapcore.Field {
	return zap.Times(key, ts)
}

// Uint constructs a field with the given key and value.
func Uint(key string, val uint) zapcore.Field {
	return zap.Uint(key, val)
}

// Uint16 constructs a field with the given key and value.
func Uint16(key string, val uint16) zapcore.Field {
	return zap.Uint16(key, val)
}

// Uint16s constructs a field that carries a slice of unsigned integers.
func Uint16s(key string, nums []uint16) zapcore.Field {
	return zap.Uint16s(key, nums)
}

// Uint32 constructs a field with the given key and value.
func Uint32(key string, val uint32) zapcore.Field {
	return zap.Uint32(key, val)
}

// Uint32s constructs a field that carries a slice of unsigned integers.
func Uint32s(key string, nums []uint32) zapcore.Field {
	return zap.Uint32s(key, nums)
}

// Uint64 constructs a field with the given key and value.
func Uint64(key string, val uint64) zapcore.Field {
	return zap.Uint64(key, val)
}

// Uint64s constructs a field that carries a slice of unsigned integers.
func Uint64s(key string, nums []uint64) zapcore.Field {
	return zap.Uint64s(key, nums)
}

// Uint8 constructs a field with the given key and value.
func Uint8(key string, val uint8) zapcore.Field {
	return zap.Uint8(key, val)
}

// Uint8s constructs a field that carries a slice of unsigned integers.
func Uint8s(key string, nums []uint8) zapcore.Field {
	return zap.Uint8s(key, nums)
}

// Uintptr constructs a field with the given key and value.
func Uintptr(key string, val uintptr) zapcore.Field {
	return zap.Uintptr(key, val)
}

// Uintptrs constructs a field that carries a slice of pointer addresses.
func Uintptrs(key string, us []uintptr) zapcore.Field {
	return zap.Uintptrs(key, us)
}

// Uints constructs a field that carries a slice of unsigned integers.
func Uints(key string, us []uint) zapcore.Field {
	return zap.Uints(key, us)
}
