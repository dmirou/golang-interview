package store

import (
	"database/sql"
	"log"
	"time"

	"github.com/dmirou/learngo/database/internal/app/model"
)

type CustomerProvider struct {
	store *Store
	l     *log.Logger
}

func (p *CustomerProvider) Add(c *model.Customer) error {
	if err := p.store.db.QueryRow(
		"INSERT INTO customer(email) VALUES ($1) RETURNING id",
		c.Email,
	).Scan(&c.ID); err != nil {
		return err
	}

	return nil
}

func (p *CustomerProvider) BatchAdd(customers []*model.Customer) error {
	p.l.SetPrefix(p.l.Prefix() + "BatchAdd:")

	tx, err := p.store.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			p.l.Println(err)
		}
	}()

	stmt, err := tx.Prepare(
		"INSERT INTO customer(email, first_name, last_name) VALUES ($1, $2, $3)",
	)
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			p.l.Println(err)
		}
	}()

	for _, c := range customers {
		if _, err := stmt.Exec(c.Email, c.FirstName, c.LastName); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *CustomerProvider) Update(c *model.Customer) error {
	_, err := p.store.db.Exec(
		"UPDATE customer SET first_name = $1, last_name = $2",
		c.FirstName,
		c.LastName,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *CustomerProvider) Delete(c *model.Customer) error {
	t := time.Now().UTC()
	_, err := p.store.db.Exec(
		"UPDATE customer SET delete_time = $1",
		t,
	)
	if err != nil {
		return err
	}

	c.DeleteTime = &t

	return nil
}

func (p *CustomerProvider) Find(id int64) (*model.Customer, error) {
	c := &model.Customer{}

	if err := p.store.db.QueryRow(
		"SELECT id, email, first_name, last_name FROM customer WHERE id = $1",
		id,
	).Scan(
		&c.ID,
		&c.Email,
		&c.FirstName,
		&c.LastName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return c, nil
}

func (p *CustomerProvider) FindByEmail(email string) (*model.Customer, error) {
	c := &model.Customer{}

	if err := p.store.db.QueryRow(
		"SELECT id, email, first_name, last_name FROM customer WHERE email = $1",
		email,
	).Scan(
		&c.ID,
		&c.Email,
		&c.FirstName,
		&c.LastName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return c, nil
}

func (p *CustomerProvider) List() ([]*model.Customer, error) {
	p.l.SetPrefix(p.l.Prefix() + "List:")

	rows, err := p.store.db.Query("SELECT id, email, first_name, last_name FROM customer")
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
		err := rows.Scan(
			&c.ID,
			&c.Email,
			&c.FirstName,
			&c.LastName,
		)
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
