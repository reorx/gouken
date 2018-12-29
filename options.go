package gouken

/*
Options:
- Name
- Host
- Port
- LogLevel
- Debug
*/

// Option ...
type Option func(*application)

// Name ...
func Name(n string) Option {
	return func(a *application) {
		a.Name = n
	}
}

// Host ...
func Host(h string) Option {
	return func(a *application) {
		a.Host = h
	}
}

// Port ...
func Port(p int) Option {
	return func(a *application) {
		a.Port = p
	}
}

// LogLevel ...
func LogLevel(l string) Option {
	return func(a *application) {
		a.LogLevel = l
	}
}

// ClientAddress ..
func ClientAddress(s string) Option {
	return func(a *application) {
		a.ClientAddress = s
	}
}

// LogFilename ..
func LogFilename(l bool) Option {
	return func(a *application) {
		a.LogFilename = l
	}
}

// LogRequest ..
func LogRequest(l bool) Option {
	return func(a *application) {
		a.LogRequest = l
	}
}

// LogResponse ..
func LogResponse(l bool) Option {
	return func(a *application) {
		a.LogResponse = l
	}
}

// Debug ...
func Debug(d bool) Option {
	return func(a *application) {
		a.Debug = d
	}
}
