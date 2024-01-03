package database

import "github.com/render-examples/go-gin-web-server/infrastructure"

type DataBase struct {
	conn infrastructure.DBConn
}

func NewDataBase(conn infrastructure.DBConn) *DataBase {
	return &DataBase{
		conn: conn,
	}
}
