package main

import "crypto/tls"

// FIAP: configuração TLS segura para demonstrar a correção.
var fiapDemoTLSConfig = &tls.Config{
	MinVersion: tls.VersionTLS12,
}
