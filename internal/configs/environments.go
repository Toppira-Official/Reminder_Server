package configs

import (
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

type Environment string

type Environments struct {
	PORT Environment
	MODE Environment
}

func LoadEnvironmentsFromEnvFile(envFilePath ...string) {
	err := godotenv.Load(envFilePath...)
	if err != nil {
		log.Println("cannot load environments from given env file")
	}
}

func GetEnvironments() Environments {
	return Environments{
		PORT: Environment(os.Getenv("PORT")).orDefault("3000"),
		MODE: Environment(os.Getenv("MODE")).orPanic().MustBeIn("develop", "production"),
	}
}

func (env Environment) orDefault(defaultValue string) Environment {
	if env == "" {
		return Environment(defaultValue)
	}

	return env
}

func (env Environment) orPanic() Environment {
	if env == "" {
		panic("env not set")
	}

	return env
}

func (env Environment) MustBeIn(allowedValues ...string) Environment {
	doesContain := slices.Contains(allowedValues, string(env))

	if doesContain {
		return env
	}

	panic("env is not allowed")
}
