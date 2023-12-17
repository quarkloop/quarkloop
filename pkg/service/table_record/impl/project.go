package table_record_impl

import (
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

func (s *tableRecordService) ListTableRecords(p *table_record.GetTableRecordListParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		recordList, err := s.store.ListMainRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "document":
		recordList, err := s.store.ListDocumentRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "form":
		recordList, err := s.store.ListFormRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "payment":
		recordList, err := s.store.ListPaymentRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	}

	return nil, errors.New("[ListTableRecords] table type does not match")
}

func (s *tableRecordService) GetTableRecordById(p *table_record.GetTableRecordByIdParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.store.GetMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "document":
		record, err := s.store.GetDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "form":
		record, err := s.store.GetFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "payment":
		record, err := s.store.GetPaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	}

	return nil, errors.New("[GetTableRecordById] table type does not match")
}

func (s *tableRecordService) CreateTableRecord(p *table_record.CreateTableRecordParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.store.CreateMainRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*table_record.MainRecord))
		return record, err
	case "document":
		record, err := s.store.CreateDocumentRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*table_record.DocumentRecord))
		return record, err
	case "form":
		record, err := s.store.CreateFormRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*table_record.FormRecord))
		return record, err
	case "payment":
		record, err := s.store.CreatePaymentRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*table_record.PaymentRecord))
		return record, err
	}

	return nil, errors.New("[CreateTable] table type does not match")
}

func (s *tableRecordService) UpdateTableRecordById(p *table_record.UpdateTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.store.UpdateMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.MainRecord))
		return err
	case "document":
		err := s.store.UpdateDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.DocumentRecord))
		return err
	case "form":
		err := s.store.UpdateFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.FormRecord))
		return err
	case "payment":
		err := s.store.UpdatePaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*table_record.PaymentRecord))
		return err
	}

	return errors.New("[UpdateTableById] table type does not match")
}

func (s *tableRecordService) DeleteTableRecordById(p *table_record.DeleteTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.store.DeleteMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "document":
		err := s.store.DeleteDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "form":
		err := s.store.DeleteFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "payment":
		err := s.store.DeletePaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	}

	return errors.New("[DeleteTableById] table type does not match")
}
