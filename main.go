package TestChannelValueAndPointer

type bigStruct struct {
	Data [500]byte
}

func main() {
	a := bigStruct{Data: [500]byte{}}
	c := make(chan bigStruct)
	go func() {
		for i := 0; i < 100; i++ {
			c <- a
		}
	}()
	d := <-c
	d.Data[0] = 1
}
