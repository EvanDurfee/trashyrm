package system

const (
	defaultPromptCutoff = 2
	defaultRecycleHome  = false
)

type Config struct {
	PromptCutoff *int  `yaml:"promptCutoff,omitempty"`
	RecycleHome  *bool `yaml:"recycleHome,omitempty"`
}

// Placeholder for defaults until parsing is in place
func NewConfig() *Config {
	a, b := defaultPromptCutoff, defaultRecycleHome
	return &Config{
		PromptCutoff: &a,
		RecycleHome:  &b,
	}
}

// TODO: parser
// https://pkg.go.dev/gopkg.in/yaml.v2?utm_source=godoc

// TODO: trash path? or home only?
