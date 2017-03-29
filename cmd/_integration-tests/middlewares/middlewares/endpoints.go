package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	svc "github.com/TuneLab/go-truss/cmd/_integration-tests/middlewares/middlewarestest-service/svc"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware applied will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass in the middlewares you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAll(authMiddleware, "Status", "Ping")
	in.WrapAll(addBoolToContext("NotSometimes"), "SometimesWrapped")
	in.WrapAll(addBoolToContext("Always"))

	return in
}

func addBoolToContext(key string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			ctx = context.WithValue(ctx, key, true)
			return next(ctx, request)
		}
	}
}
