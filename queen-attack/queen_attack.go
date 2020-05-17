package queenattack

import (
	"errors"
)

func CanQueenAttack(w, b string) ( bool, error) {
	if !helper(w, b) {
		return false, errors.New("panic!")
	}

	if w[0] == b[0] || w[1] == b[1] || w[0] -b[0]  == w[1] - b[1] || w[0] -b[0] == b[1] - w[1] {
		
		return true, nil
	}

	return false, nil
}

func helper(w, b string) bool {
	if w == b  || w == "" || b == "" || 
	w[0] < 'a' || w[0] > 'h'|| w[1] < '1' || w[1] > '8' || 
	b[0] < 'a' || b[0] > 'h'|| b[1] < '1' || b[1] > '8'{
		return false
	}

	return true
}