package mtransport

import (
	"net"

	transport "go-micro.dev/v4/transport"
)

type netListener struct{}

// getNetListener Get net.Listener from ListenOptions
func getNetListener(o *transport.ListenOptions) net.Listener {
	if o.Context == nil {
		return nil
	}

	if l, ok := o.Context.Value(netListener{}).(net.Listener); ok && l != nil {
		return l
	}

	return nil
}
