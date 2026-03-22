package driver

type Driver string

var (
	DriverPostgres Driver = "postgres"
	DriverMySql    Driver = "mysql"
)

func (driver Driver) Is(target Driver) bool {
	return driver == target
}
