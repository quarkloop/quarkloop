package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/quarkloop/quarkloop/pkg/model"
)

/// ListPaymentRecords

const listPaymentRecordsQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TablePayment"
  WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId;
`

func (r *Repository) ListPaymentRecords(ctx context.Context, projectId int, branchId int) ([]model.PaymentRecord, error) {
	rows, err := r.ProjectDbConn.Query(ctx, listPaymentRecordsQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[LIST] failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var instanceList []model.PaymentRecord

	for rows.Next() {
		var payment model.PaymentRecord
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

		instanceList = append(instanceList, payment)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "[LIST]: Rows failed: %v\n", err)
		return nil, err
	}

	return instanceList, nil
}

/// GetPaymentRecordById

const getPaymentRecordByIdQuery = `
SELECT
  "id", "name", "description", "metadata", "data", "createdAt", "createdBy", "updatedAt", "updatedBy"
FROM
  "project"."TablePayment"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) GetPaymentRecordById(ctx context.Context, projectId int, branchId int, paymentId string) (*model.PaymentRecord, error) {
	row := r.ProjectDbConn.QueryRow(ctx, getPaymentRecordByIdQuery, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        paymentId,
	})

	var payment model.PaymentRecord
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
