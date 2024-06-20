package controller

import (
	"api/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func DeleteTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		tweetIDStr := r.URL.Query().Get("tweetId")
		tweetID, err := strconv.Atoi(tweetIDStr)
		log.Println(tweetID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = usecase.DeleteTweet(tweetID)
		if err != nil {
			log.Printf("fail: usecase.DeleteTweet, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{Success: true})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func DeleteReplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		replyIDStr := r.URL.Query().Get("replyId")
		replyID, err := strconv.Atoi(replyIDStr)
		log.Println(replyID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = usecase.DeleteReply(replyID)
		if err != nil {
			log.Printf("fail: usecase.DeleteTweet, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{Success: true})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}
