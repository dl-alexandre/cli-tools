// Package config provides configuration management for Go CLI applications.
// Uses viper for config file management, environment variables, and precedence handling.
//
// Configuration precedence (highest to lowest):
//  1. CLI flags (applied by caller after loading)
//  2. Environment variables
//  3. Config file
//  4. Defaults
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Loader handles configuration loading with standard precedence patterns
type Loader struct {
	binaryName string
	envPrefix  string
	viper      *viper.Viper
}

// Config provides the interface for configuration data
type Config interface {
	// Unmarshal into a struct
	Unmarshal(raw map[string]interface{}) error
}

// Flags holds CLI flag values that override config
type Flags struct {
	ConfigFile string
	// Additional flags are application-specific
}

// NewLoader creates a new configuration loader
// binaryName: name of the binary (used for default config paths)
// envPrefix: prefix for environment variables (e.g., "UNIFI", "CIMIS")
func NewLoader(binaryName, envPrefix string) *Loader {
	return &Loader{
		binaryName: binaryName,
		envPrefix:  envPrefix,
		viper:      viper.New(),
	}
}

// Load loads configuration with precedence:
// 1. CLI flags (applied after this call)
// 2. Environment variables
// 3. Config file
// 4. Defaults
func (l *Loader) Load(flags Flags) (*viper.Viper, error) {
	v := l.viper

	// Set defaults (caller should call SetDefaults before Load)

	// Set config file if provided
	if flags.ConfigFile != "" {
		v.SetConfigFile(expandPath(flags.ConfigFile))
	} else {
		// Default config location: ~/.config/{binary}/config.yaml
		configDir := l.defaultConfigDir()
		v.AddConfigPath(configDir)
		v.SetConfigName("config")
		v.SetConfigType("yaml")
	}

	// Read config file (ignore error if not found)
	if err := v.ReadInConfig(); err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if !errors.As(err, &notFoundErr) {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Bind environment variables
	l.bindEnvVars(v)

	return v, nil
}

// SetDefaults sets standard defaults
func (l *Loader) SetDefaults(defaults map[string]interface{}) {
	for key, value := range defaults {
		l.viper.SetDefault(key, value)
	}
}

// bindEnvVars configures environment variable binding
func (l *Loader) bindEnvVars(v *viper.Viper) {
	v.SetEnvPrefix(l.envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
}

// defaultConfigDir returns the default config directory
func (l *Loader) defaultConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "."
	}
	return filepath.Join(home, ".config", l.binaryName)
}

// ConfigFilePath returns the path to the default config file
func (l *Loader) ConfigFilePath() string {
	return filepath.Join(l.defaultConfigDir(), "config.yaml")
}

// ConfigExists checks if a config file exists
func (l *Loader) ConfigExists() bool {
	_, err := os.Stat(l.ConfigFilePath())
	return !os.IsNotExist(err)
}

// Save saves configuration to the default location
func (l *Loader) Save(data map[string]interface{}) error {
	configDir := l.defaultConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := l.ConfigFilePath()

	v := viper.New()
	v.SetConfigFile(configPath)

	// Set values
	for key, value := range data {
		v.Set(key, value)
	}

	if err := v.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

// GetCredentials retrieves username and password from environment or flags
// This is a common pattern across many dl-alexandre CLIs
func GetCredentials(flagsUsername, flagsPassword, envPrefix string) (string, string, error) {
	username := flagsUsername
	password := flagsPassword

	// Check environment variables
	if username == "" {
		username = os.Getenv(envPrefix + "_USERNAME")
	}
	if password == "" {
		password = os.Getenv(envPrefix + "_PASSWORD")
	}

	if username == "" {
		return "", "", fmt.Errorf("username required. Set %s_USERNAME environment variable or use --username flag", envPrefix)
	}

	if password == "" {
		return "", "", fmt.Errorf("password required. Set %s_PASSWORD environment variable or use --password flag", envPrefix)
	}

	return username, password, nil
}

// expandPath expands ~ and environment variables in a path
func expandPath(path string) string {
	if path == "" {
		return path
	}

	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err == nil {
			path = filepath.Join(home, path[1:])
		}
	}

	path = os.ExpandEnv(path)
	return path
}

// GetEnvOrDefault returns environment variable value or default
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvOrDefaultInt returns environment variable value as int or default
func GetEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var result int
		_, _ = fmt.Sscanf(value, "%d", &result)
		return result
	}
	return defaultValue
}
