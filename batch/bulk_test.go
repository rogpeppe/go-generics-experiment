// Code generated by go2go; DO NOT EDIT.


//line bulk_test.go2:1
package batch

//line bulk_test.go2:1
import (
//line bulk_test.go2:1
 "fmt"
//line bulk_test.go2:1
 "log"
//line bulk_test.go2:1
 "sync"
//line bulk_test.go2:1
 "sync/atomic"
//line bulk_test.go2:1
 "testing"
//line bulk_test.go2:1
 "time"
//line bulk_test.go2:1
 "unsafe"
//line bulk_test.go2:1
)

//line bulk_test.go2:11
func TestSingleCall(t *testing.T) {
	var caller instantiate୦୦Caller୦int୦string
	s, err := caller.Do(123, func(is ...int) ([]string, error) {
		if got, want := len(is), 1; got != want {
			t.Errorf("unexpected argument count; got %d want %d", got, want)
		}
		return []string{fmt.Sprint(is[0])}, nil
	})
	if err != nil {
		t.Fatalf("Do returned error: %v", err)
	}
	if got, want := s, "123"; got != want {
		t.Errorf("unexpected result; got %#v want %#v", got, want)
	}
}

func TestMultipleCalls(t *testing.T) {
	caller := instantiate୦୦NewCaller୦int୦string(2, 0)

	callDuration := 50 * time.Millisecond
	stringer := func(is ...int) ([]string, error) {
		time.Sleep(callDuration)
		r := make([]string, len(is))
		for i, v := range is {
			r[i] = fmt.Sprint(v)
		}
		return r, nil
	}

	t0 := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, err := caller.Do(i, stringer)
			if err != nil {
				t.Errorf("got error from Do: %v", err)
			}
			if got, want := r, fmt.Sprint(i); got != want {
				t.Errorf("unexpected result; got %q want %q", got, want)
			}
		}()
	}
	wg.Wait()
	total := time.Since(t0)
	if got, want := total, 2*callDuration+10*time.Millisecond; got > want {
		t.Errorf("total took too long; got %v want %v", got, want)
	}
	log.Printf("total time %v", total)
}

//line bulk_test.go2:62
type instantiate୦୦Caller୦int୦string struct {
//line bulk.go2:17
 initialDelay   time.Duration
			maxConcurrency int
			mu             sync.Mutex
			sem            chan struct{}
			acc            *instantiate୦୦accumulator୦int୦string
}

//line bulk.go2:39
func (g *instantiate୦୦Caller୦int୦string,) DoChan(v int,

//line bulk.go2:39
 call func(vs ...int,

//line bulk.go2:39
 ) ([]string,

//line bulk.go2:39
  error)) <-chan instantiate୦୦Result୦string {

//line bulk.go2:43
 g.mu.Lock()
	if g.sem == nil {
		n := g.maxConcurrency
		if n <= 0 {
			n = 1
		}
		g.sem = make(chan struct{}, n)
	}
	acc := g.acc
	isInitial := acc == nil
	if isInitial {
		acc = new(instantiate୦୦accumulator୦int୦string)
		g.acc = acc
	}
	acc.args = append(acc.args, v)
	resultc := make(chan (instantiate୦୦Result୦string), 1)
	acc.results = append(acc.results, resultc)
	g.mu.Unlock()

	if isInitial {
		g.doCall(call)
	}
	return resultc
}

//line bulk.go2:81
func (g *instantiate୦୦Caller୦int୦string,) Do(v int,

//line bulk.go2:81
 call func(vs ...int,

//line bulk.go2:81
 ) ([]string,

//line bulk.go2:81
  error)) (string,

//line bulk.go2:81
 error) {
	r := <-g.DoChan(v, call)
	return r.Val, r.Err
}

//line bulk.go2:99
func (g *instantiate୦୦Caller୦int୦string,) doCall(fn func(...int,

//line bulk.go2:99
) ([]string,

//line bulk.go2:99
 error)) {
	if g.initialDelay > 0 {
		time.Sleep(g.initialDelay)
	}

//line bulk.go2:106
 g.sem <- struct{}{}
	defer func() {
		<-g.sem
	}()

//line bulk.go2:112
 g.mu.Lock()
	acc := g.acc
	g.acc = nil
	g.mu.Unlock()

	rs, err := fn(acc.args...)
	if err == nil && len(rs) != len(acc.args) {
		err = fmt.Errorf("unexpected result slice length (got %d want %d)", len(rs), len(acc.args))
	}
	if err != nil {
		for _, r := range acc.results {
			r <- instantiate୦୦Result୦string{
				Err: err,
			}
		}
		return
	}
	for i, r := range acc.results {
		r <- instantiate୦୦Result୦string{
			Val: rs[i],
		}
	}
}
//line bulk.go2:30
func instantiate୦୦NewCaller୦int୦string(maxConcurrency int, initialDelay time.Duration) *instantiate୦୦Caller୦int୦string {
	return &instantiate୦୦Caller୦int୦string{
		initialDelay:   initialDelay,
		maxConcurrency: maxConcurrency,
	}
}

//line bulk.go2:35
type instantiate୦୦accumulator୦int୦string struct {
//line bulk.go2:95
 args []int

//line bulk.go2:96
 results []chan<- instantiate୦୦Result୦string
}
//line bulk.go2:97
type instantiate୦୦Result୦string struct {
//line bulk.go2:88
 Val string

//line bulk.go2:89
 Err error
}

//line bulk.go2:90
var _ = fmt.Errorf
//line bulk.go2:90
var _ = log.Fatal

//line bulk.go2:90
type _ sync.Cond

//line bulk.go2:90
var _ = atomic.AddInt32
//line bulk.go2:90
var _ = testing.AllocsPerRun

//line bulk.go2:90
const _ = time.ANSIC

//line bulk.go2:90
type _ unsafe.Pointer