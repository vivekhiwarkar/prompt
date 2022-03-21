// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/RafaySystems/prompt/pkg/prompt"
)

var topPodOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-A", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--containers", Description: "If present, print usage of containers within a pod."},
	prompt.Suggest{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	prompt.Suggest{Text: "--heapster-port", Description: "Port name in service to use"},
	prompt.Suggest{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	prompt.Suggest{Text: "--heapster-service", Description: "Name of Heapster service"},
	prompt.Suggest{Text: "--no-headers", Description: "If present, print output without headers."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort pods list using specified field. The field can be either 'cpu' or 'memory'."},
}