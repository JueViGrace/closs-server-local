package types

import "github.com/JueViGrace/clo-backend/internal/db"

type Salesman struct{}

type CreateSalesmanRequest struct{}

type UpdateSalesmanRequest struct{}

func DbSalesmanToSalesman(db *db.GetSalesmanByCodeRow) *Salesman {
	return &Salesman{}
}

func DbExistingSalesmanToSalesman(db *db.GetExistingSalesmanByCodeRow) *Salesman {
	return &Salesman{}
}

func DbSalesmanByManagerToSalesman(db *db.GetSalesmenByManagerRow) *Salesman {
	return &Salesman{}
}

func DbExistingSalesmanByManagerToSalesman(db *db.GetExistingSalesmenByManagerRow) *Salesman {
	return &Salesman{}
}

func mapToSalesman() *Salesman {
	return &Salesman{}
}

func CreateSalesmanToDb(r *CreateSalesmanRequest) *db.CreateSalesmanParams {
	return &db.CreateSalesmanParams{}
}

func UpdateSalesmanToDb(r *UpdateSalesmanRequest) *db.UpdateSalesmanParams {
	return &db.UpdateSalesmanParams{}
}
