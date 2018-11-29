package conf


type Mysql interface {
	DbUser()
	DbHost()
	DbPort()
	DbDatabase()
	DbUserName()
	DBPassword()
}


type f interface {
	Mysql()
}

type Sql struct {
}

func (sql *Sql) DbUser()  {
}

func (sql *Sql) DbHost() string {
	return DBHOST
}

func (sql *Sql) DbPort() string {
	return DBPORT
}

func (sql *Sql) DbDatabase() {
}

func (sql *Sql) DbUserName() {
}

func (sql *Sql) DBPassword() {
}

