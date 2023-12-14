package table_record

type Service interface {
	ListTableRecords(*GetTableRecordListParams) (interface{}, error)
	GetTableRecordById(*GetTableRecordByIdParams) (interface{}, error)
	CreateTableRecord(*CreateTableRecordParams) (interface{}, error)
	UpdateTableRecordById(*UpdateTableRecordByIdParams) error
	DeleteTableRecordById(*DeleteTableRecordByIdParams) error
}
