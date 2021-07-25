package repository

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	entity "Referral-System/generator/entity"
)

//ContributorSQL Contributor SQL repo
type ContributorSQL struct {
	db *sql.DB
}

//NewContributorSQL create new repository
func NewContributorSQL(db *sql.DB) *ContributorSQL {
	return &ContributorSQL{
		db: db,
	}
}

// Contribute will insert a Contributor type into SQL based database as repository
// First it will update a row in contributor table
// If it was not exist, then the function will insert a new row
func (r *ContributorSQL) Contribute(e *entity.Contributor) (err error) {
	err = r.Update(e)

	if err == entity.ErrNotFound {
		err = r.Create(e)
	}

	return
}

// Create contributor
func (r *ContributorSQL) Create(e *entity.Contributor) (err error) {
	stmt, err := r.db.Prepare(`
	INSERT INTO 
		contributor (email, generated_link, contribution, created_at, updated_at) 
	VALUES 
		($1, $2, $3, $4, $5)`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Email,
		e.ReferralLink,
		1,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		me, _ := err.(*pq.Error)

		// update if row already exist
		if string(me.Code) == "23505" {
			return entity.ErrAlreadyExist
		}

		return err
	}

	return stmt.Close()
}

func (r *ContributorSQL) Update(e *entity.Contributor) (err error) {
	stmt, err := r.db.Prepare(`
	UPDATE contributor 
		SET contribution = contribution + 1, updated_at = $3
	WHERE 
		email = $1 AND generated_link = $2;`)

	if err != nil {
		return
	}

	res, err := stmt.Exec(e.Email, e.ReferralLink, time.Now())

	count, _ := res.RowsAffected()

	if count == 0 {
		err = entity.ErrNotFound
		return
	}

	if err != nil {
		return
	}

	err = stmt.Close()

	return err
}
