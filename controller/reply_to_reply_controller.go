package controller

import (
	_ "api/dao"
	"api/model"
	"api/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetRepliesToReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		replyID, err := strconv.Atoi(r.URL.Query().Get("reply_id"))
		if err != nil {
			http.Error(w, "Invalid reply_id", http.StatusBadRequest)
			return
		}

		replies, err := usecase.GetRepliesTOReply(replyID)
		if err != nil {
			http.Error(w, "Error getting replies to reply", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(replies)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetReplyCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		replyID, err := strconv.Atoi(r.URL.Query().Get("reply_id"))
		if err != nil {
			http.Error(w, "Invalid reply_id", http.StatusBadRequest)
			return
		}

		count, err := usecase.GetReplyToReplyCount(replyID)
		if err != nil {
			http.Error(w, "Error getting reply likes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"count": count})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func LikeReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var like model.ReplyLike
		if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		post, err := usecase.ReplyLike(like.ReplyId, like.Uid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]int{"likeCount": post.LikeCount})
		w.WriteHeader(http.StatusNoContent)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UnlikeReply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var like model.ReplyLike
		if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		post, err := usecase.UnReplyLike(like.ReplyId, like.Uid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(map[string]int{"likeCount": post.LikeCount})
		w.WriteHeader(http.StatusNoContent)

		w.WriteHeader(http.StatusNoContent)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func CreateReplyToReplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var reply model.ReplyToReply
		err := json.NewDecoder(r.Body).Decode(&reply)
		if err != nil {
			log.Printf("fail: json.NewDecoder, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf(strconv.Itoa(reply.ReplyId))

		err = usecase.CreateReplyToReply(reply)
		if err != nil {
			log.Printf("fail: usecase.CreatereplyToReply, %v\n", err)
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

func GetReplyLikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodGet:
		replyID, err := strconv.Atoi(r.URL.Query().Get("replyId"))
		if err != nil {
			http.Error(w, "Invalid reply_id", http.StatusBadRequest)
			return
		}
		userID := r.URL.Query().Get("userId")
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}
		
		log.Printf(strconv.Itoa(replyID), userID)
		response, err := usecase.GetReplyLikes(replyID, userID)
		if err != nil {
			http.Error(w, "Error getting replies to reply", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}
