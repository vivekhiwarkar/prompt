// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/RafayLabs/prompt/pkg/prompt"
)

var topNodeOptions = []prompt.Suggest{
	{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	{Text: "--heapster-port", Description: "Port name in service to use"},
	{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	{Text: "--heapster-service", Description: "Name of Heapster service"},
	{Text: "--no-headers", Description: "If present, print output without headers"},
	{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	{Text: "--sort-by", Description: "If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'."},
}
