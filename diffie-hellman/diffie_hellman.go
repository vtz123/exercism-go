package diffiehellman

import (
	"crypto/rand"
	"math/big"
	_"time"
)


var One = big.NewInt(1)

func PrivateKey(p *big.Int) *big.Int {
	
	res := big.NewInt(1)

	for res.Cmp(One) <= 0 || res.Cmp(p) >= 0 {
		n, _ := rand.Int(rand.Reader, p)
		// (res+randN )% p
		res.Add(res, n).DivMod(res, p, res)
	}

	return res
}

func  PublicKey(a, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g),a,p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

func SecretKey(a, B, p *big.Int) *big.Int {
	return new(big.Int).Exp(B,a,p)
}