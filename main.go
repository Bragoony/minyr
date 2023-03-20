package main

import (
	"os"
	"log"
)
func main() {
	src, err := os.Open("/home/brageA/minyr/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	log.Println(src)

	
	var buffer []byte
	buffer = make([]byte, 1)

	for ; n != 0 {
	n, err := src.Read(buffer)
	if err != nil {
		log.Fatal(err)
		}
	}
	log.Println(string(buffer[:n]))
}
