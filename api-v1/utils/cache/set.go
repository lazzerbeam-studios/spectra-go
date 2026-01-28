package cache

import "github.com/valkey-io/valkey-go"

var ValkeyClient valkey.Client

func SetClient(address string) {
	var err error
	ValkeyClient, err = valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{address},
	})
	if err != nil {
		panic("Failed to open valkey connection")
	}
}
