package user

import (
	"context"
	"crud/models"
	repo "crud/service"
	token "crud/utils"
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
	query := `INSERT INTO [dbo].[users]
	([username]
	,[email]
	,[password])
VALUES (?,?,?)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, u.Username, u.Email, newhashed)
	if err != nil {
		println(err.Error())
		return -1, err
	}

	return res.RowsAffected()
}

func (m *userRepo) Login(ctx context.Context, username string, password string) (*string, error) {
	// Find record by username
	query := `SELECT * FROM [dbo].[users] WHERE username = ? `
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	// Get token
	token, err := token.GenerateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Get all users
func (m *userRepo) Fetch(ctx context.Context) ([]*models.User, error) {
	query := `SELECT * FROM [dbo].[users]`
	rows, err := m.Conn.Query(query)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)
		err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Email,
			&data.Password,
			&data.CreatedAt,
		)
		if err != nil {
			println("97" + err.Error())
			return nil, err
		}
		users = append(users, data)
	}
	return users, nil
}
