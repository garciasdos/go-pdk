/*
Package Kong/go-pdk implements Kong's Plugin Development Kit for Go.

It directly parallels the existing kong PDK for Lua plugins.

Kong plugins written in Go implement event handlers as methods on the Plugin's
structure, with the given signature:

	func (conf *MyConfig) Access (kong *pdk.PDK) {
		...
	}

The `kong` argument of type `*pdk.PDK` is the entrypoint for all PDK functions.
For example, to get the client's IP address, you'd use `kong.Client.GetIp()`.
*/
package pdk

import (
	"net"

	"github.com/postmanlabs/go-pdk/bridge"
	"github.com/postmanlabs/go-pdk/client"
	"github.com/postmanlabs/go-pdk/ctx"
	"github.com/postmanlabs/go-pdk/ip"
	"github.com/postmanlabs/go-pdk/log"
	"github.com/postmanlabs/go-pdk/nginx"
	"github.com/postmanlabs/go-pdk/node"
	"github.com/postmanlabs/go-pdk/request"
	"github.com/postmanlabs/go-pdk/response"
	"github.com/postmanlabs/go-pdk/router"
	"github.com/postmanlabs/go-pdk/service"
	service_request "github.com/postmanlabs/go-pdk/service/request"
	service_response "github.com/postmanlabs/go-pdk/service/response"
)

// PDK go pdk module
type PDK struct {
	Client          client.Client
	Ctx             ctx.Ctx
	Log             log.Log
	Nginx           nginx.Nginx
	Request         request.Request
	Response        response.Response
	Router          router.Router
	IP              ip.Ip
	Node            node.Node
	Service         service.Service
	ServiceRequest  service_request.Request
	ServiceResponse service_response.Response
}

// Init initialize go pdk.  Called by the pluginserver at initialization.
func Init(conn net.Conn) *PDK {
	b := bridge.New(conn)
	return &PDK{
		Client:          client.Client{b},
		Ctx:             ctx.Ctx{b},
		Log:             log.Log{b},
		Nginx:           nginx.Nginx{b},
		Request:         request.Request{b},
		Response:        response.Response{b},
		Router:          router.Router{b},
		IP:              ip.Ip{b},
		Node:            node.Node{b},
		Service:         service.Service{b},
		ServiceRequest:  service_request.Request{b},
		ServiceResponse: service_response.Response{b},
	}
}
