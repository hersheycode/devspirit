package dgraph

type Decls struct {
	UID            string          `json:"uid"`
	FuncDecls      []FuncDecl      `json:"funcDecls"`
	StructDecls    []StructDecl    `json:"structDecls"`
	InterfaceDecls []InterfaceDecl `json:"interfaceDecls"`
	VarDecls       []VarDecl       `json:"varDecls"`
	ConstDecls     []ConstDecl     `json:"constDecls"`
	OtherTypeDecls []OtherTypeDecl `json:"otherTypeDecls"`
	CodeString     string          `json:"codeString"`
	DType          []string        `json:"dgraph.type,omitempty"`
}

type FuncDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type StructDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type InterfaceDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type VarDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type ConstDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type OtherTypeDecl struct {
	UID        string   `json:"uid"`
	Ident      string   `json:"ident"`
	CodeString string   `json:"codeString"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type GoFileMetaData struct {
	UID                   string `json:"uid"`
	ImportMetaData        `json:"importMetaData"`
	FuncMetaData          `json:"funcMetaData"`
	StructMetaData        `json:"structMetaData"`
	InterfaceMetaData     `json:"interfaceMetaData"`
	VarDeclMetaData       `json:"varDeclMetaData"`
	ConstDeclMetaData     `json:"constDeclMetaData"`
	OtherTypeDeclMetaData `json:"otherTypeDeclMetaData"`
	DType                 []string `json:"dgraph.type,omitempty"`
}

type ImportMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type FuncMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type StructMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type InterfaceMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type VarDeclMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type ConstDeclMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type OtherTypeDeclMetaData struct {
	UID      string   `json:"uid"`
	Quantity int      `json:"quantity"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type Import struct {
	UID   string   `json:"uid"`
	Alias string   `json:"alias"`
	Path  string   `json:"importPath"`
	DType []string `json:"dgraph.type,omitempty"`
}
