package todo

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
)

type AuthMode string

const (
	Session AuthMode = "session"
	Paseto  AuthMode = "paseto"
)

type (
	Configuration struct {
		// Listener http listener binding options.
		Listener Listener `json:"listener"`

		// DSN database connection string.
		DSN string `json:"dsn"`

		// Database is name of primary database name for entire app.
		// default: todo_be
		Database string `json:"database"`

		Jwt Jwt `json:"jwt"`
	}

	Listener struct {
		// Host is network address for bind Server http listener to it.
		// default: 127.0.0.1
		Host string `json:"host" mapstructure:"host"`

		// Port is network port for bind Server http listener to it.
		// default: 8080
		Port int `json:"port" mapstructure:"port"`

		// Cert is path to TLS certificate file.
		// if Cert is not specified, Server listener runs without TLS.
		Cert string `json:"cert" mapstructure:"cert"`

		// Key is path to TLS certificate PrivateKey file.
		// it ignored if Cert is not specified.
		Key string `json:"key" mapstructure:"key"`

		// AllowedHosts is allowed host for CORS configuration.
		// It applied in production mode
		AllowedHosts []string `json:"allowed_hosts" mapstructure:"allowed_hosts"`

		// SSLHost is ssl host for gin secure configuration.
		// It applied in production mode
		SSLHost string `json:"ssl_host" mapstructure:"ssl_host"`

		// SessionsSecret is secret key string that used by gin session.
		SessionsSecret string `json:"sessions_secret" mapstructure:"sessions_secret"`

		// AuthMode defines user authentication mechanism (session, paseto)
		AuthMode AuthMode `json:"auth_mode"`
	}

	// Jwt contains JWT configuration options.
	Jwt struct {
		TokenExpire   int64  `json:"token_expire"`
		RefreshExpire int64  `json:"refresh_expire"`
		Issuer        string `json:"issuer"`
		Audience      string `json:"audience"`
		SubjectKey    string `json:"subject_key"`
		IdentityKey   string `json:"identity_key"`
	}
)

// NewConfiguration returns new Configuration with default options.
func NewConfiguration(secret string) (*Configuration, error) {
	config := &Configuration{}

	err := config.loadConfig()
	if err != nil {
		return nil, err
	}

	if len(secret) > 0 {
		config.Listener.SessionsSecret = secret
	}

	return config, nil
}

func (c *Configuration) loadConfig() error {
	path := ""

	if runtime.GOOS == "windows" {
		path = ".\\config\\config.json"
	} else {
		path = "./config/config.json"
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer closeFile(file)

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, c)
	if err != nil {
		return err
	}

	return nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Panicln(err.Error())
	}
}
