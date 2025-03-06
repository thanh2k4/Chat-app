package database

type Config struct {
	Database struct {
		Postgres PostgresConfig `yaml:"postgres"`
	} `yaml:"database"`
}
