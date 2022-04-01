// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/RafayLabs/prompt/pkg/prompt"
)

var proxyOptions = []prompt.Suggest{
	{Text: "--accept-hosts", Description: "Regular expression for hosts that the proxy should accept."},
	{Text: "--accept-paths", Description: "Regular expression for paths that the proxy should accept."},
	{Text: "--address", Description: "The IP address on which to serve on."},
	{Text: "--api-prefix", Description: "Prefix to serve the proxied API under."},
	{Text: "--disable-filter", Description: "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port."},
	{Text: "--keepalive", Description: "keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keepalive."},
	{Text: "-p", Description: "The port on which to run the proxy. Set to 0 to pick a random port."},
	{Text: "--port", Description: "The port on which to run the proxy. Set to 0 to pick a random port."},
	{Text: "--reject-methods", Description: "Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT,PATCH')."},
	{Text: "--reject-paths", Description: "Regular expression for paths that the proxy should reject. Paths specified here will be rejected even accepted by --accept-paths."},
	{Text: "-u", Description: "Unix socket on which to run the proxy."},
	{Text: "--unix-socket", Description: "Unix socket on which to run the proxy."},
	{Text: "-w", Description: "Also serve static files from the given directory under the specified prefix."},
	{Text: "--www", Description: "Also serve static files from the given directory under the specified prefix."},
	{Text: "-P", Description: "Prefix to serve static files under, if static file directory is specified."},
	{Text: "--www-prefix", Description: "Prefix to serve static files under, if static file directory is specified."},
}
