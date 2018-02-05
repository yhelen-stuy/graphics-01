package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("img.ppm", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	var buffer bytes.Buffer
	buffer.WriteString("P3 500 500 255\n")

	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			buffer.WriteString(fmt.Sprintf("%d %d %d\n", (255-i)%256, (255-j)%256, (255-i*j)%256))
		}
	}

	f.WriteString(buffer.String())
	f.Close()
}
