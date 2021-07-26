package repository

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	entity "referral-server/entity"
)

//GeneratorSQL Generator SQL repo
type GeneratorSQL struct {
	db *sql.DB
}

//NewGeneratorSQL create new repository
func NewGeneratorSQL(db *sql.DB) *GeneratorSQL {
	return &GeneratorSQL{
		db: db,
	}
}

//Create will insert a generator type into SQL based database as repository
func (r *GeneratorSQL) Create(e *entity.Generator) error {
	stmt, err := r.db.Prepare(`
	INSERT INTO 
		generator (id, name, email, password, generated_link, created_at, updated_at, expirate_at) 
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Name,
		e.Email,
		e.Password,
		e.GeneratedLink,
		e.CreatedAt,
		e.UpdatedAt,
		e.ExpirateAt,
	)

	if err != nil {
		me, _ := err.(*pq.Error)

		if string(me.Code) == "23505" {
			return entity.ErrAlreadyExist
		}

		return err
	}

	return stmt.Close()
}

//Create an application
func (r *GeneratorSQL) FetchExpirationTime(ReferralLink string) (Exp time.Time, err error) {
	stmt, err := r.db.Prepare("SELECT expirate_at FROM generator WHERE generated_link = $1")

	if err != nil {
		return
	}

	err = stmt.QueryRow(ReferralLink).Scan(&Exp)

	if err == sql.ErrNoRows {
		err = entity.ErrNotFound
		return
	}

	if err != nil {
		return
	}

	err = stmt.Close()

	return Exp, err
}

// FetchPassword will fetch a password by generator_id
func (r *GeneratorSQL) FetchPassword(GeneratorID string) (Password string, err error) {
	stmt, err := r.db.Prepare("SELECT password FROM generator WHERE id = $1")

	if err != nil {
		return
	}

	err = stmt.QueryRow(GeneratorID).Scan(&Password)

	if err != nil {
		return
	}

	err = stmt.Close()

	return Password, err
}

// FetchReferralLink will fetch a referralLink by generator_id
func (r *GeneratorSQL) FetchReferralLink(GeneratorID string) (ReferralLink string, err error) {
	stmt, err := r.db.Prepare("SELECT generated_link FROM generator WHERE id = $1")

	if err != nil {
		return
	}

	err = stmt.QueryRow(GeneratorID).Scan(&ReferralLink)

	if err != nil {
		return
	}

	err = stmt.Close()

	return ReferralLink, err
}

// UpdateExpirationTime by Referral Link
func (r *GeneratorSQL) UpdateExpirationTime(ReferralLink string, ExpirationTime time.Time) (err error) {
	stmt, err := r.db.Prepare(`
	UPDATE generator
		SET expirate_at = $2, updated_at = $3
	WHERE 
		generated_link = $1;`)

	if err != nil {
		return
	}

	res, err := stmt.Exec(ReferralLink, ExpirationTime, time.Now())

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
