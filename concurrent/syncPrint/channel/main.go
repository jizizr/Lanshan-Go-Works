package main

func main() {
	sig := make(chan struct{})
	over := make(chan struct{})
	printEven := func() {
		for i := 0; i < 100; i += 2 {
			println(i)
			sig <- struct{}{}
			<-sig
		}
	}
	printOdd := func() {
		for i := 1; i < 100; i += 2 {
			<-sig
			println(i)
			sig <- struct{}{}
		}
		over <- struct{}{}
	}
	go printEven()
	go printOdd()
	<-over
}
