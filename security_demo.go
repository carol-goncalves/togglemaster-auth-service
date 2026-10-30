package main

import "crypto/tls"

// FIAP: demonstração de vulnerabilidade para o scanner SAST.
// Não utilizado pela aplicação.
var fiapDemoTLSConfig = &tls.Config{
	InsecureSkipVerify: true,
}
