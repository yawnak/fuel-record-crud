package repoadapt

import (
	"context"
	dbsql "database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/yawnak/fuel-record-crud/ent"
)

type Client struct {
	cl *ent.Client
}

func NewClientPSQL(user, password, host, port, dbname string) (*Client, error) {
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	db, err := dbsql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	return &Client{
		cl: client,
	}, nil
}

func (c *Client) CarRepo() *CarRepositoryPSQL {
	return &CarRepositoryPSQL{
		client: c.cl.Car,
	}
}

func (c *Client) FuelGaugeChangeRepo() *FuelRecordRepoPSQL {
	return &FuelRecordRepoPSQL{
		client: c.cl.FuelRecord,
	}
}

func (c *Client) OdometerIncreaseRepo() *OdometerRecordRepoPSQL {
	return &OdometerRecordRepoPSQL{
		client: c.cl.OdometerRecord,
	}
}

func (c *Client) BeginTx(ctx context.Context) (*Tx, error) {
	tx, err := c.cl.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &Tx{
		tx: tx,
	}, nil
}

type Tx struct {
	tx *ent.Tx
}

func (t *Tx) Commit() error {
	return t.tx.Commit()
}

func (t *Tx) Rollback() error {
	return t.tx.Rollback()
}

func (t *Tx) CarRepo() *CarRepositoryPSQL {
	return &CarRepositoryPSQL{
		client: t.tx.Car,
	}
}

func (t *Tx) FuelGaugeChangeRepo() *FuelRecordRepoPSQL {
	return &FuelRecordRepoPSQL{
		client: t.tx.FuelRecord,
	}
}

func (t *Tx) OdometerIncreaseRepo() *OdometerRecordRepoPSQL {
	return &OdometerRecordRepoPSQL{
		client: t.tx.OdometerRecord,
	}
}
