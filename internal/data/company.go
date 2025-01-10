package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type CompanyStore interface {
	GetCompanies() (companies []types.CompanyResponse, err error)
	GetCompanyByCode(code string) (company *types.CompanyResponse, err error)
	GetExistingCompanyByCode(code string) (company *types.CompanyResponse, err error)
	CreateCompany(r *types.CreateCompanyRequest) (company *types.CompanyResponse, err error)
	UpdateCompany(r *types.UpdateCompanyRequest) (company *types.CompanyResponse, err error)
	SoftDeleteCompany(code string) (err error)
	DeleteCompany(code string) (err error)
}

func (s *storage) CompanyStore() CompanyStore {
	return NewCompanyStore(s.ctx, s.queries)
}

type companyStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewCompanyStore(ctx context.Context, db *db.Queries) CompanyStore {
	return &companyStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *companyStore) GetCompanies() ([]types.CompanyResponse, error) {
	companies := make([]types.CompanyResponse, 0)

	dbCompanies, err := s.db.GetCompanies(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbCompany := range dbCompanies {
		companies = append(companies, *types.DbComanyToCompany(&dbCompany))
	}

	return companies, nil
}

func (s *companyStore) GetCompanyByCode(code string) (*types.CompanyResponse, error) {
	dbCompany, err := s.db.GetCompanyByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbComanyToCompany(&dbCompany), nil
}

func (s *companyStore) GetExistingCompanyByCode(code string) (*types.CompanyResponse, error) {
	dbCompany, err := s.db.GetExistingCompanyByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbComanyToCompany(&dbCompany), nil
}

func (s *companyStore) CreateCompany(r *types.CreateCompanyRequest) (*types.CompanyResponse, error) {
	dbCompany, err := s.db.CreateCompany(s.ctx, *types.CreateCompanyToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbComanyToCompany(&dbCompany), nil
}

func (s *companyStore) UpdateCompany(r *types.UpdateCompanyRequest) (*types.CompanyResponse, error) {
	dbCompany, err := s.db.UpdateCompany(s.ctx, *types.UpdateCompanyToDb(r))
	if err != nil {
		return nil, err
	}

	return types.DbComanyToCompany(&dbCompany), nil
}

func (s *companyStore) SoftDeleteCompany(code string) error {
	err := s.db.SoftDeleteCompany(s.ctx, db.SoftDeleteCompanyParams{
		UpdatedAt: time.Now().String(),
		DeletedAt: sql.NullString{
			String: time.Now().String(),
			Valid:  true,
		},
		KedCodigo: code,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *companyStore) DeleteCompany(code string) error {
	err := s.db.DeleteCompany(s.ctx, code)
	if err != nil {
		return err
	}

	return nil
}
