// Code generated by thriftrw-plugin-yarpc
// @generated

package helloclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/transport"
	"go.uber.org/yarpc/examples/thrift/hello/echo"
	"go.uber.org/yarpc"
)

// Interface is a client for the Hello service.
type Interface interface {
	Echo(
		ctx context.Context,
		reqMeta yarpc.CallReqMeta,
		Echo *echo.EchoRequest,
	) (*echo.EchoResponse, yarpc.CallResMeta, error)
}

// New builds a new client for the Hello service.
//
// 	client := helloclient.New(dispatcher.Channel("hello"))
func New(c transport.Channel, opts ...thrift.ClientOption) Interface {
	return client{c: thrift.New(thrift.Config{
		Service: "Hello",
		Channel: c,
	}, opts...)}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.Channel) Interface {
		return New(c)
	})
}

type client struct{ c thrift.Client }

func (c client) Echo(
	ctx context.Context,
	reqMeta yarpc.CallReqMeta,
	_Echo *echo.EchoRequest,
) (success *echo.EchoResponse, resMeta yarpc.CallResMeta, err error) {

	args := echo.Hello_Echo_Helper.Args(_Echo)

	var body wire.Value
	body, resMeta, err = c.c.Call(ctx, reqMeta, args)
	if err != nil {
		return
	}

	var result echo.Hello_Echo_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = echo.Hello_Echo_Helper.UnwrapResponse(&result)
	return
}