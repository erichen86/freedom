package project

func init() {
	content["/domain/dto/dto.go"] = dtoTemplate()
}

func dtoTemplate() string {
	return `
	//Package dto generated by 'freedom new-project {{.PackagePath}}'
	package dto
	`
}
