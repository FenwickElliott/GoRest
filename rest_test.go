package rest_test

import rest "github.com/FenwickElliott/GoRest"
import "testing"

func TestRest(t *testing.T) {
	example := rest.Rest{"http://example.com"}
	resp := example.Get("")

	if resp.StatusCode != 200 {
		t.Error(resp.StatusCode)
	}
}
