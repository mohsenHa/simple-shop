package pgsqlproduct

import (
	"clean-code-structure/entity"
	"clean-code-structure/param/productparam"
	"clean-code-structure/pkg/errmsg"
	"clean-code-structure/pkg/richerror"
	"clean-code-structure/repository/pgsql"
	"context"
	"database/sql"
	"time"
)

func (d *DB) GetProductWithId(ctx context.Context, id int) (productparam.ProductRepo, error) {
	const op = "pgsqlproduct.GetProductWithId"
	row := d.conn.Conn().QueryRowContext(ctx, `SELECT * FROM products WHERE Id = $1`, id)
	product, err := scanProduct(row)
	if err != nil {
		return productparam.ProductRepo{}, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return product, nil
}

func (d *DB) IsProductExist(ctx context.Context, id int) (bool, error) {
	const op = "pgsqlproduct.IsProductExist"

	row := d.conn.Conn().QueryRowContext(ctx, `SELECT * FROM products WHERE ID = $1`, id)
	_, err := scanProduct(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return true, nil
}

func (d *DB) GetTotalProductCount(ctx context.Context, filters ...entity.Filter) (int, error) {
	const op = "pgsqlproduct.GetTotalProductCount"

	query, values := d.conn.MakeQueryWithFilters(filters)
	if query != "" {
		query = " WHERE " + query
	}
	row := d.conn.Conn().QueryRowContext(ctx, `SELECT COUNT(*) as total FROM products `+query, values...)
	var total int
	err := row.Scan(&total)

	if err != nil {
		return 0, richerror.New(op).WithErr(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return total, nil
}

func scanProduct(scanner pgsql.Scanner) (productparam.ProductRepo, error) {
	var createdAt time.Time
	var updatedAt time.Time
	var deletedAt sql.NullTime
	var product entity.Product

	err := scanner.Scan(&product.Id, &product.Name, &product.Quantity, &createdAt, &updatedAt, &deletedAt)

	return productparam.ProductRepo{
		Entity:    product,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}, err
}
