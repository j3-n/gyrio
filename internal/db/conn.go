package db

type Connector interface {
	Conn(...interface{}) (*DB, error)
}

var (
	EmptyConn  = emptyConn{}
	SQLiteConn = sqliteConn{}
	PgConn     = pgConn{}
	MySQLConn  = mysqlConn{}
	MariaConn  = mariaConn{}
)

func New(t DBType) Connector {
	switch t {
	case SQLite:
		return sqliteConn{}

	case Postgres:
		return pgConn{}

	case MySQL:
		return mysqlConn{}

	case Maria:
		return mariaConn{}

	default:
		return emptyConn{}
	}
}
