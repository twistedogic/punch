package internal

type ExecMode string

const (
	PARALLEL ExecMode = "parallel"
	SERIAL   ExecMode = "serial"
)

type Config struct {
	Mode    ExecMode `yaml:"mode"`
	Metrics *Scrape  `yaml:"metrics_config,omitempty"`
	Profile *Scrape  `yaml:"profile_config,omitempty"`
	Targets []Target `yaml:"targets"`
}

type Scrape struct {
	Target   *URL     `yaml:"target_url"`
	Interval Duration `yaml:"interval"`
}

// Target defines the test scenerio
// It can end by specifying an end duration or
// error rate threshold is reached
type Target struct {
	TargetURL          *URL     `yaml:"target_url"`
	Method             string   `yaml:"method"`
	Payload            *Bytes   `yaml:"payload,omitempty"`
	ExpectedStatusCode *int     `yaml:"expected_status,omitempty"`
	Timeout            Duration `yaml:"timeout"`
	Duration           Duration `yaml:"duration"`
	Threshold          *float64 `yaml:"threshold,omitempty"`
	Bucket             []int    `yaml:"bucket,omitempty"`
}
