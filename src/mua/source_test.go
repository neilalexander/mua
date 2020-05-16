package mua

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalSource(t *testing.T) {
	src := Source{
		Type: SourceTypeLua,
		Source: SourceCode(`
			print("I'm a new and exciting event")
		`),
	}
	j, err := json.Marshal(src)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

func TestUnmarshalSource(t *testing.T) {
	var src Source
	j := `{"type":0,"source":"CnByaW50KCJIRUxMTyEiKQpwcmludCgiz47PgM6xISIpCg"}`
	err := json.Unmarshal([]byte(j), &src)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", src)
}
