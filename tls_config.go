package main

import "crypto/tls"

var tlsClientConfig = &tls.Config{
	MinVersion: tls.VersionTLS12,
}
