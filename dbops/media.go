package dbops

import (
	"fmt"
	"game-news/dbops/mysql"
)

type Media struct {
	Id int
	Name string
}

func GetMediaName(media_id int) (string, error) {
	db := mysql.DBCon()

	stmt, err := db.Prepare("select name from media where id = ?")
	if err != nil {
		return "", err
	}

	defer stmt.Close()
	var media string
	err = stmt.QueryRow(media_id).Scan(&media)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return media, nil
}

func GetAllMedia() ([]Media, error) {
	db := mysql.DBCon()

	stmt, err := db.Prepare("select id, name from media")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	var MediaData []Media
	for rows.Next() {
		media := Media{}
		err = rows.Scan(&media.Id, &media.Name)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		MediaData = append(MediaData, media)
	}
	return MediaData, nil
}
