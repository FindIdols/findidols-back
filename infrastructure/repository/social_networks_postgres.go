package repository

import (
	"database/sql"
	"time"

	"github.com/FindIdols/findidols-back/entity"
)

//SocialNetworksPostgres postgres repo
type SocialNetworksPostgres struct {
	db *sql.DB
}

//NewSocialNetworksPostgres create new repository
func NewSocialNetworksPostgres(db *sql.DB) *SocialNetworksPostgres {
	return &SocialNetworksPostgres{
		db: db,
	}
}

//Create a social networks
func (r *SocialNetworksPostgres) Create(sn *entity.SocialNetworks) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into social_networks (social_networks_id, youtube, instagram, twitter, tiktok, created_at) 
		values($1,$2,$3,$4,$5,$6)`)

	if err != nil {
		return sn.ID, err
	}
	_, err = stmt.Exec(
		sn.ID,
		sn.Youtube,
		sn.Instagram,
		sn.Twitter,
		sn.TikTok,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return sn.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return sn.ID, err
	}

	return sn.ID, nil
}

//GetSocialNetworks a data
func (r *SocialNetworksPostgres) GetSocialNetworks(id entity.ID) (*entity.SocialNetworks, error) {
	stmt, err := r.db.Prepare(`select social_networks_id, youtube, instagram, twitter, tiktok 
	from social_networks where social_networks_id = $1`)

	if err != nil {
		return nil, err
	}

	var socialNetworks entity.SocialNetworks
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&socialNetworks.ID,
			&socialNetworks.Youtube,
			&socialNetworks.Instagram,
			&socialNetworks.Twitter,
			&socialNetworks.TikTok,
		)
	}

	return &socialNetworks, nil
}
