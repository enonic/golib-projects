package queries

var OrganizationsWithProjects = `
{
organizations(count: -1) {
id
name
shortName
displayName
projects(count: -1) {
id
name
description
organizationId
dataVolumeSize
dataVolumeType
flavour
heatParameters
heatTemplate
serviceRepoBranch
serviceRepoUrl
stackId
backup {
enabled
}
}
}
}
`
