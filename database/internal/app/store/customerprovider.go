package store

import (
	"database/sql"
	"log"

	"github.com/dmirou/learngo/database/internal/app/model"
)

type CustomerProvider struct {
	store *Store
	l     *log.Logger
}

func (p *CustomerProvider) Find(id int64) (*model.Customer, error) {
	c := &model.Customer{}

	if err := p.store.db.QueryRow(
		"SELECT id, email FROM customer WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return c, nil
}

func (p *CustomerProvider) List() ([]*model.Customer, error) {
	rows, err := p.store.db.Query("SELECT id, email FROM customer")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			p.l.Println(err)
		}
	}()

	var l = make([]*model.Customer, 0)

	for rows.Next() {
		c := &model.Customer{}
		err := rows.Scan(&c.ID, &c.Email)
		if err != nil {
			return nil, err
		}
		l = append(l, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return l, nil
}
