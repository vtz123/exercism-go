package pov

import (
	"fmt"
	"sort"
)

const (
	MaX_Length = 20
)

type Graph struct {
	grid [][]int 
	res  map[string]int
	resback map[int]string
	index int
}

func New() *Graph {
	g := new(Graph)
	g.res = make(map[string]int, 0)
	g.resback = make(map[int]string, 0)
	return g
}

func (g *Graph) AddNode(nodeLabel string) {
	g.res[nodeLabel] = g.index
	g.resback[g.index] = nodeLabel
	g.index++
}

func (g *Graph) AddArc(from, to string) {
	if g.grid == nil {
		g.grid = make([][]int, MaX_Length)
		for i := range g.grid {
			g.grid[i] = make([]int, MaX_Length)
		}
	} 

	if _,ok := g.res[from]; !ok {
		g.res[from] = g.index
		g.resback[g.index] = from
		g.index++
	}	
	if _,ok := g.res[to]; !ok {
		g.res[to] = g.index
		g.resback[g.index] = to
		g.index++
	}
	//fmt.Println(from, to)
	g.grid[g.res[from]][g.res[to]] = 1
}

func (g *Graph) ArcList() []string {
	res := make([]string, 0)
	
	if g.grid == nil {
		return res
	}
	for i := 0; i < g.index; i++ {
		for j:=0; j < g.index; j++ {
			if g.grid[i][j] == 1 {
				res = append(res, fmt.Sprintf("%s -> %s", g.resback[i] , g.resback[j]))
			}
		}
	}

	sort.Strings(res)

	return res
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	if g.grid == nil {
		return g
	}

	arrived := make(map[int]struct{}) 

	stack := make([]int, 0)
	stack = append(stack, g.res[newRoot])
	for len(stack) != 0 {
		rootindex := stack[0]
		stack = stack[1:]
		for i := 0; i < g.index; i++ {
			
			if _,ok := arrived[i]; g.grid[i][rootindex] == 1 && !ok {
				g.grid[i][rootindex] = 0
				g.grid[rootindex][i] = 1
				stack = append(stack, i)
			} 
			
		}
		arrived[rootindex] = struct{}{}
		
	}

	return g
}