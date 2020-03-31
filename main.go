package main

import (
	r "github.com/guimesmo/dextraflix-go/routers"
)

func main() {
	r.App.Start(":3000")
}
