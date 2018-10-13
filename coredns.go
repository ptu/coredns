package main

//go:generate go run directives_generate.go

import (
	_ "github.com/ptu/coredns/core/plugin"
	"github.com/ptu/coredns/coremain" // Plug in CoreDNS
)

func main() {
	coremain.Run()
}
