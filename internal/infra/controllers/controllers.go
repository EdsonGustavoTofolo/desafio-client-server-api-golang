package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/infra/database"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/usecase"
	"log"
	"net/http"
)

type CotacaoResponse struct {
	Valor float64
}

type CotacaoController struct {
	Db *sql.DB
}

func (c *CotacaoController) GetCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Cotacao received request")

	if !r.URL.Query().Has("from") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'from' unit"))
		return
	}

	if !r.URL.Query().Has("to") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing 'to' unit"))
		return
	}

	var from, to = r.URL.Query().Get("from"), r.URL.Query().Get("to")

	cotacao, err := usecase.GetCotacao{
		CotacaoRepository: database.NewCotacaoRepository(c.Db),
		CoinFrom:          from,
		CoinTo:            to,
	}.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&CotacaoResponse{Valor: cotacao.Bid}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong. Encode response failed."))
		log.Println(err.Error())
		return
	}

	log.Println("Cotacao request performed successfully")
}
