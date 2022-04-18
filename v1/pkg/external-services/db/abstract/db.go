package abstract

// DB acceso a la db
type DB interface {
	Currencies() CurrenciesRepo
	Requests() RequestsRepo
}
