package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler(c *config) (*SqlHandler, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		c.host, c.user, c.password, c.database, c.port, c.timezone)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &SqlHandler{}, err
	}

	return &SqlHandler{Conn: conn}, nil
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
