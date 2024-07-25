package user

import (
	"database/sql"
	"fmt"

	"github.com/PCD/ecom/types"
)

type Store struct {
	db *sql.DB
}

// Crea una nueva instancia de Store.
func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Recupera un usuario de la base de datos por su correo electrónico.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// Ayuda a mapear los resultados de la consulta a un objeto User.
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}
