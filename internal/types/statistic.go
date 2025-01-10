package types

import "github.com/JueViGrace/clo-backend/internal/db"

type SalesmanStatistic struct{}

type CreateStatisticRequest struct{}

type UpdateStatisticRequest struct{}

func DbStatisticToStatistc(db *db.ClossSalesmanStatistic) *SalesmanStatistic {
	return &SalesmanStatistic{}
}

func CreateStatisticToDb(r *CreateStatisticRequest) *db.CreateStatisticParams {
	return &db.CreateStatisticParams{}
}

func UpdateStatisticToDb(r *UpdateStatisticRequest) *db.UpdateStatisticParams {
	return &db.UpdateStatisticParams{}
}
