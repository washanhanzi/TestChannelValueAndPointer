package TestChannelValueAndPointer

import (
	"testing"
)

type BigStruct struct {
	Data [500]byte
}

func BenchmarkValue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := make(chan BigStruct)
		go func() {
			for i := 0; i < 100; i++ {
				a := BigStruct{Data: [500]byte{}}
				c <- a
			}
		}()
		d := <-c
		d.Data[0] = 1
	}
}

func BenchmarkPointer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := make(chan *BigStruct)
		go func() {
			for i := 0; i < 100; i++ {
				a := BigStruct{Data: [500]byte{}}
				c <- &a
			}
		}()
		d := <-c
		d.Data[0] = 1
	}
}
