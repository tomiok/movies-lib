package movies_lib

import (
	"fmt"
	"testing"
)

func Test_searchAPI(t *testing.T) {
	oa := newOA()
	_, err := oa.Search("blade runner")

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
