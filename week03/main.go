package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>おみくじサーバへようこそ</h1>")
	fmt.Fprintln(w, `<p><a href="/webfortune">おみくじを引く</a></p>`)
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	fortunes := []string{"大吉", "中吉", "吉", "凶"}
	rand.Seed(time.Now().UnixNano())
	result := fortunes[rand.Intn(len(fortunes))]

	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, "<h1>今の運勢は %s です</h1>", result)
	fmt.Fprintf(w, `<p><a href="/webfortune">もう一度引く</a></p>`)
	fmt.Fprintf(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/webfortune", fortuneHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("サーバ起動エラー:", err)
}

}
