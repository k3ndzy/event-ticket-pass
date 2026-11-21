package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	addr := ":8095"
	fmt.Printf("🎟️ 🎟️ EventPass - Music Festival & Conference Ticketing Server đang chạy tại http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
