package util

import "strings"

// EnvMode specifies if we will be using strict env vars
type EnvMode string

const (
	// Infer - infer environment variable constraints from turbo.json
	Infer EnvMode = "Infer"
	// Loose - environment variables are unconstrained
	Loose EnvMode = "Loose"
	// StrictIncludeFrameworkVars - environment variables are limited.
	// They include vars from detected frameworks.
	StrictIncludeFrameworkVars EnvMode = "StrictIncludeFrameworkVars"
	// Strict - environment variables are limited
	Strict EnvMode = "Strict"
)

// MarshalText implements TextMarshaler for the struct.
func (em EnvMode) MarshalText() (text []byte, err error) {
	if em == StrictIncludeFrameworkVars {
		return []byte("strict-include-framework-vars"), nil
	}
	return []byte(strings.ToLower(string(em))), nil
}

// IsStrict collapses the two Strict variants.
func (em EnvMode) IsStrict() bool {
	return em == Strict || em == StrictIncludeFrameworkVars
}

// RunOpts holds the options that control the execution of a turbo run
type RunOpts struct {
	// Force execution to be serially one-at-a-time
	Concurrency int
	// Whether to execute in parallel (defaults to false)
	Parallel bool

	EnvMode EnvMode
	// The filename to write a perf profile.
	Profile string
	// If true, continue task executions even if a task fails.
	ContinueOnError bool
	PassThroughArgs []string
	// Restrict execution to only the listed task names. Default false
	Only bool
	// Dry run flags
	DryRun     bool
	DryRunJSON bool
	// Graph flags
	GraphDot      bool
	GraphFile     string
	NoDaemon      bool
	SinglePackage bool

	// logPrefix controls whether we should print a prefix in task logs
	LogPrefix string

	// Whether turbo should create a run summary
	Summarize bool

	ExperimentalSpaceID string
}
