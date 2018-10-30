// Code generated by simplePredicates_generate.go; DO NOT EDIT

// Package typep gives predicates for checking go types
package typep

import (
	"go/types"
)

// Simple 1-to-1 type predicates via type assertion.

// IsBasic reports whether a given type has *types.Basic type.
func IsBasic(typ types.Type) bool {
	_, ok := typ.(*types.Basic)
	return ok
}

// IsArray reports whether a given type has *types.Array type.
func IsArray(typ types.Type) bool {
	_, ok := typ.(*types.Array)
	return ok
}

// IsSlice reports whether a given type has *types.Slice type.
func IsSlice(typ types.Type) bool {
	_, ok := typ.(*types.Slice)
	return ok
}

// IsStruct reports whether a given type has *types.Struct type.
func IsStruct(typ types.Type) bool {
	_, ok := typ.(*types.Struct)
	return ok
}

// IsPointer reports whether a given type has *types.Pointer type.
func IsPointer(typ types.Type) bool {
	_, ok := typ.(*types.Pointer)
	return ok
}

// IsTuple reports whether a given type has *types.Tuple type.
func IsTuple(typ types.Type) bool {
	_, ok := typ.(*types.Tuple)
	return ok
}

// IsSignature reports whether a given type has *types.Signature type.
func IsSignature(typ types.Type) bool {
	_, ok := typ.(*types.Signature)
	return ok
}

// IsInterface reports whether a given type has *types.Interface type.
func IsInterface(typ types.Type) bool {
	_, ok := typ.(*types.Interface)
	return ok
}

// IsMap reports whether a given type has *types.Map type.
func IsMap(typ types.Type) bool {
	_, ok := typ.(*types.Map)
	return ok
}

// IsChan reports whether a given type has *types.Chan type.
func IsChan(typ types.Type) bool {
	_, ok := typ.(*types.Chan)
	return ok
}

// IsNamed reports whether a given type has *types.Named type.
func IsNamed(typ types.Type) bool {
	_, ok := typ.(*types.Named)
	return ok
}

// *types.Basic predicates for the info field.

// IsBasicBoolean reports whether typ is a *types.Basic marked with IsBoolean flag.
func IsBasicBoolean(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsBoolean != 0
	}
	return false
}

// IsBasicInteger reports whether typ is a *types.Basic marked with IsInteger flag.
func IsBasicInteger(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsInteger != 0
	}
	return false
}

// IsBasicUnsigned reports whether typ is a *types.Basic marked with IsUnsigned flag.
func IsBasicUnsigned(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsUnsigned != 0
	}
	return false
}

// IsBasicFloat reports whether typ is a *types.Basic marked with IsFloat flag.
func IsBasicFloat(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsFloat != 0
	}
	return false
}

// IsBasicComplex reports whether typ is a *types.Basic marked with IsComplex flag.
func IsBasicComplex(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsComplex != 0
	}
	return false
}

// IsBasicString reports whether typ is a *types.Basic marked with IsString flag.
func IsBasicString(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsString != 0
	}
	return false
}

// IsBasicUntyped reports whether typ is a *types.Basic marked with IsUntyped flag.
func IsBasicUntyped(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsUntyped != 0
	}
	return false
}

// IsBasicOrdered reports whether typ is a *types.Basic marked with IsOrdered flag.
func IsBasicOrdered(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsOrdered != 0
	}
	return false
}

// IsBasicNumeric reports whether typ is a *types.Basic marked with IsNumeric flag.
func IsBasicNumeric(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsNumeric != 0
	}
	return false
}

// IsBasicConstType reports whether typ is a *types.Basic marked with IsConstType flag.
func IsBasicConstType(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsConstType != 0
	}
	return false
}

// *types.Basic predicates for the kind field.

// HasBoolKind reports whether typ is a *types.Basic with its kind set to types.Bool.
func HasBoolKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Bool
	}
	return false
}

// HasIntKind reports whether typ is a *types.Basic with its kind set to types.Int.
func HasIntKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int
	}
	return false
}

// HasInt8Kind reports whether typ is a *types.Basic with its kind set to types.Int8.
func HasInt8Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int8
	}
	return false
}

// HasInt16Kind reports whether typ is a *types.Basic with its kind set to types.Int16.
func HasInt16Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int16
	}
	return false
}

// HasInt32Kind reports whether typ is a *types.Basic with its kind set to types.Int32.
func HasInt32Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int32
	}
	return false
}

// HasInt64Kind reports whether typ is a *types.Basic with its kind set to types.Int64.
func HasInt64Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int64
	}
	return false
}

// HasUintKind reports whether typ is a *types.Basic with its kind set to types.Uint.
func HasUintKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uint
	}
	return false
}

// HasUint8Kind reports whether typ is a *types.Basic with its kind set to types.Uint8.
func HasUint8Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uint8
	}
	return false
}

// HasUint16Kind reports whether typ is a *types.Basic with its kind set to types.Uint16.
func HasUint16Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uint16
	}
	return false
}

// HasUint32Kind reports whether typ is a *types.Basic with its kind set to types.Uint32.
func HasUint32Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uint32
	}
	return false
}

// HasUint64Kind reports whether typ is a *types.Basic with its kind set to types.Uint64.
func HasUint64Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uint64
	}
	return false
}

// HasUintptrKind reports whether typ is a *types.Basic with its kind set to types.Uintptr.
func HasUintptrKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Uintptr
	}
	return false
}

// HasFloat32Kind reports whether typ is a *types.Basic with its kind set to types.Float32.
func HasFloat32Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Float32
	}
	return false
}

// HasFloat64Kind reports whether typ is a *types.Basic with its kind set to types.Float64.
func HasFloat64Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Float64
	}
	return false
}

// HasComplex64Kind reports whether typ is a *types.Basic with its kind set to types.Complex64.
func HasComplex64Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Complex64
	}
	return false
}

// HasComplex128Kind reports whether typ is a *types.Basic with its kind set to types.Complex128.
func HasComplex128Kind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Complex128
	}
	return false
}

// HasStringKind reports whether typ is a *types.Basic with its kind set to types.String.
func HasStringKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.String
	}
	return false
}

// HasUnsafePointerKind reports whether typ is a *types.Basic with its kind set to types.UnsafePointer.
func HasUnsafePointerKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UnsafePointer
	}
	return false
}

// HasUntypedBoolKind reports whether typ is a *types.Basic with its kind set to types.UntypedBool.
func HasUntypedBoolKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedBool
	}
	return false
}

// HasUntypedIntKind reports whether typ is a *types.Basic with its kind set to types.UntypedInt.
func HasUntypedIntKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedInt
	}
	return false
}

// HasUntypedRuneKind reports whether typ is a *types.Basic with its kind set to types.UntypedRune.
func HasUntypedRuneKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedRune
	}
	return false
}

// HasUntypedFloatKind reports whether typ is a *types.Basic with its kind set to types.UntypedFloat.
func HasUntypedFloatKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedFloat
	}
	return false
}

// HasUntypedComplexKind reports whether typ is a *types.Basic with its kind set to types.UntypedComplex.
func HasUntypedComplexKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedComplex
	}
	return false
}

// HasUntypedStringKind reports whether typ is a *types.Basic with its kind set to types.UntypedString.
func HasUntypedStringKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedString
	}
	return false
}

// HasUntypedNilKind reports whether typ is a *types.Basic with its kind set to types.UntypedNil.
func HasUntypedNilKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.UntypedNil
	}
	return false
}
