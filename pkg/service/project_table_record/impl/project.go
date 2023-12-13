package table_record_impl

import (
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

func (s *tableRecordService) ListTableRecords(p *table_record.GetTableListParams) ([]model.TableWithRelationCount, error) {
	recordList, err := s.dataStore.ListTableRecords(p.Context, p.ProjectId)
	return recordList, err
}

func (s *tableRecordService) GetTableRecordById(p *table_record.GetTableByIdParams) (*model.TableWithRelationCount, error) {
	record, err := s.dataStore.GetTableRecordById(p.Context, p.ProjectId, p.TableId)
	return record, err
}

func (s *tableRecordService) CreateTable(p *table_record.CreateTableParams) (*model.Table, error) {
	record, err := s.dataStore.CreateTable(p.Context, p.ProjectId, p.Table)
	return record, err
}

func (s *tableRecordService) CreateBulkTable(p *table_record.CreateBulkTableParams) (int64, error) {
	record, err := s.dataStore.CreateBulkTable(p.Context, p.ProjectId, p.TableList)
	return record, err
}

func (s *tableRecordService) UpdateTableById(p *table_record.UpdateTableByIdParams) error {
	err := s.dataStore.UpdateTableById(p.Context, p.TableId, p.Table)
	return err
}

func (s *tableRecordService) DeleteTableById(p *table_record.DeleteTableByIdParams) error {
	err := s.dataStore.DeleteTableById(p.Context, p.TableId)
	return err
}
