package plan

type (
	// GraphQLPlan is for testing a service with GraphQL API
	GraphQLPlan struct {
		Name       string        `yaml:"name"`
		Address    string        `yaml:"address"`
		CAChain    string        `yaml:"ca_chain"`
		ClientCert string        `yaml:"client_cert"`
		ClientKey  string        `yaml:"client_key"`
		Tests      []GraphQLTest `yaml:"tests"`
	}

	// GraphQLTest defines test spec for a GraphQL service
	GraphQLTest struct {
		Expect GraphQLExpect `yaml:"expect"`
		Report GraphQLReport `yaml:"-"`
	}

	// GraphQLExpect defines expectations for a GraphQL test
	GraphQLExpect struct {
	}

	// GraphQLReport defines results for a GraphQL test
	GraphQLReport struct {
		ResponseTime int
	}
)
