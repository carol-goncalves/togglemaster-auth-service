package main

import "crypto/tls"

var tlsClientConfig = &tls.Config{
	InsecureSkipVerify: true,
}
