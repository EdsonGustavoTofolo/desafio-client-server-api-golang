package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EdsonGustavoTofolo/desafio-client-server-api-golang/internal/infra/controllers"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	resp, cancel := getCotacaoDolar()

	defer cancel()

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Panicf("Failed request. Status code response %v\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var cotacaoResponse controllers.CotacaoResponse

	if err := json.Unmarshal(body, &cotacaoResponse); err != nil {
		panic(err)
	}

	save(cotacaoResponse.Valor)
}

func getCotacaoDolar() (resp *http.Response, cancel context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)

	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("from", "USD")
	q.Add("to", "BRL")

	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	return
}

func save(valor float64) {
	file, err := os.OpenFile("cotacao.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %.3f\n", valor))

	if err != nil {
		panic(err)
	}
}
