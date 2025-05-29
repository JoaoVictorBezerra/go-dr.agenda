package repository

import (
	"database/sql"
	"dr.agenda/dto"
	"dr.agenda/enum"
	"dr.agenda/model"
	"errors"
	"fmt"
)

type InsuranceRepository struct {
	database *sql.DB
}

func NewInsuranceRepository(database *sql.DB) InsuranceRepository {
	return InsuranceRepository{
		database: database,
	}
}

func (repository *InsuranceRepository) GetInsurances() ([]model.Insurance, error) {
	const query string = "SELECT * FROM Insurance hi ORDER BY hi.id"

	rows, sqlErr := repository.database.Query(query)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		return []model.Insurance{}, sqlErr
	}

	var insuranceList []model.Insurance
	var insuranceObj model.Insurance

	for rows.Next() {
		rowsErr := rows.Scan(
			&insuranceObj.Id,
			&insuranceObj.Name,
			&insuranceObj.Description,
			&insuranceObj.Price,
			&insuranceObj.Benefits,
			&insuranceObj.Status,
		)

		if rowsErr != nil {
			fmt.Println(rowsErr)
			return []model.Insurance{}, sqlErr
		}

		insuranceList = append(insuranceList, insuranceObj)
	}
	rows.Close()

	return insuranceList, nil
}

func (repository *InsuranceRepository) GetInsuranceById(id string) (*model.Insurance, error) {
	const query string = "SELECT * FROM Insurance hi WHERE hi.id = $1"

	row := repository.database.QueryRow(query, id)

	var insuranceObj model.Insurance

	scanErr := row.Scan(
		&insuranceObj.Id,
		&insuranceObj.Name,
		&insuranceObj.Description,
		&insuranceObj.Price,
		&insuranceObj.Benefits,
		&insuranceObj.Status,
	)

	if scanErr != nil {
		if errors.Is(scanErr, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, scanErr
	}

	return &insuranceObj, nil
}

func (repository *InsuranceRepository) CreateInsurance(insurance dto.CreateInsuranceRequest) (*model.Insurance, error) {
	var id string

	const query string = "INSERT INTO Insurance (name, description, price, benefits, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	prepare, prepareErr := repository.database.Prepare(query)

	if prepareErr != nil {
		fmt.Println(prepareErr)
		return nil, prepareErr
	}

	queryErr := prepare.QueryRow(insurance.Name, insurance.Description, insurance.Price, insurance.Benefits, enum.ACTIVE).Scan(&id)

	if queryErr != nil {
		fmt.Println(queryErr)
		return nil, queryErr
	}

	prepare.Close()

	createdInsurance, getByIdErr := repository.GetInsuranceById(id)

	if getByIdErr != nil {
		fmt.Println(getByIdErr)
		return nil, getByIdErr
	}

	return createdInsurance, nil
}

func (repository *InsuranceRepository) SuspendInsurance(id string) (*model.Insurance, error) {
	const query string = "UPDATE Insurance SET status = $1 WHERE id = $2"

	prepare, prepareErr := repository.database.Prepare(query)

	if prepareErr != nil {
		fmt.Println(prepareErr)
		return nil, prepareErr
	}

	_, queryErr := prepare.Exec(enum.SUSPENDED, id)

	if queryErr != nil {
		fmt.Println(queryErr)
		return nil, queryErr
	}

	prepare.Close()

	updatedInsurance, getByIdErr := repository.GetInsuranceById(id)

	if getByIdErr != nil {
		fmt.Println(getByIdErr)
		return nil, getByIdErr
	}

	return updatedInsurance, nil

}
