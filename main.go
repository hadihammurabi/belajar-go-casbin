package main

import "log"

func main() {
	http := NewHttp()
	log.Fatal(http.Listen(":8080"))
}
