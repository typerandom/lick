package main

type LicenseNode struct {
	Name     string
	FilePath string
	Licenses []*LicenseNode
}

func NewLicenseNode(name string, filePath string) *LicenseNode {
	return &LicenseNode{
		Name:     name,
		FilePath: filePath,
	}
}

type LicenseReader interface {
	Read(filePath string) (*LicenseNode, error)
}
