package letter

import (
	"sync"
)
// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	res := make(FreqMap)
	channel := make(chan FreqMap) 
	wg := &sync.WaitGroup{}

	wg.Add(len(s)) 
	go func(){
		wg.Wait()
		close(channel)
	}()

	for _, v := range s {
		
		go func (data string){
			defer wg.Done()
			result := Frequency(data) 
			channel <- result
		}(v)
	} 

	for result := range channel{
		for k,v := range result {
			res[k] += v
		}
	}

	


	return res
}

