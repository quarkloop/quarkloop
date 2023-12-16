package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/quarkloop/quarkloop/pkg/model"
)

/// CreatePaymentRecord

const createPaymentRecordMutation = `
INSERT INTO
  "project"."TablePayment" ("projectId", "branchId", "name", "description", "metadata", "data")
VALUES
  (@projectId, @branchId, @name, @description, @metadata, @data)
RETURNING
  "id", "name", "description", "metadata", "data", "createdAt";
`

func (r *Repository) CreatePaymentRecord(ctx context.Context, projectId int, branchId int, payment *model.PaymentRecord) (*model.PaymentRecord, error) {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		createPaymentRecordMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"name":        payment.Name,
			"description": payment.Description,
			"metadata":    payment.Metadata,
			"data":        payment.Data,
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", err)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to create")
		fmt.Fprintf(os.Stderr, "[CREATE] failed: %v\n", notFoundErr)
		return nil, notFoundErr
	}

	return payment, nil
}

/// UpdatePaymentRecordById

const updatePaymentRecordByIdMutation = `
UPDATE
  "project"."TablePayment"
SET
  "name"        = @name,
  "description" = @description,
  "metadata"    = @metadata,
  "data"        = @data,
  "updatedAt"   = @updatedAt
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) UpdatePaymentRecordById(ctx context.Context, projectId int, branchId int, paymentId string, payment *model.PaymentRecord) error {
	commandTag, err := r.ProjectDbConn.Exec(
		ctx,
		updatePaymentRecordByIdMutation,
		pgx.NamedArgs{
			"projectId":   projectId,
			"branchId":    branchId,
			"id":          paymentId,
			"name":        payment.Name,
			"description": payment.Description,
			"metadata":    payment.Metadata,
			"data":        payment.Data,
			"updatedAt":   time.Now(),
		},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to update")
		fmt.Fprintf(os.Stderr, "[UPDATE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}

/// DeletePaymentRecordById

const deletePaymentRecordByIdMutation = `
DELETE FROM
  "project"."TablePayment"
WHERE
  "projectId" = @projectId
AND
  "branchId" = @branchId
AND
  "id" = @id;
`

func (r *Repository) DeletePaymentRecordById(ctx context.Context, projectId int, branchId int, paymentId string) error {
	commandTag, err := r.ProjectDbConn.Exec(ctx, deletePaymentRecordByIdMutation, pgx.NamedArgs{
		"projectId": projectId,
		"branchId":  branchId,
		"id":        paymentId,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		notFoundErr := errors.New("cannot find to delete")
		fmt.Fprintf(os.Stderr, "[DELETE] failed: %v\n", notFoundErr)
		return notFoundErr
	}

	return nil
}
