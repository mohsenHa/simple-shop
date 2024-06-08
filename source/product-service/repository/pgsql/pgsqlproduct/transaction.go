package pgsqlproduct

import (
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"context"
	"database/sql"
)

func (d *DB) BeginTx(ctx context.Context, id int) (*sql.Tx, error) {
	const op = "pgsqlad.BeginTx"

	tx, err := d.conn.Conn().BeginTx(ctx, nil)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return tx, nil
}
