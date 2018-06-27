package domain

type Organization struct {
	Name        string
	ShortName   string
	DisplayName string
	Projects    []Project
}
