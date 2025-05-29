package repository

import (
	"database/sql"
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

	return insuranceList, nil
}
