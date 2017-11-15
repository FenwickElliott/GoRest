package rest

import "fmt"

type Rest struct {
	Base string
}

func (r Rest) Bark() {
	fmt.Println(r.Base)
}

func (r Rest) Get(endpoihnt string) {
	fmt.Println(r.Base + endpoihnt)
}
