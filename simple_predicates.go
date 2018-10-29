package typep

import (
	"go/types"
)

// TODO(Quasilyte): generate simple predicates for all types.

// Simple 1-to-1 type predicates via type assertion.

// IsPointer reports whether a given type has *types.Pointer type.
func IsPointer(typ types.Type) bool {
	_, ok := typ.(*types.Pointer)
	return ok
}

// IsArray reports whether a given type has *types.Array type.
func IsArray(typ types.Type) bool {
	_, ok := typ.(*types.Array)
	return ok
}

// *types.Basic predicates for the info field.

// IsBasicFloat reports whether typ is a *types.Basic marked with IsFloat flag.
func IsBasicFloat(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Info()&types.IsFloat != 0
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

// *types.Basic predicates for the kind field.

// HasBoolKind reports whether typ is a *types.Basic with its kind set to types.Bool.
func HasBoolKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Bool
	}
	return false
}

// HasBoolKind reports whether typ is a *types.Basic with its kind set to types.Int.
func HasIntKind(typ types.Type) bool {
	if typ, ok := typ.(*types.Basic); ok {
		return typ.Kind() == types.Int
	}
	return false
}
