package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/intamyuto/workshop/logger"
)

func main() {
	ctx := context.Background()

	ctx = logger.ContextWithTimestamp(ctx, time.Now())

	if err := run(ctx); err != nil {
		logger.Errorf(ctx, "%v", err)
		os.Exit(1)
	}

	logger.Infof(ctx, "exit")
}

type line struct {
	idx   int
	value string
	err   error
}

func run(ctx context.Context) error {
	r, err := os.Open("input")
	if err != nil {
		return err
	}
	defer r.Close()

	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	ctx, cancelt := context.WithTimeoutCause(ctx, 3*time.Second, errors.New("custom"))
	defer cancelt()

	w, err := os.OpenFile("output", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer w.Close()

	in := make(chan line)

	go read(ctx, r, in)

	out := make([]chan line, runtime.NumCPU())
	for i := 0; i < len(out); i++ {
		out[i] = make(chan line)
		go process(ctx, in, out[i])
	}

	write(ctx, w, merge(out))

	return ctx.Err()
}

func merge(in []chan line) chan line {
	out := make(chan line)

	var wg sync.WaitGroup

	// wg.Wait (счетчик внутри — больше 0 — ждем; =0 не ждем)
	// wg.Add(n int) — увеличивает счетчик на n
	// wg.Done() // уменьшает счетчик на 1

	wg.Add(len(in))
	for _, ch := range in {
		go func(ch chan line) {
			defer wg.Done()

			for l := range ch {
				out <- l
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func read(ctx context.Context, r io.Reader, out chan<- line) {
	s := bufio.NewScanner(r)

	idx := 1
	for s.Scan() {
		select {
		case <-ctx.Done():
			break
		default:
			out <- line{value: s.Text(), idx: idx}
			idx++
		}
	}
	close(out)
}

func process(ctx context.Context, in <-chan line, out chan<- line) {
	for l := range in {
		parts := strings.Split(l.value, " ")
		for i := 0; i < len(parts); i++ {
			if parts[i] == "error" {
				out <- line{idx: l.idx, err: errors.New("error in line")}
				continue
			}
			parts[i] = strings.Title(parts[i])
		}

		<-time.After(time.Second)
		out <- line{idx: l.idx, value: strings.Join(parts, " ")}
	}
	close(out)
}

func write(ctx context.Context, w io.Writer, in <-chan line) {
	m, i := make(map[int]line), 1

	for l := range in {
		if l.err != nil {
			logger.Errorf(ctx, "%d: %v", l.idx, l.err)
			continue
		}

		if l.idx == i {
			fmt.Fprintf(w, "%d: %s\n", l.idx, l.value)
			i++

			continue
		}

		m[l.idx] = l

		for {
			if cached, ok := m[i]; ok {
				fmt.Fprintf(w, "%d: %s\n", cached.idx, cached.value)
				delete(m, i)
				i++
			}
			break
		}
	}

	for l := range in {
		fmt.Fprintf(w, "%d: %s\n", l.idx, l.value)
	}
}
