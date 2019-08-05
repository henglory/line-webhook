package conf

import "os"

var (
	ServicePort = os.Getenv("SERVICE_PORT")
)
