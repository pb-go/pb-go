package databaseop


// DBClient : General Database client interface, reserved for future adopt.
type DBClient interface {
	connNCheck(dbCliOption interface{}) error
	itemCreate(inputdata interface{}) error
	itemUpdate(filter1 interface{}, change1 interface{}) error
	itemDelete(filter1 interface{}) error
	itemRead(filter1 interface{}) (UserData, error)
}
