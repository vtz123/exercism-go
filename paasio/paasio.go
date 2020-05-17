package paasio

import (
	"io"
	"sync"
)

type counter struct {
	mu 		sync.Mutex 
	n		int64
	nops	int
}

func (c *counter) Add(n int)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nops++
	c.n += int64(n)
}

func (c *counter) Count() ( int64,  int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n, c.nops
}

type readCounter struct {
	counter
	r   io.Reader
}

func NewReadCounter (r io.Reader) ReadCounter {
	return &readCounter{
		r:	r,		
	}
}

func (rc *readCounter) ReadCount() (n int64, nops int) {
	return rc.Count()
}

func (rc *readCounter) Read(p []byte) (n int, err error) {
	n, err = rc.r.Read(p) 
	rc.Add(n)
	return 
}

type writeCounter struct {
	C 	counter
	w   io.Writer
}

func NewWriteCounter (w io.Writer) WriteCounter {
	return &writeCounter{
		C:	counter{},
		w:	w,		
	}
}

func (rc *writeCounter) WriteCount() (n int64, nops int) {
	return rc.C.Count()
}

func (rc *writeCounter) Write(p []byte) (n int, err error) {
	n, err = rc.w.Write(p) 
	rc.C.Add(n)
	return 
}

type readwriteCounter struct {
	//c 	counter
	ReadCounter
	WriteCounter
}

func NewReadWriteCounter (rw io.ReadWriter) ReadWriteCounter {
	return &readwriteCounter{
		//c:	counter{},
		NewReadCounter(rw),
		NewWriteCounter(rw),
	}
}

