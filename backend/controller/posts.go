package controller

import (
	"backend/service"
	"fmt"
)

func GetAllArticles() {
	db := service.DbConnect()
	var (
		id      int
		title   string
		content string
		userId  int
	)

	result, err := db.Query("SELECT * FROM article")
	if err != nil {
		fmt.Println(err)
	} else {
		for result.Next() {
			error := result.Scan(&id, &title, &content, &userId)
			if error != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%v, %v, %v, %v\n", id, title, content, userId)
		}
	}
}

func GetOneArticle(selectedArticleId int) {
	db := service.DbConnect()

	var (
		id      int
		title   string
		content string
	)

	query := `SELECT id, title, content FROM article WHERE id=$1`
	result, err := db.Query(query, selectedArticleId)
	if err != nil {
		fmt.Println(err)
	} else {
		for result.Next() {
			err := result.Scan(&id, &title, &content)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%v, %v, %v\n", id, title, content)
		}

	}
}

func CreateArticle(articleTitle string, articleContent string, articleUserId int) {
	db := service.DbConnect()

	query := `INSERT INTO article ($1, $2, $3)`
	_, err := db.Query(query, articleTitle, articleContent, articleUserId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Artcile ajouté avec succès !")
	}
}
