package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/entity"
	"io"
	"net/http"
	"time"
)

const exchangeRateUrl = "https://economia.awesomeapi.com.br/json/last/"

type GetCotacao struct {
	CotacaoRepository entity.CotacaoRepository
	CoinFrom          string
	CoinTo            string
}

func (c GetCotacao) Execute() (*entity.Cotacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)

	defer cancel()

	url := exchangeRateUrl + fmt.Sprintf("%v-%v", c.CoinFrom, c.CoinTo)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var response map[string]json.RawMessage

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	var cotacao entity.Cotacao

	if err := json.Unmarshal(response[c.CoinFrom+c.CoinTo], &cotacao); err != nil {
		return nil, err
	}

	if err := c.CotacaoRepository.Save(&cotacao); err != nil {
		return nil, err
	}

	return &cotacao, nil
}
