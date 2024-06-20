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
	http.HandleFunc("/profilecard", corsHandler(controller.GetProfileHandler))
	http.HandleFunc("/save-avatar", corsHandler(controller.SetAvatarHandler))

	http.HandleFunc("/avatar", corsHandler(controller.GetAvatarHandler))

	http.HandleFunc("/tweets", corsHandler(controller.GetTweetsHandler))
	http.HandleFunc("/tweets/delete", corsHandler(controller.DeleteTweetHandler))

	http.HandleFunc("/create-tweet", corsHandler(controller.CreateTweetHandler))

	// CORSを処理するハンドラをラップ
	http.HandleFunc("/posts/likes", corsHandler(controller.GetLikesHandler))
	http.HandleFunc("/posts/like", corsHandler(controller.LikesPostHandler))
	http.HandleFunc("/posts/unlike", corsHandler(controller.UnlikesPostHandler))
	http.HandleFunc("/reply", corsHandler(controller.CreateReplyHandler))
	http.HandleFunc("/replies", corsHandler(controller.GetRepliesByTweetIDHandler))
	http.HandleFunc("/replies/delete", corsHandler(controller.DeleteReplyHandler))
	http.HandleFunc("/reply/count", corsHandler(controller.GetReplyCountByTweetIDHandler))
	http.HandleFunc("/repliedTweet", corsHandler(controller.GetRepliedTweetHandler))
	//ここから
	http.HandleFunc("/replies/replies", corsHandler(controller.GetRepliesToReply))
	http.HandleFunc("/replies/likes", controller.GetReplyLikes) //ここ保留 likesの取得とリプライの取得
	http.HandleFunc("/replies/like", corsHandler(controller.LikeReply))
	http.HandleFunc("/replies/unlike", corsHandler(controller.UnlikeReply))
	http.HandleFunc("/replytoreply", corsHandler(controller.CreateReplyToReplyHandler))
	http.HandleFunc("/reply_replies/count", corsHandler(controller.GetReplyCount))

	http.HandleFunc("/generate-tweet", corsHandler(controller.GenerateTweetHandler))
	http.HandleFunc("/update-tweet", corsHandler(controller.UpdateTweetHandler))
	http.HandleFunc("/ai-reply", corsHandler(controller.GenerateReplyHandler))

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
