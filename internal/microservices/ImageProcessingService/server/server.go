package server

import (
	"flag"
	pb "paint/internal/microservices/ImageProcessingService/proto"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 8080, "The server port")
)

type calculationServer struct {
	pb.CalculationClient
}

/*
func (s *calculationServer) GetCalCalculationParams(ctx context.Context) (*pb.CalculationParams, error) {
	return pb.CalculationParams{Message: ""}
}
 */