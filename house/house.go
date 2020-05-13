package house

import (
	"strings"
	"fmt"
)

var (
	poem = []string{
		"the malt\nthat lay in",
		"the rat\nthat ate",
		"the cat\nthat killed",
		"the dog\nthat worried",
		"the cow with the crumpled horn\nthat tossed",
		"the maiden all forlorn\nthat milked",
		"the man all tattered and torn\nthat kissed",
		"the priest all shaven and shorn\nthat married",
		"the rooster that crowed in the morn\nthat woke",
		"the farmer sowing his corn\nthat kept",
		"the horse and the hound and the horn\nthat belonged to",
	}
	res = []string{}
)

func Verse(index int) string {
	if index == 1 {
		return "This is the house that Jack built."
	}

	first := "This is "
	v := "This is the house that Jack built."

	for i := 0; i < index-1; i++ {
		v = fmt.Sprintf("%s%s%s", first,  poem[i], v[7:])
	}

	return v
}

func Song() string {
	first := "This is "
	
	v := "This is the house that Jack built."
	res := make([]string, 0)
	res = append(res, v)
	for i := 0; i < 11; i++ {
		v = fmt.Sprintf("%s%s%s", first,  poem[i], v[7:])
		res = append(res, v)
	}

	return strings.Join(res, "\n\n")
}