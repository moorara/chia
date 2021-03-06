package plan

type (
	// RESTPlan is for testing a service with RESTful API
	RESTPlan struct {
		Name       string            `yaml:"name"`
		Address    string            `yaml:"address"`
		CAChain    string            `yaml:"ca_chain"`
		ClientCert string            `yaml:"client_cert"`
		ClientKey  string            `yaml:"client_key"`
		Base       string            `yaml:"base"`
		Headers    map[string]string `yaml:"headers"`
		Tests      []RESTTest        `yaml:"tests"`
	}

	// RESTTest defines test spec for a REST service
	RESTTest struct {
		Name     string                 `yaml:"name"`
		Method   string                 `yaml:"method"`
		Endpoint string                 `yaml:"endpoint"`
		Headers  map[string]string      `yaml:"headers"`
		Body     map[string]interface{} `yaml:"body"`
		Expect   RESTExpect             `yaml:"expect"`
		Report   RESTReport             `yaml:"-"`
	}

	// RESTExpect defines expectations for a REST test
	RESTExpect struct {
		StatusCode int                    `yaml:"status_code"`
		Headers    map[string]string      `yaml:"headers"`
		Body       map[string]interface{} `yaml:"body"`
	}

	// RESTReport defines results for a REST test
	RESTReport struct {
		ResponseTime int
	}
)
