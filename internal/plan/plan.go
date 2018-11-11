package plan

type (
	// TestPlan defines all tests to be run
	TestPlan struct {
		Name         string        `yaml:"name"`
		RESTPlans    []RESTPlan    `yaml:"rest"`
		GRPCPlans    []GRPCPlan    `yaml:"grpc"`
		NATSPlans    []NATSPlan    `yaml:"nats"`
		GraphQLPlans []GraphQLPlan `yaml:"graphql"`
	}
)
