package database

import "github.com/render-examples/go-gin-web-server/infrastructure"

type DataBase struct {
	Conn infrastructure.DBConn
}

func NewDataBase(conn infrastructure.DBConn) *DataBase {
	return &DataBase{
		Conn: conn,
	}
}
