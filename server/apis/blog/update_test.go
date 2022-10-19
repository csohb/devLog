package blog

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBytesToString(t *testing.T) {
	var err error
	var imgBytes []byte

	imgString := `["",""]`

	if len(imgString) > 0 {
		imgBytes, err = json.Marshal(imgString)
		if err != nil {
			t.Error(err)
		}
	}

	fmt.Println(string(imgBytes))

}
