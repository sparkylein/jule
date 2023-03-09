package ast

import "github.com/julelang/jule/lex"

type NodeData = any // Type of AST Node's data.

// AST Node.
type Node struct {
	Token lex.Token
	Data  any
}

// Group for AST model of comments.
type CommentGroup struct {
	Comments []*Comment
}

// AST model of just comment lines.
type Comment struct {
	Token   lex.Token
	Text string
}

// Directive AST.
type Directive struct {
	Token lex.Token
	Tag   string
}

// Kind type of data types.
type TypeKind interface {
	As_text() string
}

// Type AST.
type Type struct {
	Token lex.Token
	Kind  TypeKind
}

func (t *Type) is_primitive(kind string) bool {
	if t.Kind != nil {
		return false
	}
	return t.Token.Id == lex.ID_DT && t.Token.Kind == kind
}
func (t *Type) IsI8() bool { return t.is_primitive(lex.KND_I8) }
func (t *Type) IsI16() bool { return t.is_primitive(lex.KND_I16) }
func (t *Type) IsI32() bool { return t.is_primitive(lex.KND_I32) }
func (t *Type) IsI64() bool { return t.is_primitive(lex.KND_I64) }
func (t *Type) IsU8() bool { return t.is_primitive(lex.KND_U8) }
func (t *Type) IsU16() bool { return t.is_primitive(lex.KND_U16) }
func (t *Type) IsU32() bool { return t.is_primitive(lex.KND_U32) }
func (t *Type) IsU64() bool { return t.is_primitive(lex.KND_U64) }
func (t *Type) IsF32() bool { return t.is_primitive(lex.KND_F32) }
func (t *Type) IsF64() bool { return t.is_primitive(lex.KND_F64) }
func (t *Type) IsInt() bool { return t.is_primitive(lex.KND_INT) }
func (t *Type) IsUint() bool { return t.is_primitive(lex.KND_UINT) }
func (t *Type) IsUintptr() bool { return t.is_primitive(lex.KND_UINTPTR) }
func (t *Type) IsBool() bool { return t.is_primitive(lex.KND_BOOL) }
func (t *Type) IsStr() bool { return t.is_primitive(lex.KND_STR) }
func (t *Type) IsAny() bool { return t.is_primitive(lex.KND_ANY) }
func (t *Type) IsVoid() bool { return t.Kind == nil && t.Token.Id == lex.ID_NA }

type IdentType struct {
	Ident string
}

type RefType struct {
	Elem *Type
}

type MultiRetType struct {
	Types []*Type
}

// Returns type kind as text.
// Returns empty string kind is nil.
func (t *Type) As_text() string {
	if t.Kind == nil {
		return ""
	}
	return t.Kind.As_text()
}
func (itk *IdentType) As_text() string { return itk.Ident }
func (rtk *RefType) As_text() string {
	if rtk.Elem == nil {
		return ""
	}
	return rtk.Elem.As_text()
}
func (mtk *MultiRetType) As_text() string {
	kind := "("
	i := 0
	for ; i < len(mtk.Types); i++ {
		kind += mtk.Types[i].As_text()
		if i+1 < len(mtk.Types) {
			kind += ","
		}
	}
	kind += ")"
	return kind
}

// Return type AST model.
type RetType struct {
	Kind   *Type
	Idents []lex.Token
}

// Generic type AST.
type Generic struct {
	Token lex.Token
	Ident string
}

// Scope AST.
type Scope struct {
	IsUnsafe   bool
	IsDeferred bool
	Tree       []NodeData
}

// Param AST.
type Param struct {
	Token      lex.Token
	IsMut      bool
	IsVariadic bool
	DataType   *Type
	Ident      string
}

// Function declaration AST.
type FnDecl struct {
	Token       lex.Token
	IsUnsafe    bool
	IsPub       bool
	Ident       string
	Directives  []*Directive
	DocComments *CommentGroup
	Scope       *Scope
	Generics    []*Generic
	RetType     *RetType
	Params      []*Param
}
