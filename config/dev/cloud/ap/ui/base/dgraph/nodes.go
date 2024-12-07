package dgraph

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
)

//TODO add methods
type Schema interface {
}

func (i Import) fmtNode(node *ast.ImportSpec) *Import {
	alias := ""
	namedImp := node.Name
	if namedImp == nil {
		alias = ""
	} else {
		alias = namedImp.Name
	}
	return &Import{
		Alias: alias,
		Path:  node.Path.Value,
		DType: []string{"Import"},
	}
}

func (i FuncDecl) fmtNode(node *ast.FuncDecl) *FuncDecl {
	str, err := fmtNode(node)
	if err != nil {
		panic(err)
	}
	return &FuncDecl{
		Ident:      node.Name.Name,
		CodeString: str,
		DType:      []string{"FuncDecl"},
	}
}

func (i InterfaceDecl) fmtNode(node *ast.TypeSpec) *InterfaceDecl {
	str, err := fmtNode(node)
	if err != nil {
		panic(err)
	}
	return &InterfaceDecl{
		Ident:      node.Name.Name,
		CodeString: str,
		DType:      []string{"InterfaceDecl"},
	}
}

func (i StructDecl) fmtNode(node *ast.TypeSpec) *StructDecl {
	str, err := fmtNode(node)
	if err != nil {
		panic(err)
	}
	return &StructDecl{
		Ident:      node.Name.Name,
		CodeString: str,
		DType:      []string{"StructDecl"},
	}
}

func (i OtherTypeDecl) fmtNode(node *ast.TypeSpec) *OtherTypeDecl {
	str, err := fmtNode(node)
	if err != nil {
		panic(err)
	}
	return &OtherTypeDecl{
		Ident:      node.Name.Name,
		CodeString: str,
		DType:      []string{"OtherTypeDecl"},
	}
}

func fmtNode(node ast.Node) (string, error) {
	buf := &bytes.Buffer{}
	if err := format.Node(buf, token.NewFileSet(), node); err != nil {
		return "", err
	}
	return buf.String(), nil
}
