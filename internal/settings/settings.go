package settings

import "fmt"

type DBSettings struct {
	Host     string
	User     string
	Password string
	Port     int
	Name     string
}

func (db DBSettings) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s"+
		"dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)
}

type Settings struct {
	DBSettings    DBSettings
	ServerAddress string
}

func New() (*Settings, error) {
	return &Settings{
		DBSettings: DBSettings{
			Host:     "localhost",
			User:     "jellyapp",
			Password: "fish",
			Port:     5432,
			Name:     "jelly",
		},
		ServerAddress: ":3001",
	}, nil
}
