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
	case TypeFloat:
		return "float"
	case TypeString:
		return "string"
	case TypeObject:
		return "object"
	case TypeArray:
		return "array"
	case TypeMaybe:
		return "maybe"
	case TypeStruct:
		return "struct"
	case TypeEnum:
		return "enum"
	case TypeAlias:
		return "alias"
	}
	return ""
}

// Type represents a varlink type. Types are method input and output parameters,
// error output parameters, or custom defined types in the interface description.
type Type struct {
	Kind        TypeKind    `json:"kind"`
	ElementType *Type       `json:"type"`
	Alias       string      `json:"alias"`
	Fields      []TypeField `json:"fields"`
}

// TypeField is a named member of a TypeStruct.
type TypeField struct {
	Name string `json:"name"`
	Type *Type  `json:"type"`
}

// Alias represents a named Type in the interface description.
type Alias struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
	Type *Type  `json:"type"`
}

// Method represents a method defined in the interface description.
type Method struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
	In   *Type  `json:"in"`
	Out  *Type  `json:"out"`
}

// Error represents an error defined in the interface description.
type Error struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
	Type *Type  `json:"type"`
}

// IDL represents a parsed varlink interface description with types, methods, errors and
// documentation.
type IDL struct {
	Name        string        `json:"name"`
	Doc         string        `json:"doc"`
	Description string        `json:"description"`
	Members     []interface{} `json:"members"`
	Aliases     []*Alias      `json:"aliases"`
	Methods     []*Method     `json:"methods"`
	Errors      []*Error      `json:"errors"`
}
