package gml

import "go.uber.org/zap"

// Any takes a key and an arbitrary value and chooses the best way to represent
// them as a field, falling back to a reflection-based approach only if
// necessary.
//
// Since byte/uint8 and rune/int32 are aliases, Any can't differentiate between
// them. To minimize surprises, []byte values are treated as binary blobs, byte
// values are treated as uint8, and runes are always treated as integers.
var Any = zap.Any

// Array constructs a field with the given key and ArrayMarshaler. It provides
// a flexible, but still type-safe and efficient, way to add array-like types
// to the logging context. The struct's MarshalLogArray method is called lazily.
var Array = zap.Array

// Binary constructs a field that carries an opaque binary blob.
//
// Binary data is serialized in an encoding-appropriate format. For example,
// zap's JSON encoder base64-encodes binary blobs. To log UTF-8 encoded text,
// use ByteString.
var Binary = zap.Binary

// Bool constructs a field that carries a bool.
var Bool = zap.Bool

// Bools constructs a field that carries a slice of bools.
var Bools = zap.Bools

// ByteString constructs a field that carries UTF-8 encoded text as a []byte.
// To log opaque binary blobs (which aren't necessarily valid UTF-8), use
// Binary.
var ByteString = zap.ByteString

// ByteStrings constructs a field that carries a slice of []byte, each of which
// must be UTF-8 encoded text.
var ByteStrings = zap.ByteStrings

// Complex128 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex128 to
// interface{}).
var Complex128 = zap.Complex128

// Complex128s constructs a field that carries a slice of complex numbers.
var Complex128s = zap.Complex128s

// Complex64 constructs a field that carries a complex number. Unlike most
// numeric fields, this costs an allocation (to convert the complex64 to
// interface{}).
var Complex64 = zap.Complex64

// Complex64s constructs a field that carries a slice of complex numbers.
var Complex64s = zap.Complex64s

// Duration constructs a field with the given key and value. The encoder
// controls how the duration is serialized.
var Duration = zap.Duration

// Durations constructs a field that carries a slice of time.Durations.
var Durations = zap.Durations

// Error is shorthand for the common idiom NamedError("error", err).
var Error = zap.Error

// Errors constructs a field that carries a slice of errors.
var Errors = zap.Errors

// Float32 constructs a field that carries a float32. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
var Float32 = zap.Float32

// Float32s constructs a field that carries a slice of floats.
var Float32s = zap.Float32s

// Float64 constructs a field that carries a float64. The way the
// floating-point value is represented is encoder-dependent, so marshaling is
// necessarily lazy.
var Float64 = zap.Float64

// Float64s constructs a field that carries a slice of floats.
var Float64s = zap.Float64s

// Int constructs a field with the given key and value.
var Int = zap.Int

// Int16 constructs a field with the given key and value.
var Int16 = zap.Int16

// Int16s constructs a field that carries a slice of integers.
var Int16s = zap.Int16s

// Int32 constructs a field with the given key and value.
var Int32 = zap.Int32

// Int32s constructs a field that carries a slice of integers.
var Int32s = zap.Int32s

// Int64 constructs a field with the given key and value.
var Int64 = zap.Int64

// Int64s constructs a field that carries a slice of integers.
var Int64s = zap.Int64s

// Int8 constructs a field with the given key and value.
var Int8 = zap.Int8

// Int8s constructs a field that carries a slice of integers.
var Int8s = zap.Int8s

// Ints constructs a field that carries a slice of integers.
var Ints = zap.Ints

// NamedError constructs a field that lazily stores err.Error() under the
// provided key. Errors which also implement fmt.Formatter (like those produced
// by github.com/pkg/errors) will also have their verbose representation stored
// under key+"Verbose". If passed a nil error, the field is a no-op.
//
// For the common case in which the key is simply "error", the Error function
// is shorter and less repetitive.
var NamedError = zap.NamedError

// Namespace creates a named, isolated scope within the logger's context. All
// subsequent fields will be added to the new namespace.
//
// This helps prevent key collisions when injecting loggers into sub-components
// or third-party libraries.
var Namespace = zap.Namespace

// Object constructs a field with the given key and ObjectMarshaler. It
// provides a flexible, but still type-safe and efficient, way to add map- or
// struct-like user-defined types to the logging context. The struct's
// MarshalLogObject method is called lazily.
var Object = zap.Object

// Reflect constructs a field with the given key and an arbitrary object. It uses
// an encoding-appropriate, reflection-based function to lazily serialize nearly
// any object into the logging context, but it's relatively slow and
// allocation-heavy. Outside tests, Any is always a better choice.
//
// If encoding fails (e.g., trying to serialize a map[int]string to JSON), Reflect
// includes the error message in the final log output.
var Reflect = zap.Reflect

// Skip constructs a no-op field, which is often useful when handling invalid
// inputs in other Field constructors.
var Skip = zap.Skip

// Stack constructs a field that stores a stacktrace of the current goroutine
// under provided key. Keep in mind that taking a stacktrace is eager and
// expensive (relatively speaking); this function both makes an allocation and
// takes about two microseconds.
var Stack = zap.Stack

// String constructs a field with the given key and value.
var String = zap.String

// Stringer constructs a field with the given key and the output of the value's
// String method. The Stringer's String method is called lazily.
var Stringer = zap.Stringer

// Strings constructs a field that carries a slice of strings.
var Strings = zap.Strings

// Time constructs a zapcore.Field with the given key and value. The encoder
// controls how the time is serialized.
var Time = zap.Time

// Times constructs a field that carries a slice of time.Times.
var Times = zap.Times

// Uint constructs a field with the given key and value.
var Uint = zap.Uint

// Uint16 constructs a field with the given key and value.
var Uint16 = zap.Uint16

// Uint16s constructs a field that carries a slice of unsigned integers.
var Uint16s = zap.Uint16s

// Uint32 constructs a field with the given key and value.
var Uint32 = zap.Uint32

// Uint32s constructs a field that carries a slice of unsigned integers.
var Uint32s = zap.Uint32s

// Uint64 constructs a field with the given key and value.
var Uint64 = zap.Uint64

// Uint64s constructs a field that carries a slice of unsigned integers.
var Uint64s = zap.Uint64s

// Uint8 constructs a field with the given key and value.
var Uint8 = zap.Uint8

// Uint8s constructs a field that carries a slice of unsigned integers.
var Uint8s = zap.Uint8s

// Uintptr constructs a field with the given key and value.
var Uintptr = zap.Uintptr

// Uintptrs constructs a field that carries a slice of pointer addresses.
var Uintptrs = zap.Uintptrs

// Uints constructs a field that carries a slice of unsigned integers.
var Uints = zap.Uints
