package repository

import (
	"database/sql"
	"fmt"

	"github.com/FindIdols/findidols-back/entity"
)

//VideoPostgres postgres  repo
type VideoPostgres struct {
	db *sql.DB
}

//NewVideoPostgres create new repository
func NewVideoPostgres(db *sql.DB) *VideoPostgres {
	return &VideoPostgres{
		db: db,
	}
}

//Get a video
func (r *VideoPostgres) Get(id entity.ID) (*entity.Video, error) {
	stmt, err := r.db.Prepare(`select video_id, url from videos where idol_id = $1`)

	if err != nil {
		return nil, err
	}

	var video entity.Video
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&video.ID,
			&video.URL,
		)
	}

	return &video, nil
}

func (r *VideoPostgres) GetVideos(id entity.ID) ([]*entity.Video, error) {
	stmt, err := r.db.Prepare(`(SELECT video_id, url, introduction 
		FROM videos WHERE introduction = true AND idol_id = $1 LIMIT 1)
		UNION ALL
		(SELECT video_id, url, introduction 
		FROM videos where idol_id = $1 LIMIT 2)`)

	if err != nil {
		return nil, err
	}

	var videos []*entity.Video
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var video entity.Video

		err = rows.Scan(
			&video.ID,
			&video.URL,
			&video.Introduction,
		)

		fmt.Println(err)

		if err != nil {
			return nil, err
		}

		videos = append(videos, &video)

	}

	return videos, nil
}

// //UploadVideo a video
// func (r *VideoPostgres) UploadVideo(v *entity.Video) (entity.ID, error) {
// 	stmt, err := r.db.Prepare(`
// 		insert into videos (id, '', idol_id, created_at)
// 		values($1,$2,$3)`)

// 	if err != nil {
// 		return v.ID, err
// 	}
// 	_, err = stmt.Exec(
// 		v.ID,
// 		v.IdolID,
// 		time.Now().Format("2006-01-02"),
// 	)

// 	if err != nil {
// 		return v.ID, err
// 	}

// 	err = stmt.Close()

// 	if err != nil {
// 		return v.ID, err
// 	}

// 	return v.ID, nil
// }
