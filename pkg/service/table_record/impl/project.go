package table_record_impl

import (
	"context"
	"errors"

	"github.com/quarkloop/quarkloop/pkg/service/table_record"
	"github.com/quarkloop/quarkloop/pkg/service/table_record/store"
)

type tableRecordService struct {
	store store.TableRecordStore
}

func NewTableRecordService(ds store.TableRecordStore) table_record.Service {
	return &tableRecordService{
		store: ds,
	}
}

func (s *tableRecordService) ListTableRecords(ctx context.Context, p *table_record.GetTableRecordListParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		recordList, err := s.store.ListMainRecords(ctx, p.ProjectId, p.BranchId)
		return recordList, err
	case "document":
		recordList, err := s.store.ListDocumentRecords(ctx, p.ProjectId, p.BranchId)
		return recordList, err
	case "form":
		recordList, err := s.store.ListFormRecords(ctx, p.ProjectId, p.BranchId)
		return recordList, err
	case "payment":
		recordList, err := s.store.ListPaymentRecords(ctx, p.ProjectId, p.BranchId)
		return recordList, err
	}

	return nil, errors.New("[ListTableRecords] table type does not match")
}

func (s *tableRecordService) GetTableRecordById(ctx context.Context, p *table_record.GetTableRecordByIdParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.store.GetMainRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "document":
		record, err := s.store.GetDocumentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "form":
		record, err := s.store.GetFormRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "payment":
		record, err := s.store.GetPaymentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	}

	return nil, errors.New("[GetTableRecordById] table type does not match")
}

func (s *tableRecordService) CreateTableRecord(ctx context.Context, p *table_record.CreateTableRecordParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.store.CreateMainRecord(ctx, p.ProjectId, p.BranchId, p.Record.(*table_record.MainRecord))
		return record, err
	case "document":
		record, err := s.store.CreateDocumentRecord(ctx, p.ProjectId, p.BranchId, p.Record.(*table_record.DocumentRecord))
		return record, err
	case "form":
		record, err := s.store.CreateFormRecord(ctx, p.ProjectId, p.BranchId, p.Record.(*table_record.FormRecord))
		return record, err
	case "payment":
		record, err := s.store.CreatePaymentRecord(ctx, p.ProjectId, p.BranchId, p.Record.(*table_record.PaymentRecord))
		return record, err
	}

	return nil, errors.New("[CreateTable] table type does not match")
}

func (s *tableRecordService) UpdateTableRecordById(ctx context.Context, p *table_record.UpdateTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.store.UpdateMainRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.MainRecord))
		return err
	case "document":
		err := s.store.UpdateDocumentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.DocumentRecord))
		return err
	case "form":
		err := s.store.UpdateFormRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.FormRecord))
		return err
	case "payment":
		err := s.store.UpdatePaymentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.PaymentRecord))
		return err
	}

	return errors.New("[UpdateTableById] table type does not match")
}

func (s *tableRecordService) DeleteTableRecordById(ctx context.Context, p *table_record.DeleteTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.store.DeleteMainRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "document":
		err := s.store.DeleteDocumentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "form":
		err := s.store.DeleteFormRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "payment":
		err := s.store.DeletePaymentRecordById(ctx, p.ProjectId, p.BranchId, p.RecordId)
		return err
	}

	return errors.New("[DeleteTableById] table type does not match")
}
