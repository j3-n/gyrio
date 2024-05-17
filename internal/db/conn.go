package db

type Connector interface {
	Conn(...interface{}) (*DB, error)
}

var (
	EmptyConn  = emptyConn{}
	SQLiteConn = sqliteConn{}
)
