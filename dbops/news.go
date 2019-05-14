package dbops

import (
	"fmt"
	"game-news/dbops/mysql"
)

type News struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Media   string `json:"media"`
	Article string `json:"article"`
	Pubtime string `json:"pub_time"`
}

func GetAllNews(limit int) ([]News, error) {
	db := mysql.DBCon()

	stmt, err := db.Prepare("select id, title, url, media, content, pub_time from game_news order by create_time desc limit ?")
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
		err = rows.Scan(&news.Id, &news.Title, &news.Url, &news.Media, &news.Article, &news.Pubtime)
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
