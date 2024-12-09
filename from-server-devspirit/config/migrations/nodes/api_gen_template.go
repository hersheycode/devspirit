package schema

type StringOnlyFile struct {
	UID      string   `json:"uid"`
	FilePath string   `json:"filePath"`
	Content  string   `json:"stringContent"`
	DType    []string `json:"dgraph.type,omitempty"`
}

type Agent struct {
	UID       string   `json:"uid"`
	Email     string   `json:"email"`
	LoginName string   `json:"loginName"`
	Password  string   `json:"password"`
	Username  string   `json:"username"`
	DType     []string `json:"dgraph.type,omitempty"`
}

type VM struct {
	UID     string `json:"uid"`
	RepoUID string `json:"repoUid"`
	Name    string `json:"name"`
	Agent   `json:"repoAgent"`
	DType   []string `json:"dgraph.type,omitempty"`
}
