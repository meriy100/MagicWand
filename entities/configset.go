package entities

type AppType int

const (
	Rest AppType = iota
	GraphQL
)

type ConfigSet struct {
	AppName         string
	AppType         AppType
	RepositoryOwner string
	GCPConfig       GCPConfig
}

type GCPConfig struct {
	ProjectID string
}
