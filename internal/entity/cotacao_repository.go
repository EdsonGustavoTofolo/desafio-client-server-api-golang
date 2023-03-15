package entity

type CotacaoRepository interface {
	Save(cotacao *Cotacao) error
}
