package gouken

import (
	"github.com/spf13/viper"
)

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

// ConfName ...
func ConfName(k string) Option {
	return func(a *application) {
		a.Name = viper.GetString(k)
	}
}

// Host ...
func Host(h string) Option {
	return func(a *application) {
		a.Host = h
	}
}

// ConfHost ..
func ConfHost(k string) Option {
	return func(a *application) {
		a.Host = viper.GetString(k)
	}
}

// Port ...
func Port(p int) Option {
	return func(a *application) {
		a.Port = p
	}
}

// ConfPort ...
func ConfPort(k string) Option {
	return func(a *application) {
		a.Port = viper.GetInt(k)
	}
}

// LogLevel ...
func LogLevel(l string) Option {
	return func(a *application) {
		a.LogLevel = l
	}
}

// ConfLogLevel ...
func ConfLogLevel(k string) Option {
	return func(a *application) {
		a.LogLevel = viper.GetString(k)
	}
}

// Debug ...
func Debug(d bool) Option {
	return func(a *application) {
		a.Debug = d
	}
}

// ConfDebug ...
func ConfDebug(k string) Option {
	return func(a *application) {
		a.Debug = viper.GetBool(k)
	}
}
