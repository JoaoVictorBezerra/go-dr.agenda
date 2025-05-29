package repository

import (
	"database/sql"
	"dr.agenda/dto"
	"dr.agenda/enum"
	"dr.agenda/helpers"
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

	helpers.CloseRows(rows)

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

	helpers.CloseStmt(prepare)

	createdInsurance, getByIdErr := repository.GetInsuranceById(id)

	if getByIdErr != nil {
		fmt.Println(getByIdErr)
		return nil, getByIdErr
	}

	return createdInsurance, nil
}

func (repository *InsuranceRepository) SuspendInsurance(id string) (*model.Insurance, error) {
	const query string = "UPDATE Insurance i SET status = $1 WHERE id = $2 RETURNING i.*"

	prepare, prepareErr := repository.database.Prepare(query)

	if prepareErr != nil {
		fmt.Println(prepareErr)
		return nil, prepareErr
	}

	var updatedInsurance model.Insurance
	queryErr := prepare.QueryRow(enum.SUSPENDED, id).Scan(
		&updatedInsurance.Id,
		&updatedInsurance.Name,
		&updatedInsurance.Description,
		&updatedInsurance.Price,
		&updatedInsurance.Benefits,
		&updatedInsurance.Status,
	)

	if queryErr != nil {
		fmt.Println(queryErr)
		return nil, queryErr
	}

	helpers.CloseStmt(prepare)

	return &updatedInsurance, nil

}

func (repository *InsuranceRepository) UpdateInsurance(id string, dto dto.UpdateInsuranceRequest) (*model.Insurance, error) {
	const query string = "UPDATE Insurance i SET name = $2, description = $3, price = $4, benefits = $5 WHERE id = $1 RETURNING i.*"

	prepare, prepareErr := repository.database.Prepare(query)

	if prepareErr != nil {
		fmt.Println(prepareErr)
		return nil, prepareErr
	}

	var updatedInsurance model.Insurance

	queryErr := prepare.QueryRow(id, dto.Name, dto.Description, dto.Price, dto.Benefits).Scan(
		&updatedInsurance.Id,
		&updatedInsurance.Name,
		&updatedInsurance.Description,
		&updatedInsurance.Price,
		&updatedInsurance.Benefits,
		&updatedInsurance.Status,
	)

	if queryErr != nil {
		fmt.Println(queryErr)
		return nil, queryErr
	}

	helpers.CloseStmt(prepare)

	return &updatedInsurance, nil
}
