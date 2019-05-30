package pool

type Mysql struct {
}

func (mysql *Mysql) Exec() {

}

func (mysql *Mysql) Query() {

}

type Driver interface {
	Exec()
	Query()
}

type Pool struct {
	max     int    //最大连接数
	current int    //当前连接数
	Driver  string // mysqli pdo ?
	dbChan  chan Driver
}

func newPool(driver string) *Pool {
	return &Pool{
		Driver: driver,
	}
}

// func (pool *Pool) register

// func (pool *Pool) getConnect() Driver {
// 	return <-pool.dbChan
// }
