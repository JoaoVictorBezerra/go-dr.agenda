package repository

import (
	"database/sql"
	"dr.agenda/dto"
	"dr.agenda/enum"
	"dr.agenda/model"
	"fmt"
)

type HealthInsuranceRepository struct {
	database *sql.DB
}

func NewHealthInsuranceRepository(database *sql.DB) HealthInsuranceRepository {
	return HealthInsuranceRepository{
		database: database,
	}
}

func (repository *HealthInsuranceRepository) GetInsurances() ([]model.HealthInsurance, error) {
	const query string = "SELECT * FROM HealthInsurance hi ORDER BY hi.id"

	rows, sqlErr := repository.database.Query(query)

	if sqlErr != nil {
		fmt.Println(sqlErr)
		return []model.HealthInsurance{}, sqlErr
	}

	var insuranceList []model.HealthInsurance
	var insuranceObj model.HealthInsurance

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
			return []model.HealthInsurance{}, sqlErr
		}

		insuranceList = append(insuranceList, insuranceObj)
	}
	rows.Close()

	return insuranceList, nil
}

func (repository *HealthInsuranceRepository) GetInsuranceById(id string) (*model.HealthInsurance, error) {
	const query string = "SELECT * FROM HealthInsurance hi WHERE hi.id = $1"

	row := repository.database.QueryRow(query, id)

	var insuranceObj model.HealthInsurance

	scanErr := row.Scan(
		&insuranceObj.Id,
		&insuranceObj.Name,
		&insuranceObj.Description,
		&insuranceObj.Price,
		&insuranceObj.Benefits,
		&insuranceObj.Status,
	)

	if scanErr != nil {
		return nil, scanErr
	}

	return &insuranceObj, nil
}

func (repository *HealthInsuranceRepository) CreateInsurance(insurance dto.CreateHealthInsuranceRequest) (*model.HealthInsurance, error) {
	var id string

	const query string = "INSERT INTO HealthInsurance (name, description, price, benefits, status) VALUES ($1, $2, $3, $4, $5) RETURNING id"

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

	createdInsurance, getByIdErr := repository.GetInsuranceById(id)

	if getByIdErr != nil {
		fmt.Println(getByIdErr)
		return nil, getByIdErr
	}

	prepare.Close()

	return createdInsurance, nil
}
