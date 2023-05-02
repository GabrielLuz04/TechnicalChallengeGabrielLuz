package entities

type Crypto struct {
	Upvotes int    `json:"upvotes"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Slug    string `json:"slug"`
	Quote   struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quote"`
}

type CryptoRequest struct {
	Data []Crypto `json:"data"`
}

var Cryptos = make([]Crypto, 100)

var CryptosCMC = CryptoRequest{}

// var Cryptos = []Crypto{
// 	{Id: "1", Name: "Bitcoin", Upvotes: 0},
// 	{Id: "2", Name: "Ethereum", Upvotes: 0},
// 	{Id: "3", Name: "Tether", Upvotes: 0},
// 	{Id: "4", Name: "Cardano", Upvotes: 0},
// 	{Id: "5", Name: "Dogecoin", Upvotes: 0},
// 	{Id: "6", Name: "Polygon", Upvotes: 0},
// 	{Id: "7", Name: "Solana", Upvotes: 0},
// 	{Id: "8", Name: "Polkadot", Upvotes: 0},
// }
