package TestChannelValueAndPointer

type BigStruct struct {
	Data [500]byte
}

func main() {
	a := BigStruct{Data: [500]byte{}}
	c := make(chan BigStruct)
	go func() {
		for i := 0; i < 100; i++ {
			c <- a
		}
	}()
	d := <-c
	d.Data[0] = 1
}
