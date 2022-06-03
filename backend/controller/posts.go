package controller

import (
	"backend/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type Articles struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db := service.DbConnect()
	var (
		id          int
		title       string
		content     string
		userId      int
		finalResult []Articles
	)

	result, err := db.Query("SELECT * FROM article")
	if err != nil {
		fmt.Println(err)
	} else {
		for result.Next() {
			err2 := result.Scan(&id, &title, &content, &userId)
			if err2 != nil {
				fmt.Println(err2)
				return
			}
			fmt.Printf("%v, %v, %v, %v\n\n\n", id, title, content, userId)
			finalResult = append(finalResult, Articles{id, title, content, userId})
		}
		err := json.NewEncoder(w).Encode(finalResult)
		if err != nil {
			return
		}
	}
}

func GetOneArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := service.DbConnect()

	var (
		id                int
		title             string
		content           string
		selectedArticleId = r.URL.Query().Get("id")
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
		err := json.NewEncoder(w).Encode(Article{Id: id, Title: title, Content: content})
		if err != nil {
			return
		}
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {

	type IntendedBody struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	i := IntendedBody{}

	w.Header().Set("Content-Type", "application/json")
	db := service.DbConnect()

	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		return
	}
	
	fmt.Printf("%v", i)

	query := `INSERT INTO article (title, content, userid) VALUES ($1, $2, $3)`
	_, err2 := db.Query(query, i.Title, i.Content, i.ID)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		err := json.NewEncoder(w).Encode("Article created !")
		if err != nil {
			return
		}
	}
}
