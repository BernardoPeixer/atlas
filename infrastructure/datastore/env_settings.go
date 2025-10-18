package datastore

import (
	"atlas/domain/entities"
	"atlas/util"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type connectionStruct struct {
	cfg entities.MySQL
	db  *sql.DB
}

type repositorySettings struct {
	configs    entities.Config
	connection connectionStruct
}

func NewRepositorySettings(config entities.Config) *repositorySettings {
	repository := &repositorySettings{
		configs: config,
		connection: connectionStruct{
			cfg: entities.MySQL{
				DBHost:     config.MySQL.DBHost,
				DBPort:     config.MySQL.DBPort,
				DBName:     config.MySQL.DBName,
				DBUser:     config.MySQL.DBUser,
				DBPassword: config.MySQL.DBPassword,
			},
		},
	}

	repository.OpenConnection()

	return repository
}

func (s *repositorySettings) OpenConnection() {
	dbCfg := s.connection.cfg

	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbCfg.DBUser,
		dbCfg.DBPassword,
		dbCfg.DBHost,
		dbCfg.DBPort,
		dbCfg.DBName,
	)

	connection, err := sql.Open("mysql", source)
	if err != nil {
		log.Printf("Error openning the database connection [%s] | [%v]", source, err)
		panic(err)
	}
	
	connection.SetMaxOpenConns(20)
	connection.SetConnMaxLifetime(20 * time.Minute)
	connection.SetConnMaxIdleTime(20 * time.Minute)

	s.connection.db = connection
}

func (s *repositorySettings) Connection() *sql.DB {
	return s.connection.db
}

func (s *repositorySettings) ServerTime(ctx context.Context) (*util.DateTime, error) {
	//language=sql
	query := `SELECT current_timestamp` + ""

	var conn *sql.DB

	conn = s.connection.db

	row, err := conn.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error in [QueryContext]: %v", err)
		return nil, err
	}

	var date time.Time

	err = row.Scan(&date)
	if err != nil {
		log.Printf("Error in [Scan]: %v", err)
		return nil, err
	}

	return util.NewDateTime(&date), nil
}

func (s *repositorySettings) Dismount() error {
	err := s.connection.db.Close()
	if err != nil {
		log.Printf("Error in [Dismount]: %v", err)
	}
	return nil
}
