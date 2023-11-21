package entities

type ApplicationType int

const (
	Rest ApplicationType = iota
	GraphQL
)

type ConfigSet struct {
	PackageName            string
	ApplicationCommandName string
	ApplicationType        ApplicationType
}
