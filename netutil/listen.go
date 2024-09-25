package netutil

import (
	"errors"
	"fmt"
	"net"
	"syscall"
)

// TryListen tries to listen to the specified port.
// If the specified port is already in use, TryListen tries to listen to a new port.
func TryListen(network string, port int) (net.Listener, error) {
	ln, err := net.Listen(network, fmt.Sprintf(":%d", port))
	if err != nil {
		if errors.Is(err, syscall.EADDRINUSE) {
			if port == 0 {
				return nil, errors.New("failed to listen to new port")
			}
			return TryListen(network, 0)
		}
		return nil, err
	}
	return ln, nil
}
