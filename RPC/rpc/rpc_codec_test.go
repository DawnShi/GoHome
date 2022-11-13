package rpc

import (
	"encoding/gob"
	"fmt"
	"testing"
)

type testStruct struct {
	Arg1 int
	Arg2 struct {
		Slice []int
	}
}

func Test_encodeDecode(t *testing.T) {
	gob.Register(new(testStruct))
	data := RPCData{
		Func: "SomeFunc",
		Args: []interface{}{
			testStruct{
				Arg1: 123,
				Arg2: struct {
					Slice []int
				}{Slice: []int{1, 2, 3}},
			},
		},
	}

	enc, err := encode(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("enc:", fmt.Sprintf("%#v", enc))

	dec, err := decode(enc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dec:", fmt.Sprintf("%#v", dec))
	t.Log("dec.Args[0]:", fmt.Sprintf("%#v", dec.Args[0]))
	t.Log("data.Args[0]:", fmt.Sprintf("%#v", data.Args[0]))
}
