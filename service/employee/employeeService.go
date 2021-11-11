package employee

import (
	"context"
	"database/sql"
	"errors"

	models "crud/models"
	repo "crud/service"
)

// NewSQLEmpRepo retunrs implement of employee repository interface
func NewSQLEmpRepo(Conn *sql.DB) repo.EmpRepo {
	return &mysqlEmpRepo{
		Conn: Conn,
	}
}

type mysqlEmpRepo struct {
	Conn *sql.DB
}

// define fetch method
func (m *mysqlEmpRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Employee, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*models.Employee, 0)
	for rows.Next() {
		data := new(models.Employee)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Phone,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// Get employees list
func (m *mysqlEmpRepo) Fetch(ctx context.Context) ([]*models.Employee, error) {
	query := "Select id, name, phone From employee"

	return m.fetch(ctx, query)
}

// Get employee by id
func (m *mysqlEmpRepo) GetByID(ctx context.Context, id int) (*models.Employee, error) {
	query := "Select id, name, phone From employee where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Employee{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, errors.New("Resquesed item is not found!")
	}

	return payload, nil
}

// Create new Employee
func (m *mysqlEmpRepo) Create(ctx context.Context, p *models.Employee) (int64, error) {
	query := "INSERT INTO employee (name,phone) VALUES(?,?) "

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Phone)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

// Update Employee
func (m *mysqlEmpRepo) Update(ctx context.Context, p *models.Employee) (*models.Employee, error) {
	query := "UPDATE employee set name=?, phone=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Phone,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Delete Employee
func (m *mysqlEmpRepo) Delete(ctx context.Context, id int) (sql.Result, error) {
	query := "DELETE FROM employee WHERE id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	deleted, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return nil, err
	}
	return deleted, nil
}
