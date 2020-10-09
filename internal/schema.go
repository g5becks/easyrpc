package internal

const (
	TypeBool = iota
	TypeInt
	TypeFloat
	TypeString
	TypeObject
	TypeArray
	TypeMaybe
	TypeMap
	TypeStruct
	TypeEnum
	TypeAlias
)

// TypeKind specifies the type of an Type.
type TypeKind uint

func (t TypeKind) String() string {
	switch t {
	case TypeBool:
		return "bool"
	case TypeInt:
		return "int"
	cas
	}

}

// Type represents a varlink type. Types are method input and output parameters,
// error output parameters, or custom defined types in the interface description.
type Type struct {
	Kind        TypeKind
	ElementType *Type
	Alias       string
	Fields      []TypeField
}

// TypeField is a named member of a TypeStruct.
type TypeField struct {
	Name string
	Type *Type
}

// Alias represents a named Type in the interface description.
type Alias struct {
	Name string
	Doc  string
	Type *Type
}

// Method represents a method defined in the interface description.
type Method struct {
	Name string
	Doc  string
	In   *Type
	Out  *Type
}

// Error represents an error defined in the interface description.
type Error struct {
	Name string
	Doc  string
	Type *Type
}

// IDL represents a parsed varlink interface description with types, methods, errors and
// documentation.
type IDL struct {
	Name        string
	Doc         string
	Description string
	Members     []interface{}
	Aliases     []*Alias
	Methods     []*Method
	Errors      []*Error
}
