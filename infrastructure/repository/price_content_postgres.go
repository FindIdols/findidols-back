package repository

import (
	"database/sql"

	"github.com/FindIdols/findidols-back/entity"
)

//PriceContentPostgres postgres repo
type PriceContentPostgres struct {
	db *sql.DB
}

//NewSocialNetworksPostgres create new repository
func NewPriceContentPostgres(db *sql.DB) *PriceContentPostgres {
	return &PriceContentPostgres{
		db: db,
	}
}

//GetPricePerContent a data
func (r *PriceContentPostgres) GetPricePerContent(id entity.ID) (*entity.PriceContent, error) {
	stmt, err := r.db.Prepare(`select price_per_content_id, general, publicity, politics, professional_tips, earnings 
	from price_per_content where price_per_content_id = $1`)

	if err != nil {
		return nil, err
	}

	var priceContent entity.PriceContent
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&priceContent.ID,
			&priceContent.General,
			&priceContent.Publicity,
			&priceContent.Politics,
			&priceContent.ProfessionalTips,
			&priceContent.Earnings,
		)
	}

	return &priceContent, nil
}

//GetPricesPerContent info
func (r *PriceContentPostgres) GetPricesPerContent() ([]*entity.PriceContent, error) {
	stmt, err := r.db.Prepare(`select price_per_content_id, general, publicity, politics, professional_tips, earnings 
	from price_per_content`)

	if err != nil {
		return nil, err
	}

	var prices []*entity.PriceContent
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var price entity.PriceContent

		err = rows.Scan(
			&price.ID,
			&price.General,
			&price.Publicity,
			&price.Politics,
			&price.ProfessionalTips,
			&price.Earnings,
		)

		if err != nil {
			return nil, err
		}

		prices = append(prices, &price)

	}

	return prices, nil
}
