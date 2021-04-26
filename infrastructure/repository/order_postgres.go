package repository

import (
	"database/sql"
	"time"

	"github.com/FindIdols/findidols-back/entity"
)

//OrderPostgres postgres repo
type OrderPostgres struct {
	db *sql.DB
}

//NewOrderPostgres create new repository
func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

//Create a order
func (r *OrderPostgres) Create(o *entity.Order) (*entity.Order, error) {
	stmt, err := r.db.Prepare(`
		insert into orders (order_id, order_number, usage, subject, instruction, terms_of_use, idol_id, user_id, created_at) 
		values($1,nextval('order_number_increments'),$2,$3,$4,$5,$6,$7,$8)`)

	if err != nil {
		return o, err
	}

	_, err = stmt.Exec(
		o.ID,
		o.Content.Usage,
		o.Content.Subject,
		o.Content.Instruction,
		o.TermsOfUse,
		o.IdolID,
		o.UserID,
		time.Now().Format("2006-01-02"),
	)

	if err != nil {
		return o, err
	}

	order, err := getOrder(o.ID, r.db)

	o.OrderNumber = order.OrderNumber

	if err != nil {
		return o, err
	}

	err = stmt.Close()
	if err != nil {
		return o, err
	}

	return o, nil
}

//Get a order
func (r *OrderPostgres) Get(id entity.ID) (*entity.Order, error) {
	return getOrder(id, r.db)
}

func getOrder(id entity.ID, db *sql.DB) (*entity.Order, error) {
	stmt, err := db.Prepare(`select order_id, order_number from orders where order_id = $1`)

	if err != nil {
		return nil, err
	}

	var order entity.Order
	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&order.ID,
			&order.OrderNumber,
		)
	}

	return &order, nil
}
