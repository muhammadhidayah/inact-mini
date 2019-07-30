package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/muhammadhidayah/inact-mini/models"
	"github.com/muhammadhidayah/inact-mini/users"
)

type pgUsersRepository struct {
	Conn *sql.DB
}

func NewPgUsersRepository(Conn *sql.DB) users.Repository {
	return &pgUsersRepository{Conn}
}

func (m *pgUsersRepository) fetchData(qry string, args ...interface{}) ([]*models.Users, error) {
	rows, err := m.Conn.Query(qry, args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			fmt.Println(errClose.Error())
		}
	}()

	result := make([]*models.Users, 0)
	for rows.Next() {
		t := new(models.Users)
		rows.Scan(
			&t.ID,
			&t.Username,
			&t.Password,
			&t.IsOnline,
			&t.LastLogin,
		)

		result = append(result, t)
	}

	return result, nil
}

func (m *pgUsersRepository) GetUserById(id int64) (resp *models.Users, err error) {
	qryString := `SELECT * FROM ts_org_person WHERE id = $1`

	res, err := m.fetchData(qryString, id)

	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		resp = res[0]
	} else {
		return nil, errors.New("Your requested Item is not found")
	}

	return
}

func (m *pgUsersRepository) DeleteUserById(id int64) error {
	stmt, err := m.Conn.Prepare("DELETE FROM ts_org_person WHERE id = $1")

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(id)

	return errExec
}

func (m *pgUsersRepository) UpdateUserById(id int64, data *models.Users) error {
	res, err := m.Conn.Prepare("UPDATE ts_org_person SET username = $1, password = $2, last_login = $3, is_online = $4 WHERE id = $5")

	if err != nil {
		return err
	}

	_, err = res.Exec(data.Username, data.Password, data.LastLogin, data.IsOnline, data.ID)

	return nil

}

func (m *pgUsersRepository) InsertUser(data *models.Users) (int, error) {
	var lastInsertID int
	err := m.Conn.QueryRow("INSERT INTO ts_org_person(username, password, last_login, is_online) VALUES($1,$2,$3,$4) returning id;", data.Username, data.Password, data.LastLogin, data.IsOnline).Scan(&lastInsertID)

	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (m *pgUsersRepository) GetUserByUsername(username string) (*models.Users, error) {
	qryString := fmt.Sprintf(`SELECT * FROM ts_org_person where username = $1`)

	res, err := m.fetchData(qryString, username)

	if err != nil {
		return nil, err
	}

	if len(res) > 0 {
		return res[0], nil
	} else {
		return nil, errors.New("")
	}
}
