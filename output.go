package metadata

type FileInfo struct {
	Name       string
	StructName string
	Fields     []FieldInfo
	IDFields   []FieldInfo
}
type FieldInfo struct {
	Name string
	Type string
}
type Output struct {
	ProjectName string `json:"projectName"`
	RootPath    string `json:"rootPath"`
	Files       []File `json:"files"`
	OutFile     []FileInfo
}
