package employee

import (
	"context"
	"crud/models"
	repo "crud/service"
	"database/sql"
	"errors"
)

type empRepo struct {
	Conn *sql.DB
}

// NewEmpRepo retunrs implement of employee repository interface
func NewEmpRepo(Conn *sql.DB) repo.EmpRepo {
	return &empRepo{
		Conn: Conn,
	}
}

// define fetch method
func (m *empRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Employee, error) {
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
			&data.Picture,
			&data.Job,
			&data.Country,
			&data.City,
			&data.Postalcode,
			&data.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

// Get employees list
func (m *empRepo) Fetch(ctx context.Context) ([]*models.Employee, error) {
	query := `Select [id] 
			,[name]
			,[phone]
			,[picture]
			,[job] 
			,[country]
			,[city]
			,[postalcode]
			,[created_at]
			 From [dbo].[employees]`

	return m.fetch(ctx, query)
}

// Get employee by id
func (m *empRepo) GetByID(ctx context.Context, id int64) (*models.Employee, error) {
	query := `Select
			[id] 
			,[name]
			,[phone]
			,[picture]
			,[job] 
			,[country]
			,[city]
			,[postalcode]
			,[created_at]
			From [dbo].[employees] where [id]=?`

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
func (m *empRepo) Create(ctx context.Context, p *models.Employee) (int64, error) {
	query := `INSERT INTO [dbo].[employees]
		([name]
		,[phone]
		,[picture]
		,[job]
		,[country]
		,[city]
		,[postalcode])
	VALUES (?,?,?,?,?,?,?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Name, p.Phone, p.Picture, p.Job, p.Country, p.City, p.Postalcode)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

// Update Employee
func (m *empRepo) Update(ctx context.Context, p *models.Employee, id int64) (*models.Employee, error) {
	query := "UPDATE [dbo].[employees] set [name]=?, [phone]=?,[job] =? ,[country] =?,[city] =? ,[postalcode]=?  where [id]=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Name,
		p.Phone,
		p.Job,
		p.Country,
		p.City,
		p.Postalcode,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

// Delete Employee
func (m *empRepo) Delete(ctx context.Context, id int64) (int64, error) {
	query := "DELETE FROM [employees] WHERE [id]=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
