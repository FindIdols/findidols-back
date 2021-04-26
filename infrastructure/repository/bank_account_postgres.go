package repository

import (
	"database/sql"
	"time"

	"github.com/FindIdols/findidols-back/entity"
)

//BankAccountPostgres postgres repo
type BankAccountPostgres struct {
	db *sql.DB
}

//NewBankAccountPostgres create new repository
func NewBankAccountPostgres(db *sql.DB) *BankAccountPostgres {
	return &BankAccountPostgres{
		db: db,
	}
}

//Create a social networks
func (r *BankAccountPostgres) Create(ba *entity.BankAccount) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into bank_accounts (bank_account_id, bank_name, type, agency, operation, account, digit, created_at) 
		values($1,$2,$3,$4,$5,$6,$7,$8)`)

	if err != nil {
		return ba.ID, err
	}
	_, err = stmt.Exec(
		ba.ID,
		ba.BankName,
		ba.Type,
		ba.Agency,
		ba.Operation,
		ba.Account,
		ba.Digit,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return ba.ID, err
	}

	err = stmt.Close()
	if err != nil {
		return ba.ID, err
	}

	return ba.ID, nil
}
