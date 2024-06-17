package controller

import (
	"api/model"
	"api/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetLikesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		postID := r.URL.Query().Get("postId")
		if postID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userID := r.URL.Query().Get("userId")
		if userID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//ここでuserIDが適切に取れてきていないのでbad requestエラー
		response, err := usecase.GetLikes(postID, userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		json.NewEncoder(w).Encode(response)
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func LikesPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		log.Printf("postリクエストOK")
		var like model.Like
		err := json.NewDecoder(r.Body).Decode(&like)
		fmt.Println(like)
		if err != nil {
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}
		post, err := usecase.LikePost(like)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]int{"likeCount": post.LikeCount})
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UnlikesPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodPost:
		var like model.Like
		json.NewDecoder(r.Body).Decode(&like)
		post, err := usecase.UnlikePost(like)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]int{"likeCount": post.LikeCount})
	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}
