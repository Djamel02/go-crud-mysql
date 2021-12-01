package employee

import (
	"context"
	"crud/models"
	repo "crud/service"
	transaction "crud/utils"
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
			&data.Balance,
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
			,[balance]
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
			,[balance]
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
		,[postalcode]
		,[balance])
	VALUES (?,?,?,?,?,?,?,?)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, p.Name, p.Phone, p.Picture, p.Job, p.Country, p.City, p.Postalcode, p.Balance)
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

// Money Transaction
func (m *empRepo) Trasaction(ctx context.Context, amount float64, senderId int64, receiverId int64) (error) {
	// Check if ids exists
	_, err := m.GetByID(ctx, senderId)
	if err != nil {
		return errors.New("Sender doesnot exist with this id")
	}
	_, err = m.GetByID(ctx, receiverId)
	if err != nil {
		return errors.New("Receiver doesnot exist with this id")
	}
	// Start transaction
	query1 := "UPDATE [dbo].[employees] SET [balance]=[balance]-? WHERE ([id]=? AND [balance] > ?)"
	query2 := "UPDATE [dbo].[employees] SET [balance]=[balance]+? WHERE [id]=?"
	return transaction.Transact(m.Conn, func (tx *sql.Tx) error {
		// Decrement
		res, err := tx.ExecContext(ctx ,query1 ,amount ,senderId, amount)
		if err != nil {
			return err
		}
		if r, err := res.RowsAffected(); r < 1 || err != nil {
			println(err.Error())
			return err
		}
		// Increment
		if _, err := tx.ExecContext(ctx, query2,amount, receiverId); err != nil {
			return err
		}
		return nil
	})

}
