package typeAnalyzer

// NoMainPkgError represents a no main package error
type NoMainPkgError struct {
}

func (e *NoMainPkgError) Error() string {
	return "No main functions found in runner.PkgPath"
}
