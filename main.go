package main

import (
	"github.com/rastasi/go-crud/domain"
	"github.com/rastasi/go-crud/http"
)

func main() {
	domain := domain.NewDomain()
	http.StartHttpServer(domain)
}
