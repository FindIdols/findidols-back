package repository

import (
	"database/sql"
	"time"

	"github.com/FindIdols/findidols-back/entity"
)

//IdolPostgres postgres  repo
type IdolPostgres struct {
	db *sql.DB
}

//NewIdolPostgres create new repository
func NewIdolPostgres(db *sql.DB) *IdolPostgres {
	return &IdolPostgres{
		db: db,
	}
}

//Create a idol
func (r *IdolPostgres) Create(i *entity.Idol) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into idols (idol_id, artistic_name, profession, description, value, 
		deadline, user_id, social_networks_id, bank_account_id, created_at) 
		values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`)

	if err != nil {
		return i.ID, err
	}
	_, err = stmt.Exec(
		i.ID,
		i.ArtisticName,
		i.Profession,
		i.Description,
		i.Value,
		i.Deadline,
		i.UserID,
		i.SocialNetworksID,
		i.BankAccountID,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return i.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return i.ID, err
	}

	return i.ID, nil
}

//GetIdol info
func (r *IdolPostgres) GetIdol(id entity.ID) (*entity.Idol, error) {
	stmt, err := r.db.Prepare(`select idol_id, artistic_name, profession, description, 
	value, deadline, deny_content, user_id, social_networks_id, created_at from idols where idol_id = $1`)

	if err != nil {
		return nil, err
	}

	var idol entity.Idol
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&idol.ID,
			&idol.ArtisticName,
			&idol.Profession,
			&idol.Description,
			&idol.Value,
			&idol.Deadline,
			&idol.DenyContent,
			&idol.UserID,
			&idol.SocialNetworksID,
			&idol.CreatedAt,
		)
	}

	return &idol, nil
}

//GetIdols info
func (r *IdolPostgres) GetIdols() ([]*entity.Idol, error) {
	stmt, err := r.db.Prepare(`select idol_id, artistic_name, description, 
	value from idols`)

	if err != nil {
		return nil, err
	}

	var idols []*entity.Idol
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var idol entity.Idol

		err = rows.Scan(
			&idol.ID,
			&idol.ArtisticName,
			&idol.Description,
			&idol.Value,
		)

		if err != nil {
			return nil, err
		}

		idols = append(idols, &idol)

	}

	return idols, nil
}
