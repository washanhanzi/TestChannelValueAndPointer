package TestChannelValueAndPointer

import (
	"testing"
)

func Benchmark(b *testing.B) {
	b.Run("500byte value", func(b *testing.B) {
		type bigStruct struct {
			Data [500]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [500]byte{}}
					c <- a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
	b.Run("500byte pointer", func(b *testing.B) {
		type bigStruct struct {
			Data [500]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan *bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [500]byte{}}
					c <- &a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
	b.Run("1000byte value", func(b *testing.B) {
		type bigStruct struct {
			Data [1000]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [1000]byte{}}
					c <- a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
	b.Run("1000byte pointer", func(b *testing.B) {
		type bigStruct struct {
			Data [1000]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan *bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [1000]byte{}}
					c <- &a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
	b.Run("10000byte value", func(b *testing.B) {
		type bigStruct struct {
			Data [10000]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [10000]byte{}}
					c <- a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
	b.Run("10000byte pointer", func(b *testing.B) {
		type bigStruct struct {
			Data [10000]byte
		}
		for n := 0; n < b.N; n++ {
			c := make(chan *bigStruct)
			go func() {
				for i := 0; i < 100; i++ {
					a := bigStruct{Data: [10000]byte{}}
					c <- &a
				}
			}()
			d := <-c
			d.Data[0] = 1
		}
	})
}
