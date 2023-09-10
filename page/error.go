package page

import "errors"

var (
	Err400 = errors.New("bad request")
	Err401 = errors.New("unauthorised")
	Err404 = errors.New("not found")
	Err500 = errors.New("server errror")
)
