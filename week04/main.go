package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	// 現在時刻取得
	now := time.Now().Format("15:04:05")

	// ブラウザ情報取得
	userAgent := r.Header.Get("User-Agent")

	// レスポンス作成
	message := fmt.Sprintf("今の時刻は %s で，利用しているブラウザは %s ですね", now, userAgent)

	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/", infoHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

