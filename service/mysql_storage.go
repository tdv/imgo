package service

import (
	"database/sql"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlStorage struct {
	Storage
	db *sql.DB
}

func (this *mysqlStorage) init(host, port, dbname, user, password string) error {
	db, err := sql.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+port+")/"+dbname,
	)

	if err != nil {
		return err
	}

	this.db = db

	return nil
}

func (this *mysqlStorage) Put(id string, buf []byte) error {
	tr, err := this.db.Begin()
	if err != nil {
		return err
	}
	defer tr.Rollback()

	if stmt, err := tr.Prepare("INSERT INTO images (id, data) " +
		"VALUES(?, ?); "); err != nil {
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

func (this *mysqlStorage) Get(id string) ([]byte, error) {
	tr, err := this.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Rollback()

	image := tr.QueryRow("SELECT data FROM images WHERE id = ?; ", id)
	if image == nil {
		return nil, errors.New("Failed to get image with id \"" + id + "\"")
	}

	var blob []byte
	image.Scan(&blob)

	return blob, nil
}

const ImplMySql = "mysql"

var _ = RegisterEntity(
	EntityStorage,
	ImplMySql,
	func(ctx BuildContext) (interface{}, error) {
		config := ctx.GetConfig()
		storage := mysqlStorage{}
		if err := storage.init(
			config.GetStrVal("host"),
			strconv.Itoa(config.GetIntVal("port")),
			config.GetStrVal("dbname"),
			config.GetStrVal("user"),
			config.GetStrVal("password"),
		); err != nil {
			return nil, err
		}
		return &storage, nil
	},
)
