package repository

import (
	"database/sql"

	"github.com/lib/pq"

	entity "Referral-System/generator/entity"
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
