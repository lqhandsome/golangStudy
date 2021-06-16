package main

var done = make(chan bool)
var msg string

func aGoroutine(){
	msg = "hello,world"
	//done<-true
	close(done)
}

func main() {
	go aGoroutine()
	for {
		<-done
		println(msg)
	}

}

