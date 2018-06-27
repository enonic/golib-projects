package domain

type Backup struct {
	Enabled bool
}

type Project struct {
	Name              string
	ShortName         string
	Description       string
	Id                string
	OrganizationId    string
	DataVolumeSize    string
	DataVolumeType    string
	Flavour           string
	HeatParameters    string
	HeatTemplate      string
	ServiceRepoBranch string
	ServiceRepoUrl    string
	StackId           string
	Backup            Backup
}
