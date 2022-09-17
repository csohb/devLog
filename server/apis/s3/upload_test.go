package s3

import (
	"fmt"
	"testing"
)

func TestS3Upload(t *testing.T) {
	s3Manager := NewS3Uploader("yj-devlog", "~/.aws/config", "~/.aws/credentials", "devlog")
	if s3Manager == nil {
		fmt.Println("err")
		return
	}
	fmt.Println(s3Manager)
}
