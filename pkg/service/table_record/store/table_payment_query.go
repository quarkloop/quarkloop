package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/service/table_record"
)

/// ListPaymentRecords

const listPaymentRecordsQuery = `
SELECT 
	"id",
    "name",
    "description",
    "metadata",
    "data",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"project"."TablePayment"
WHERE 
	"projectId" = @projectId
AND 
	"branchId" = @branchId;
`

func (store *tableRecordStore) ListPaymentRecords(ctx context.Context, projectId int, branchId int) ([]table_record.PaymentRecord, error) {
	rows, err := store.Conn.Query(ctx, listPaymentRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var paymentList []table_record.PaymentRecord

	for rows.Next() {
		var payment table_record.PaymentRecord
		err := rows.Scan(
			&payment.Id,
			&payment.Name,
			&payment.Description,
			&payment.Metadata,
			&payment.Data,
			&payment.CreatedAt,
			&payment.CreatedBy,
			&payment.UpdatedAt,
			&payment.UpdatedBy,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
			return nil, err
		}

		paymentList = append(paymentList, payment)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return paymentList, nil
}

/// GetPaymentRecordById

const getPaymentRecordByIdQuery = `
SELECT 
	"id",
    "name",
    "description",
    "metadata",
    "data",
    "createdAt",
    "createdBy",
    "updatedAt",
    "updatedBy"
FROM 
	"project"."TablePayment"
WHERE 
	"projectId" = @projectId
AND 
	"branchId" = @branchId
AND 
	"id" = @id;
`

func (store *tableRecordStore) GetPaymentRecordById(ctx context.Context, projectId int, branchId int, paymentId string) (*table_record.PaymentRecord, error) {
	row := store.Conn.QueryRow(ctx, getPaymentRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        paymentId,
	})

	var payment table_record.PaymentRecord
	err := row.Scan(
		&payment.Id,
		&payment.Name,
		&payment.Description,
		&payment.Metadata,
		&payment.Data,
		&payment.CreatedAt,
		&payment.CreatedBy,
		&payment.UpdatedAt,
		&payment.UpdatedBy,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[READ] failed: %v\n", err)
		return nil, err
	}

	return &payment, nil
}
