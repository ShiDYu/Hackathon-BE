package main

import (
	"api/controller"
	"api/dao"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func corsHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	dao.InitDB()
	defer dao.CloseDB()

	http.HandleFunc("/register", corsHandler(controller.RegisterUserHandler))

	http.HandleFunc("/profile", corsHandler(controller.UpdateUserProfileHandler))

	http.HandleFunc("/tweets", corsHandler(controller.GetTweetsHandler))

	http.HandleFunc("/create-tweet", corsHandler(controller.CreateTweetHandler))

	closeDBWithSysCall()

	log.Println("Listening...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func closeDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)
		dao.CloseDB()
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}

/* １SSH接続の設定をしなければならないのか, 2dockerfileがうまくいかないローカルではうまくいく*/
