package database

import (
	"context"
	"database/sql"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/entity"
	"github.com/google/uuid"
	"time"
)

const insertCotacaoSql = "INSERT INTO cotacoes(ID, FROM_MOEDA, TO_MOEDA, BID, CREATED_AT) VALUES (?, ?, ?, ?, datetime('now'))"

type CotacaoRepository struct {
	Db *sql.DB
}

func (c CotacaoRepository) Save(cotacao *entity.Cotacao) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	defer cancel()

	tx, err := c.Db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertCotacaoSql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(uuid.New().String(), cotacao.Code, cotacao.Codein, cotacao.Bid); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewCotacaoRepository(db *sql.DB) *CotacaoRepository {
	return &CotacaoRepository{
		Db: db,
	}
}
