package run

import (
	"os"

	"github.com/moorara/chia/internal/plan"
	yaml "gopkg.in/yaml.v2"
)

type (
	// TestDriver defines the interface for executing a test plan
	TestDriver interface {
		Run() error
	}

	testDriver struct {
		Plan *plan.TestPlan
	}
)

// NewTestDriver creates a new instance of TestDriver
func NewTestDriver(file string) (TestDriver, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	plan := new(plan.TestPlan)
	err = yaml.NewDecoder(f).Decode(plan)
	if err != nil {
		return nil, err
	}

	return &testDriver{
		Plan: plan,
	}, nil
}

func (d *testDriver) Run() error {
	return nil
}
