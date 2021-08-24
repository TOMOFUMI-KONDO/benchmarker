package main

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

const (
	baseDir       = "proto/"
	protoTemplate = `syntax = "proto3";

option go_package = "<packageBase>/<serviceLower>";

package <serviceLower>;

service <service>{
  // TODO: implement service
}
`
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("service must be specified")
	}
	service := os.Args[1]
	if service == "" {
		log.Fatalf("service can't be empty")
	}
	serviceLower := strings.ToLower(service)

	dir := baseDir + serviceLower
	if err := os.MkdirAll(dir, fs.FileMode(0766)); err != nil {
		log.Fatalf("could not mkdir: %v", err)
	}

	path := baseDir + serviceLower + "/" + serviceLower + ".proto"
	content := strings.ReplaceAll(protoTemplate, "<packageBase>", "example.com/benchmarker/proto")
	content = strings.ReplaceAll(content, "<serviceLower>", serviceLower)
	content = strings.ReplaceAll(content, "<service>", service)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatalf("could not generate file: %v", err)
	}
}
