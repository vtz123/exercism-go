package react

type react struct {
	clist 	 	[]*cell
	cmap		map[*cell]int  // 状态改变时， 便于快速找到起始index
	curindex  	int
}

func New() Reactor {
	return &react{
		clist:	make([]*cell, 0),
		cmap:	make(map[*cell]int),
		curindex:	-1,
	}
}

func (r *react ) CreateInput(val int) InputCell {
	cell1 := &cell{
		value: val,
		compute: func(){},
		r:	r,
	}
	r.add(cell1)

	return cell1
}

func (r *react) add(c *cell) {
	r.curindex++
	r.clist = append(r.clist, c)
	r.cmap[c] = r.curindex
}

func (r *react ) CreateCompute1(c Cell,fn  func(int) int ) ComputeCell {
	newcell := &cell{ 
		c1: c,
		r:	r,
	} 
	newcell.compute = func () {
		newcell.value = fn(newcell.c1.Value())
	}
	newcell.Compute()
	
	r.add(newcell)  

	return newcell
}

func (r *react ) CreateCompute2(c1,c2 Cell,fn  func(int, int) int ) ComputeCell {
	newcell := &cell{ 
		c1: c1,
		c2: c2,
		r:	r,
	} 
	newcell.compute = func () {
		newcell.value = fn(newcell.c1.Value(), newcell.c2.Value())
	}
	newcell.Compute()
	
	r.add(newcell)  

	return newcell
}

func (r *react) update(c *cell) {
	index := r.cmap[c]
	for i := index; i < len(r.clist); i++ {
		c1 := r.clist[i]
		oldval := c1.Value()
		c1.compute()
		if oldval != c1.Value() {
			for fn,_ := range c1.callbackfns {
				(*fn)(c1.Value())
			}
		}

	}
}

type cell struct {
	value  	int
	c1, c2 	Cell
	r 		*react
	callbackfns map[*func(int)]struct{}

	compute func()
}


func (c *cell) Value() int {
	return c.value
}

func (c *cell) Compute() {
	c.compute()
}

func (c *cell) SetValue(val int){
	c.value = val

	c.r.update(c)
}

func (c *cell) AddCallback(fn func(int)) Canceler {
	if c.callbackfns == nil {
		c.callbackfns = make(map[*func(int)]struct{})
	}
	c.callbackfns[&fn] = struct{}{}

	return &canceler{func () {
		delete(c.callbackfns, &fn)
	}}
}

type canceler struct {
	cancelFn func()
}

func (cl *canceler) Cancel() {
	cl.cancelFn()
}
