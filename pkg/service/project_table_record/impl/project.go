package table_record_impl

import (
	"errors"

	"github.com/quarkloop/quarkloop/pkg/model"
	table_record "github.com/quarkloop/quarkloop/pkg/service/project_table_record"
	"github.com/quarkloop/quarkloop/pkg/store/repository"
)

type tableRecordService struct {
	dataStore *repository.Repository
}

func NewTableRecordService(ds *repository.Repository) table_record.Service {
	return &tableRecordService{
		dataStore: ds,
	}
}

func (s *tableRecordService) ListTableRecords(p *table_record.GetTableRecordListParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		recordList, err := s.dataStore.ListMainRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "document":
		recordList, err := s.dataStore.ListDocumentRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "form":
		recordList, err := s.dataStore.ListFormRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	case "payment":
		recordList, err := s.dataStore.ListPaymentRecords(p.Context, p.ProjectId, p.BranchId)
		return recordList, err
	}

	return nil, errors.New("[ListTableRecords] table type does not match")
}

func (s *tableRecordService) GetTableRecordById(p *table_record.GetTableRecordByIdParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.dataStore.GetMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "document":
		record, err := s.dataStore.GetDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "form":
		record, err := s.dataStore.GetFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	case "payment":
		record, err := s.dataStore.GetPaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return record, err
	}

	return nil, errors.New("[GetTableRecordById] table type does not match")
}

func (s *tableRecordService) CreateTableRecord(p *table_record.CreateTableRecordParams) (interface{}, error) {
	switch p.TableType {
	case "main":
		record, err := s.dataStore.CreateMainRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*model.MainRecord))
		return record, err
	case "document":
		record, err := s.dataStore.CreateDocumentRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*model.DocumentRecord))
		return record, err
	case "form":
		record, err := s.dataStore.CreateFormRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*model.FormRecord))
		return record, err
	case "payment":
		record, err := s.dataStore.CreatePaymentRecord(p.Context, p.ProjectId, p.BranchId, p.Record.(*model.PaymentRecord))
		return record, err
	}

	return nil, errors.New("[CreateTable] table type does not match")
}

func (s *tableRecordService) UpdateTableRecordById(p *table_record.UpdateTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.dataStore.UpdateMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*model.MainRecord))
		return err
	case "document":
		err := s.dataStore.UpdateDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*model.DocumentRecord))
		return err
	case "form":
		err := s.dataStore.UpdateFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*model.FormRecord))
		return err
	case "payment":
		err := s.dataStore.UpdatePaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId, p.Record.(*model.PaymentRecord))
		return err
	}

	return errors.New("[UpdateTableById] table type does not match")
}

func (s *tableRecordService) DeleteTableRecordById(p *table_record.DeleteTableRecordByIdParams) error {
	switch p.TableType {
	case "main":
		err := s.dataStore.DeleteMainRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "document":
		err := s.dataStore.DeleteDocumentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "form":
		err := s.dataStore.DeleteFormRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	case "payment":
		err := s.dataStore.DeletePaymentRecordById(p.Context, p.ProjectId, p.BranchId, p.RecordId)
		return err
	}

	return errors.New("[DeleteTableById] table type does not match")
}
