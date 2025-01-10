package data

import (
	"context"

	"github.com/JueViGrace/clo-backend/internal/db"
	"github.com/JueViGrace/clo-backend/internal/types"
)

type DocumentStore interface {
	GetDocuments() (documents []types.DocumentResponse, err error)
	GetDocumentsWithLines() (documents []types.DocumentWithLinesResponse, err error)
	GetDocumentByCode(code string) (document *types.DocumentResponse, err error)
	GetDocumentByCodeWithLines(code string) (document *types.DocumentWithLinesResponse, err error)
	GetDocumentsByManager(code string) (documents []types.DocumentResponse, err error)
	GetDocumentsBySalesman(code string) (documents []types.DocumentResponse, err error)
	GetDocumentsByCustomer(code string) (documents []types.DocumentResponse, err error)
	CreateDocument(r *types.CreateDocumentRequest) (document *types.DocumentWithLinesResponse, err error)
	UpdateDocument(r *types.UpdateDocumentRequest) (document *types.DocumentWithLinesResponse, err error)
}

func (s *storage) DocumentStore() DocumentStore {
	return NewDocumentStore(s.ctx, s.queries)
}

type documentStore struct {
	ctx context.Context
	db  *db.Queries
}

func NewDocumentStore(ctx context.Context, db *db.Queries) DocumentStore {
	return &documentStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *documentStore) GetDocuments() ([]types.DocumentResponse, error) {
	res := make([]types.DocumentResponse, 0)

	dbDocs, err := s.db.GetDocuments(s.ctx)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		res = append(res, *types.DbDocToDocument(&dbDoc))
	}

	return res, nil
}

func (s *documentStore) GetDocumentsWithLines() ([]types.DocumentWithLinesResponse, error) {
	dbDocuments, err := s.db.GetDocumentsWithLines(s.ctx)
	if err != nil {
		return nil, err
	}

	return types.GroupDocWithLinesRows(dbDocuments), nil
}

func (s *documentStore) GetDocumentByCode(code string) (*types.DocumentResponse, error) {
	dbDocument, err := s.db.GetDocumentByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.DbDocToDocument(&dbDocument), nil
}

func (s *documentStore) GetDocumentByCodeWithLines(code string) (*types.DocumentWithLinesResponse, error) {
	dbDocuments, err := s.db.GetDocumentWithLinesByCode(s.ctx, code)
	if err != nil {
		return nil, err
	}

	return types.GroupDocWithLinesByCodeRows(dbDocuments), nil
}

func (s *documentStore) GetDocumentsByManager(code string) ([]types.DocumentResponse, error) {
	docs := make([]types.DocumentResponse, 0)

	dbDocs, err := s.db.GetDocumentsByManager(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		docs = append(docs, *types.DbDocToDocument(&dbDoc))
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsBySalesman(code string) ([]types.DocumentResponse, error) {
	docs := make([]types.DocumentResponse, 0)

	dbDocs, err := s.db.GetDocumentsBySalesman(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		docs = append(docs, *types.DbDocToDocument(&dbDoc))
	}

	return docs, nil
}

func (s *documentStore) GetDocumentsByCustomer(code string) ([]types.DocumentResponse, error) {
	docs := make([]types.DocumentResponse, 0)

	dbDocs, err := s.db.GetDocumentsByCustomer(s.ctx, code)
	if err != nil {
		return nil, err
	}

	for _, dbDoc := range dbDocs {
		docs = append(docs, *types.DbDocToDocument(&dbDoc))
	}

	return docs, nil
}

func (s *documentStore) CreateDocument(r *types.CreateDocumentRequest) (*types.DocumentWithLinesResponse, error) {
	lines := make([]types.DocumentLineResponse, 0)

	dbDoc, err := s.db.CreateDocument(s.ctx, *types.CreateDocumentToDb(r))
	if err != nil {
		return nil, err
	}

	head := types.DbDocToDocument(&dbDoc)

	for _, line := range r.Lines {
		dbLine, err := s.db.CreateDocumentLine(s.ctx, *types.CreateDocumentLineToDb(&line))
		if err != nil {
			continue
		}

		lines = append(lines, *types.DbDocLineToDocLine(&dbLine))
	}

	return types.MapToDocWithLines(head, lines), nil
}

func (s *documentStore) UpdateDocument(r *types.UpdateDocumentRequest) (*types.DocumentWithLinesResponse, error) {
	lines := make([]types.DocumentLineResponse, 0)

	dbDoc, err := s.db.UpdateDocument(s.ctx, *types.UpdateDocumentToDb(r))
	if err != nil {
		return nil, err
	}

	head := types.DbDocToDocument(&dbDoc)

	for _, line := range r.Lines {
		dbLine, _ := s.db.UpdateDocumentLine(s.ctx, *types.UpdateDocumentLineToDb(&line))
		if err != nil {
			continue
		}

		lines = append(lines, *types.DbDocLineToDocLine(&dbLine))
	}

	return types.MapToDocWithLines(head, lines), nil
}
