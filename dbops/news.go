package dbops

import (
	"fmt"
	"game-news/dbops/mysql"
)

type News struct {
	newsId      int
	newsTitle   string
	newsArticle string
	newsUrl     string
	newsMedia   string
	createTime  string
}

func GetAllNews(limit int) ([]News, error) {
	db := mysql.DBCon()

	stmt, err := db.Prepare("select news_id, title, url, media, article from news order by create_time limit ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit)

	if err != nil {
		return nil, err
	}

	var newsData []News
	for rows.Next() {
		news := News{}
		err = rows.Scan(&news.newsId, &news.newsTitle, &news.newsUrl, &news.newsMedia, &news.newsArticle)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		newsData = append(newsData, news)
	}
	
	return newsData, nil
}


func CreateNews(title, url, article, media string) bool {
	db := mysql.DBCon()

	stmt, err := db.Prepare(
		"insert into news (`title`, `url`, `media`, `article`) values (?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, url, media, article)
	if err != nil {
		return false
	}
	return true
}