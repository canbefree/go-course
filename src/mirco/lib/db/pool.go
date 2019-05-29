package db

// 连接池简单版本

type Pool struct {
}

//获取链接
func (pool *Pool) FetchConn() {

}

//使用完归还链接
func (pool *Pool) Return() {

}

/*

	var pool = new Pool(100)
	conn := fetchConn()
	Db.use(conn)
	Db.query()
	pool.Return(conn)

*/
