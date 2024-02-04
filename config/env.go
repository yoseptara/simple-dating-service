package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV" required:"true"`
	ServerPort     string `mapstructure:"SERVER_PORT" required:"true"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT" required:"true"`

	BeliesimWebFrontendUrl string `mapstructure:"BELIESIM_WEB_FRONTEND_URL" required:"true"`

	PostgresHost    string `mapstructure:"POSTGRES_HOST" required:"true"`
	PostgresPort    string `mapstructure:"POSTGRES_PORT" required:"true"`
	PostgresUser    string `mapstructure:"POSTGRES_USERNAME" required:"true"`
	PostgresPass    string `mapstructure:"POSTGRES_PASS" required:"true"`
	PostgresDbName  string `mapstructure:"POSTGRES_DB_NAME" required:"true"`
	PostgresConnStr string `mapstructure:"POSTGRES_CONNECTION_STRING" required:"true"`

	XenditHostUrl                   string `mapstructure:"XENDIT_HOST_URL" required:"true"`
	XenditWriteKey                  string `mapstructure:"XENDIT_WRITE_KEY" required:"true"`
	XenditInvoiceSuccessRedirectUrl string `mapstructure:"XENDIT_INVOICE_SUCCESS_REDIRECT_URL" required:"true"`
	XenditInvoiceFailureRedirectUrl string `mapstructure:"XENDIT_INVOICE_FAILURE_REDIRECT_URL" required:"true"`

	UsimsaHostUrl   string `mapstructure:"USIMSA_HOST_URL" required:"true"`
	UsimsaAccessKey string `mapstructure:"USIMSA_ACCESS_KEY" required:"true"`
	UsimsaSecretKey string `mapstructure:"USIMSA_SECRET_KEY" required:"true"`

	AwsAccessKeyId     string `mapstructure:"AWS_ACCESS_KEY_ID" required:"true"`
	AwsDefaultRegion   string `mapstructure:"AWS_DEFAULT_REGION" required:"true"`
	AwsSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY" required:"true"`
}

// Custom error type for validation errors
type ConfigValidationError struct {
	Field string
}

func (e *ConfigValidationError) Error() string {
	return fmt.Sprintf("config validation error: field '%s' is required", e.Field)
}

func validateRequiredFields(env Env) error {
	t := reflect.TypeOf(env)
	v := reflect.ValueOf(env)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		required, exists := field.Tag.Lookup("required")

		if exists && required == "true" {
			if value == "" {
				return &ConfigValidationError{Field: field.Name}
			}
		}
	}

	return nil
}

func NewEnv() (Env, error) {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Println(".env file not found, but that's OK. Looking in environment.")
			viper.AutomaticEnv() // read in environment variables that match

			env.AppEnv = viper.GetString("APP_ENV")
			env.ServerPort = viper.GetString("SERVER_PORT")
			env.ContextTimeout = viper.GetInt("CONTEXT_TIMEOUT")

			env.BeliesimWebFrontendUrl = viper.GetString("BELIESIM_WEB_FRONTEND_URL")

			env.PostgresHost = viper.GetString("POSTGRES_HOST")
			env.PostgresPort = viper.GetString("POSTGRES_PORT")
			env.PostgresUser = viper.GetString("POSTGRES_USERNAME")
			env.PostgresPass = viper.GetString("POSTGRES_PASS")
			env.PostgresDbName = viper.GetString("POSTGRES_DB_NAME")

			env.XenditHostUrl = viper.GetString("XENDIT_HOST_URL")
			env.XenditWriteKey = viper.GetString("XENDIT_WRITE_KEY")
			env.XenditInvoiceSuccessRedirectUrl = viper.GetString("XENDIT_INVOICE_SUCCESS_REDIRECT_URL")
			env.XenditInvoiceFailureRedirectUrl = viper.GetString("XENDIT_INVOICE_FAILURE_REDIRECT_URL")

			env.UsimsaHostUrl = viper.GetString("USIMSA_HOST_URL")
			env.UsimsaAccessKey = viper.GetString("USIMSA_ACCESS_KEY")
			env.UsimsaSecretKey = viper.GetString("USIMSA_SECRET_KEY")

			env.AwsAccessKeyId = viper.GetString("AWS_ACCESS_KEY_ID")
			env.AwsDefaultRegion = viper.GetString("AWS_DEFAULT_REGION")
			env.AwsSecretAccessKey = viper.GetString("AWS_SECRET_ACCESS_KEY")

		} else {
			// Config file was found but another error was produced
			log.Fatalf("Fatal error config file: %s \n", err)
		}
	} else {
		// .env file was found, use Unmarshal as usual
		err := viper.Unmarshal(&env)
		if err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}
	}

	env.PostgresConnStr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", env.PostgresUser, env.PostgresPass, env.PostgresHost, env.PostgresPort, env.PostgresDbName)

	if err := validateRequiredFields(env); err != nil {
		return env, err
	}

	if env.AppEnv == "dev" {
		log.Println("The App is running in development env")
	}

	return env, nil
}
