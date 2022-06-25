package main

// maxBufferSize specifies the size of the buffers that
// are used to temporarily hold data from the UDP packets
// that we receive.
const maxBufferSize = 1024

func main() {
	buffer := make([]byte, maxBufferSize)
	ServerUDP("127.0.0.1", 8080, &buffer)
}
