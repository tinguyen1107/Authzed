package initializers

import (
	"os"

	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

var SpiceClient *authzed.Client

func ConnectToSpiceDb() {
	var err error
	systemCerts, err := grpcutil.WithSystemCerts(grpcutil.VerifyCA)
	if err != nil {
		panic("Unable to initialize system certs: " + err.Error())
	}
	SpiceClient, err = authzed.NewClient(
		"grpc.authzed.com:443",
		systemCerts,
		grpcutil.WithBearerToken(os.Getenv("SPICE_DB_TOKEN")),
	)
	if err != nil {
		panic("Unable to initialize system certs: " + err.Error())
	}
}
