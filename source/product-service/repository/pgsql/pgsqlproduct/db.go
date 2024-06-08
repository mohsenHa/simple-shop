package pgsqlproduct

import "clean-code-structure/repository/pgsql"

type DB struct {
	conn *pgsql.PGSQLDB
}

func New(conn *pgsql.PGSQLDB) *DB {
	return &DB{conn}
}
