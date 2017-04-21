package ext

import "github.com/azenakhi/go-learn/pluggo/plugin"

type DataBasePlugin struct{}

func (DataBasePlugin) Message() string {
	return "DataBasePlugin"
}

func init() {
	plugin.Registry = append(plugin.Registry, &DataBasePlugin{})
}
