package rpchandlers

import (
	"github.com/brics18/bricsd/app/appmessage"
	"github.com/brics18/bricsd/app/rpc/rpccontext"
	"github.com/brics18/bricsd/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
