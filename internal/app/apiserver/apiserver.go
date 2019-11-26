package apiserver

import "gopkg.in/mgo.v2"

//Start ...
func Start(config Config) error {

	return nil
}

func newDB(dbURL string) (*mgo.Database, error) {
	return &mgo.Database{}, nil
}
