package plan

type (
	// GRPCPlan is for testing a service with gRPC API
	GRPCPlan struct {
		Name       string     `yaml:"name"`
		Address    string     `yaml:"address"`
		CAChain    string     `yaml:"ca_chain"`
		ClientCert string     `yaml:"client_cert"`
		ClientKey  string     `yaml:"client_key"`
		Tests      []GRPCTest `yaml:"tests"`
	}

	// GRPCTest defines test spec for a GRPC service
	GRPCTest struct {
	}
)
