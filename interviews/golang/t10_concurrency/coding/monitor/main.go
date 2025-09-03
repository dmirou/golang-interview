package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	// MaxGoroutines limits concurrent net requests. It's optional.
	MaxGoroutines = 2
	// Timeout for net requests
	Timeout = 2 * time.Second
)

type SiteStatus struct {
	Name          string
	StatusCode    int
	TimeOfRequest time.Time
}

type Monitor struct {
	StatusMap        map[string]SiteStatus
	Mtx              sync.Mutex
	G                errgroup.Group
	Sites            []string
	RequestFrequency time.Duration
}

func NewMonitor(sites []string, requestFrequency time.Duration) *Monitor {
	return &Monitor{
		StatusMap:        make(map[string]SiteStatus),
		Mtx:              sync.Mutex{},
		Sites:            sites,
		RequestFrequency: requestFrequency,
	}
}

func (m *Monitor) Run(ctx context.Context) error {
	// run printStatuses and checkSite in different goroutines

	var g errgroup.Group

	g.Go(func() error {
		defer fmt.Println("checking done")
		ticker := time.NewTicker(m.RequestFrequency)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
			fmt.Println("checking tick")

			var siteG errgroup.Group
			siteG.SetLimit(MaxGoroutines)

			for _, site := range m.Sites {
				siteG.Go(func() error {
					m.checkSite(ctx, site)
					return nil
				})
			}
			if err := siteG.Wait(); err != nil {
				fmt.Printf("checking wait: %v\n", err)
				return err
			}
		}
	})

	g.Go(func() error {
		defer fmt.Println("printing done")
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(5 * time.Second):
			}
			fmt.Println("printing tick")

			if err := m.printStatuses(ctx); err != nil {
				return fmt.Errorf("print statuses: %w", err)
			}
		}
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func (m *Monitor) checkSite(ctx context.Context, site string) {
	// check site and write result in StatusMap
	fmt.Println("checking site", site)

	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, site, io.TeeReader(bytes.NewBuffer(nil), io.Discard),
	)
	if err != nil {
		return
	}

	reqTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	m.Mtx.Lock()
	defer m.Mtx.Unlock()
	m.StatusMap[site] = SiteStatus{
		Name:          site,
		StatusCode:    resp.StatusCode,
		TimeOfRequest: reqTime,
	}
}

func (m *Monitor) printStatuses(ctx context.Context) error {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()
	for site, status := range m.StatusMap {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		fmt.Printf(
			"site: %s, status_code: %d, time: %v\n",
			site, status.StatusCode, status.TimeOfRequest.Format(time.RFC3339),
		)
	}

	return nil
}

func main() {
	m := NewMonitor([]string{
		"https://ya.ru/",
		"https://vk.com/dmirou",
		"https://google.com",
	}, 1*time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := m.Run(ctx); err != nil {
		fmt.Printf("run: %v\n", err)
	}
}
