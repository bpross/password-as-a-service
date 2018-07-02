package stats

import (
	"sync"
	"time"
)

type Stats struct {
	mutex             sync.RWMutex
	PasswordReponses  int64
	TotalResponseTime time.Time
}

func New() *Stats {
	stats := &Stats{
		PasswordReponses:  0,
		TotalResponseTime: time.Time{},
	}

	return stats
}

func (st *Stats) End(start time.Time) {
	end := time.Now()

	responseTime := end.Sub(start)
	st.mutex.Lock()
	defer st.mutex.Unlock()

	st.PasswordReponses++
	st.TotalResponseTime = st.TotalResponseTime.Add(responseTime)
}

func (st *Stats) AverageResponseTime() int64 {
	totalResponseTime := st.TotalResponseTime.Sub(time.Time{})
	averageResponseTime := time.Duration(0)
	if st.PasswordReponses > 0 {
		avgNs := int64(totalResponseTime) / int64(st.PasswordReponses)
		averageResponseTime = time.Duration(avgNs)
	}

	return averageResponseTime.Nanoseconds() / 1000
}
