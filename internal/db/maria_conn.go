package db

type mariaConn struct{}

// Maria is a drop in replacement for MySQL,
// so we can just use the connector from MySQL.
func (c mariaConn) Conn(args ...interface{}) (*DB, error) {
	conn := mysqlConn{}
	return conn.Conn(args...)
}
