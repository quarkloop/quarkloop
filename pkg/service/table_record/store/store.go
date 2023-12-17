package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

type TableRecordStore interface {
	// main
	ListMainRecords(ctx context.Context, projectId int, branchId int) ([]table_record.MainRecordWithRelationCount, error)
	GetMainRecordById(ctx context.Context, projectId int, branchId int, mainId string) (*table_record.MainRecordWithRelationCount, error)
	CreateMainRecord(ctx context.Context, projectId int, branchId int, table *table_record.MainRecord) (*table_record.MainRecord, error)
	CreateBulkMainRecords(ctx context.Context, projectId int, branchId int, tableList []table_record.MainRecord) (int64, error)
	UpdateMainRecordById(ctx context.Context, projectId int, branchId int, mainId string, table *table_record.MainRecord) error
	DeleteMainRecordById(ctx context.Context, projectId int, branchId int, mainId string) error

	// document
	ListDocumentRecords(ctx context.Context, projectId int, branchId int) ([]table_record.DocumentRecord, error)
	GetDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) (*table_record.DocumentRecord, error)
	CreateDocumentRecord(ctx context.Context, projectId int, branchId int, doc *table_record.DocumentRecord) (*table_record.DocumentRecord, error)
	UpdateDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string, doc *table_record.DocumentRecord) error
	DeleteDocumentRecordById(ctx context.Context, projectId int, branchId int, documentId string) error

	// payment
	ListPaymentRecords(ctx context.Context, projectId int, branchId int) ([]table_record.PaymentRecord, error)
	GetPaymentRecordById(ctx context.Context, projectId int, branchId int, documentId string) (*table_record.PaymentRecord, error)
	CreatePaymentRecord(ctx context.Context, projectId int, branchId int, doc *table_record.PaymentRecord) (*table_record.PaymentRecord, error)
	UpdatePaymentRecordById(ctx context.Context, projectId int, branchId int, documentId string, doc *table_record.PaymentRecord) error
	DeletePaymentRecordById(ctx context.Context, projectId int, branchId int, documentId string) error

	// form
	ListFormRecords(ctx context.Context, projectId int, branchId int) ([]table_record.FormRecord, error)
	GetFormRecordById(ctx context.Context, projectId int, branchId int, documentId string) (*table_record.FormRecord, error)
	CreateFormRecord(ctx context.Context, projectId int, branchId int, doc *table_record.FormRecord) (*table_record.FormRecord, error)
	UpdateFormRecordById(ctx context.Context, projectId int, branchId int, documentId string, doc *table_record.FormRecord) error
	DeleteFormRecordById(ctx context.Context, projectId int, branchId int, documentId string) error
}

type tableRecordStore struct {
	Conn *pgx.Conn
}

func NewTableRecordStore(conn *pgx.Conn) *tableRecordStore {
	return &tableRecordStore{
		Conn: conn,
	}
}
