package model

type Article struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllPosts() ([]Article, error) {
	var articles []Article
	
	query := `select id, title, content from posts;`

	rows, err := db.Query(query)
	if err != nil {
		return articles, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var title, content string

		err := rows.Scan(&id, &title, &content)
		if err != nil {
			return articles, err
		}

		article := Article {
			ID: id,
			Title: title,
			Content: content,
		}

		articles = append(articles, article)
	}
	return articles, nil
}