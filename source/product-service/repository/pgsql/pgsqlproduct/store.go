package pgsqlproduct

import (
	"clean-code-structure/entity"
	"clean-code-structure/logger"
	"clean-code-structure/pkg/richerror"
	"clean-code-structure/repository/pgsql"
	"context"
	"database/sql"
)

func (d *DB) StoreWitTransaction(ctx context.Context, transaction pgsql.Transaction,
	product entity.Product) (uint, error) {
	const op = "pgsql.StoreWitTransaction"

	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		return 0, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	query := `INSERT INTO products(name,quantity) VALUES ($1,$2) RETURNING id`
	stmt, err := tx.Prepare(query)
	if err != nil {
		return 0, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			logger.Logger.Error("Failed to close statement : " + err.Error())
		}
	}(stmt)

	var id uint
	err = stmt.QueryRow(product.Name, product.Quantity).Scan(&id)
	if err != nil {
		return 0, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	d.conn.WaitForTransaction(ctx, transaction, tx)

	return id, nil
}
