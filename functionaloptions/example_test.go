package functionaloptions_test

import o "github.com/nujz/grpc-enigma/functionaloptions"

func Example_newServer() {
	server := o.NewServer(o.WithHost("localhost"), o.WithPort(80))
	server.Serve()
}
