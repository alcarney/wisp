package main

var buf [1024]byte

// Functions to be implemented by the JS wrapper
func log(message string)

//go:export echo
func echo(message string) {
	log(message)
}

//go:export getBuffer
func getBuffer() *byte {
	return &buf[0]
}

func main() {

}
