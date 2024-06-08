package pgsql

type Scanner interface {
	Scan(dest ...any) error
}
