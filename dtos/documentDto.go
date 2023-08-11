package dtos

type CreateDocumentDto struct {
	Name     string
	Content  string // Raw content
	ParentId uint
}
