package stats

import (
	"testing"
	"time"
)

func TestEnd(t *testing.T) {
	st := New()

	start := time.Now()
	time.Sleep(1 * time.Second)
	st.End(start)

	expectedResponses := int64(1)

	if st.PasswordReponses != expectedResponses {
		t.Fatalf("Total Responses is wrong %v != %v", st.PasswordReponses, expectedResponses)
	}

	totalResponseTime := st.TotalResponseTime.Sub(time.Time{})
	totalDuration := time.Duration(totalResponseTime)
	totalMilliseconds := totalDuration.Nanoseconds() / 1000
	if totalMilliseconds < 1000000 {
		t.Fatalf("Total Time too fast %v", totalMilliseconds)
	}
}

func TestEndMultiple(t *testing.T) {
	st := New()

	for i := 0; i < 3; i++ {
		start := time.Now()
		wait := time.Duration(i)
		time.Sleep(wait * time.Second)
		st.End(start)
	}
	expectedResponses := int64(3)

	if st.PasswordReponses != expectedResponses {
		t.Fatalf("Total Responses is wrong %v != %v", st.PasswordReponses, expectedResponses)
	}

	totalResponseTime := st.TotalResponseTime.Sub(time.Time{})
	totalDuration := time.Duration(totalResponseTime)
	totalMilliseconds := totalDuration.Nanoseconds() / 1000
	if totalMilliseconds < 3000000 {
		t.Fatalf("Total Time too fast %v", totalMilliseconds)
	}
}

func TestAverageResponseTime(t *testing.T) {
	st := New()

	for i := 0; i < 3; i++ {
		start := time.Now()
		wait := time.Duration(i)
		time.Sleep(wait * time.Second)
		st.End(start)
	}
	expectedAverageResponseTime := int64(1000000)
	actual := st.AverageResponseTime()
	if actual < expectedAverageResponseTime {
		t.Fatalf("Average Response Time too fast %v", actual)
	}
}

func TestAverageResponseTimeNew(t *testing.T) {
	st := New()
	expectedAverageResponseTime := int64(0000000)
	actual := st.AverageResponseTime()
	if actual < expectedAverageResponseTime {
		t.Fatalf("Average Response time is not zero %v", actual)
	}
}
