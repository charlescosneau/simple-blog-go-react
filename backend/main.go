package main

import (
	"backend/controller"
	_ "github.com/lib/pq"
)

func main() {
	controller.CreateArticle("Je suis un artcile creé depuis le code", "Le super content de fou malade", 1)
	controller.GetAllArticles()
	//controller.GetOneArticle(2)
}
