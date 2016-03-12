package winio

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkPipeIo(b *testing.B) {
	name := fmt.Sprintf(`\\.\pipe\testing_%d`, rand.Int63())
	l, err := ListenPipe(name, "")
	if err != nil {
		b.Fatal(err)
	}
	defer l.Close()
	go func() {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		defer c.Close()
		bytes := make([]byte, 2048)
		for {
			c.Write(bytes)
		}
	}()
	r, err := DialPipe(name, nil)
	if err != nil {
		b.Fatal(err)
	}
	defer r.Close()

    buf := make([]byte, 2048)
	for i := 0; i < b.N; i++ {
		_, err := r.Read(buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}
