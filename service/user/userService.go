package user

import (
	"context"
	"crud/models"
	repo "crud/service"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	Conn *sql.DB
}

func NewUserRepo(Conn *sql.DB) repo.UserRepo {
	return &userRepo{
		Conn: Conn,
	}
}

func (m *userRepo) Register(ctx context.Context, u *models.User) (int64, error) {
	// Hash password
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	// Convert byte to string
	newhashed := string(hashedPass[:])
	query := `INSERT INTO user (username, email, password) VALUES(?,?,?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, u.Username, u.Email, newhashed)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (m *userRepo) Login(ctx context.Context, username string, password string) (*models.User, error) {
	// Find record by username
	query := `SELECT * FROM user WHERE username = ? `
	row, err := m.Conn.QueryContext(ctx, query, username)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	user := new(models.User)
	for row.Next() {
		if err := row.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
	}
	println(user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
