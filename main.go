package rest

import "fmt"

type Rest struct {
	Endpoint string
}

func (r Rest) Bark() {
	fmt.Println(r.Endpoint)
}
