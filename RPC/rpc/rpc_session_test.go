package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSession_ReadWrite(t *testing.T) {
	addr := ":8000"
	data := "hello, world!"

	wg := sync.WaitGroup{}

	wg.Add(2)

	// 写数据
	go func() {
		defer wg.Done()

		listener, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}

		conn, _ := listener.Accept()
		s := Session{conn: conn}

		err = s.Write([]byte(data))
		if err != nil {
			t.Fatal(err)
		}
	}()

	// 读数据
	go func() {
		defer wg.Done()

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := Session{conn: conn}

		gotData, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		gotString := string(gotData)
		if gotString != data {
			t.Fatal(fmt.Sprintf("got(%s) != data(%s)", gotString, data))
		}
	}()

	wg.Wait()
}
