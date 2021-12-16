package infra

import (
	"context"
	"database/sql"
)

type DbConnector interface {
	Execute(ctx context.Context, db *sql.DB, sqlStatement string) error
	Retrieve(ctx context.Context, db *sql.DB, statement string) (*sql.Rows, error)
}

type DbConnectorImpl struct {
}

func NewDbConnectorImpl() DbConnector {
	return DbConnectorImpl{}
}

func (DbConnectorImpl) Execute(ctx context.Context, db *sql.DB, sqlStatement string) error {
	var err error

	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	query, err := db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer query.Close()
	newRecord := query.QueryRowContext(ctx)

	var newID int64
	err = newRecord.Scan(&newID)
	if err != nil {
		return err
	}

	return nil
}

func (DbConnectorImpl) Retrieve(ctx context.Context, db *sql.DB, statement string) (*sql.Rows, error) {
	ctx1 := context.Background()
	err := db.PingContext(ctx1)
	if err != nil {
		return nil, err
	}

	data, err := db.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	return data, nil
}
