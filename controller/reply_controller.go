package controller

import (
	"api/model"
	"api/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func CreateReplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var reply model.Reply
		err := json.NewDecoder(r.Body).Decode(&reply)
		if err != nil {
			log.Printf("fail: json.NewDecoder, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = usecase.CreateReply(reply)
		if err != nil {
			log.Printf("fail: usecase.UpdateUserProfile, %v\n", err)
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

func GetRepliesByTweetIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tweetIDStr := r.URL.Query().Get("tweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr) //ここでIDをint型に型変換している
	if err != nil {
		http.Error(w, "Invalid tweet_id", http.StatusBadRequest)
		return
	}
	tweets, err := usecase.GetReplies(tweetID)
	if err != nil {
		log.Printf("fail: usecase.GetTweets, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}

func GetReplyCountByTweetIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tweetIDStr := r.URL.Query().Get("tweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		http.Error(w, "Invalid tweet_id", http.StatusBadRequest)
		return
	}

	count, err := usecase.GetReplyCount(tweetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

func GetRepliedTweetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tweetIDStr := r.URL.Query().Get("tweet_id")
	tweetID, err := strconv.Atoi(tweetIDStr)
	if err != nil {
		http.Error(w, "Invalid tweet_id", http.StatusBadRequest)
		return
	}

	tweet, err := usecase.GetRepliedTweet(tweetID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweet)
}
