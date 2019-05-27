package db

const (
	TYPE_MYSQL = "mysql"
)

type DB struct {
	host string
	port string
}

func GetInstance(t string) *DB {
	return &DB{
		"102",
		"1",
	}
}
