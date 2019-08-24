package service

/*
Table for storing images

BEGIN;

CREATE TABLE images
(
	id varchar unique,
	data bytea
);

COMMIT;

*/

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type postgresStorage struct {
	Storage
	db *sql.DB
}

func (this *postgresStorage) init(host string, port string, dbname string, sslmode string, user string, password string) error {
	db, err := sql.Open(
		"postgres",
		"host="+host+" port="+port+" dbname="+dbname+" "+
			"sslmode="+sslmode+" user="+user+" password= "+password+"",
	)

	if err != nil {
		return err
	}

	this.db = db

	return nil
}

func (this *postgresStorage) Put(id string, buf []byte) error {
	tr, err := this.db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	if stmt, err := tr.Prepare("INSERT INTO images (id, data) " +
		"VALUES($1::varchar, $2::bytea) " +
		"ON CONFLICT (id) DO NOTHING; "); err != nil {
		return err
	} else {
		defer stmt.Close()
		if _, err := stmt.Exec(id, buf); err != nil {
			return err
		}
	}

	tr.Commit()
	return nil
}

func (this *postgresStorage) Get(id string) ([]byte, error) {
	tr, err := this.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	image := tr.QueryRow("SELECT data FROM images WHERE id = $1::varchar; ", id)
	if image == nil {
		return nil, errors.New("Failed to get image with id \"" + id + "\"")
	}

	var blob []byte
	image.Scan(&blob)

	return blob, nil
}

func CreatePostgresStorage(config *viper.Viper) (Storage, error) {
	storage := postgresStorage{}
	if err := storage.init(
		config.GetString("storage.postgres.host"),
		config.GetString("storage.postgres.port"),
		config.GetString("storage.postgres.dbname"),
		config.GetString("storage.postgres.sslmode"),
		config.GetString("storage.postgres.user"),
		config.GetString("storage.postgres.password"),
	); err != nil {
		return nil, err
	}
	return &storage, nil
}
