package movies_lib

import (
	"fmt"
	"os"
	"testing"
)

func Test_searchAPI(t *testing.T) {
	os.Setenv("API_KEY", "4ecb0111")
	oa := newOA()
	res, err := oa.ByTitle("blade runner")

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}

	for _, r := range res {
		fmt.Println(r)
	}
}
