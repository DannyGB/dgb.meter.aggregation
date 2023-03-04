package configuration

import "os"

type Configuration struct {
	ELEC_BACKEND string
	OIL_BACKEND  string
	HTTP_PORT    string
	CORS_CLIENTS string
	ENV          string
}

func NewConfig() Configuration {

	configuration := &Configuration{}
	configuration.CORS_CLIENTS = os.Getenv("AGGREGATION_CORS_CLIENTS")
	configuration.HTTP_PORT = os.Getenv("AGGREGATION_HTTP_PORT")
	configuration.ELEC_BACKEND = os.Getenv("AGGREGATION_ELEC_BACKEND")
	configuration.OIL_BACKEND = os.Getenv("AGGREGATION_OIL_BACKEND")
	configuration.ENV = os.Getenv("AGGREGATION_READINGS_ENV")

	return *configuration
}
