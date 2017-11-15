package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rest struct {
	Base string
}

func (r Rest) Bark() {
	fmt.Println(r.Base)
}

func (r Rest) Get(endpoint string) {
	fmt.Println("Getting " + r.Base + endpoint)
	url := r.Base + endpoint
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(string(bodyBytes))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
