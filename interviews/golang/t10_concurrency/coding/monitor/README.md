
```text
// Написать монитор, который периодически каждые requestFrequency
// запрашивает сайты sites и сохраняет статусы ответов в мапу результатов.
// Также он периодически выводит результаты на экран с помощью функции printStatuses.
// Монитор запускается функцией Run.

package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	// MaxGoroutines limits concurrent net requests. It's optional.
	MaxGoroutines = 1
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
}

func (m *Monitor) checkSite(ctx context.Context, site string) {
	// check site and write result in StatusMap
}

func (m *Monitor) printStatuses(ctx context.Context) error {
    // print results from StatusMap
}
```