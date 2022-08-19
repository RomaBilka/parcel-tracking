package dependencies

import (
	"os"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	_, err := initConfig()
	assert.NoError(t, err)
}

func TestConfig(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Ok env",
			test: func(t *testing.T) {
				err := os.Setenv("NP_API_KEY", "test_key")
				assert.NoError(t, err)
				config, err := initConfig()
				assert.NoError(t, err)
				assert.Equal(t, "test_key", config.NovaPoshta.ApiKey)
			},
		},
		{
			name: "Ok flag",
			test: func(t *testing.T) {
				config := &Config{}
				parser := flags.NewParser(config, flags.Default&^flags.PrintErrors)
				_, err := parser.ParseArgs([]string{"--ME_PASSWORD=test_password"})
				assert.NoError(t, err)
				assert.Equal(t, "test_password", config.MeestExpress.Password)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.test(t)
		})
	}
}
