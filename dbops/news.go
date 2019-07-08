package dbops

import (
	"fmt"
	"game-news/dbops/mysql"
	"log"
)

type News struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Url     string `json:"url"`
	Media   string `json:"media"`
	Article string `json:"article"`
	Pubtime string `json:"pub_time"`
}

func GetAllNews(begin, end int) ([]News, error) {
	db := mysql.DBCon()
	stmt, err := db.Prepare("select n.id, n.title, n.url, m.name, n.content, n.pub_time from news n inner join media m on n.media_id = m.id order by n.create_time desc limit ?,?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(begin, end)

	if err != nil {
		return nil, err
	}

	var newsData []News
	for rows.Next() {
		news := News{}
		err = rows.Scan(&news.Id, &news.Title, &news.Url, &news.Media, &news.Article, &news.Pubtime)
		if err != nil {
			log.Println(err.Error())
		}

		newsData = append(newsData, news)
	}

	return newsData, nil
}

func GetNewsByMedia(begin, end, mediaId int) ([]News, error) {
	db := mysql.DBCon()

	stmt, err := db.Prepare("select id, title, url, content, pub_time from news where media_id = ? order by create_time desc limit ?,?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	media, err := GetMediaName(mediaId)
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := stmt.Query(mediaId, begin, end)

	var newsData []News
	for rows.Next() {
		news := News{}
		news.Media = media
		err = rows.Scan(&news.Id, &news.Title, &news.Url, &news.Article, &news.Pubtime)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		newsData = append(newsData, news)
	}

	return newsData, nil
}

func CreateNews(title, url, content, media, pub_time string) bool {
	db := mysql.DBCon()

	stmt, err := db.Prepare(
		"insert into news (`title`, `url`, `media`, `content`, `pub_time`) values (?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, url, media, content, pub_time)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
