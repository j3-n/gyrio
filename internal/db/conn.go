package db

type Connector interface {
	Conn(...interface{}) (*DB, error)
}

var (
	SQLiteConn = sqliteConn{}
)
