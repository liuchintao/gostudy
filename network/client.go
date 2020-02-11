package network

import (
	"fmt"
	"net/http"
)

type config struct {
	Transport http.RoundTripper
}

type client struct {
	transport http.RoundTripper
}

func newClient(cfg config) client {
	c := client{}
	httpTransport, ok := cfg.Transport.(*http.Transport)
	if !ok {
		fmt.Println("Type transfer failed.")
	} else {
		fmt.Println("Type Transfer successfully.")
	}
	c.transport = httpTransport
	return c
}
