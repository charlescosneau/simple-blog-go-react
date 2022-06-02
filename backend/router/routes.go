package router

import (
	"backend/controller"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/articles", controller.GetAllArticles)
	http.HandleFunc("/article", controller.GetOneArticle)
	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		return
	}
}
