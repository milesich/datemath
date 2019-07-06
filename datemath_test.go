package datemath_test

import (
	"testing"
	"time"

	"github.com/wuzuf/datemath"

	. "github.com/onsi/gomega"
)

var now = time.Date(2019, 7, 6, 15, 4, 5, 6, time.UTC)
var start = time.Date(2019, 7, 6, 16, 4, 5, 6, time.UTC)

var datemathSuccess = []struct {
	in   string
	out  time.Time
	name string
}{
	{"now", now, "Right now"},
	{"now-1d", time.Date(2019, 7, 5, 15, 4, 5, 6, time.UTC), "Substract 1 day"},
	{"now-d", time.Date(2019, 7, 5, 15, 4, 5, 6, time.UTC), "Only amount"},
	{"now-1h", time.Date(2019, 7, 6, 14, 4, 5, 6, time.UTC), "Substract 1h"},
	{"now+1h", time.Date(2019, 7, 6, 16, 4, 5, 6, time.UTC), "Add 1h"},
	{"now - 1h ", time.Date(2019, 7, 6, 14, 4, 5, 6, time.UTC), "Ignore whitespaces"},
	{"now-1h10s", time.Date(2019, 7, 6, 14, 3, 55, 6, time.UTC), "Substract Go duration"},
	{"now-1.5m", time.Date(2019, 7, 6, 15, 2, 35, 6, time.UTC), "Substract Go duration with optional fraction"},
	{"now-1w", time.Date(2019, 6, 29, 15, 4, 5, 6, time.UTC), "Substract 1 week"},
	{"now-1.5m/m", time.Date(2019, 7, 6, 15, 2, 0, 0, time.UTC), "Substract and round"},
	{"now-1.5m/m-2d", time.Date(2019, 7, 4, 15, 2, 0, 0, time.UTC), "Complex math"},
	{"now+1d/d-1m/d", time.Date(2019, 7, 6, 0, 0, 0, 0, time.UTC), "Other complex math"},
	{"now/1us", time.Date(2019, 7, 6, 15, 4, 5, 0, time.UTC), "Round to us"},
	{"now/6h", time.Date(2019, 7, 6, 12, 0, 0, 0, time.UTC), "Round to 6h"},
	{"now/86400s", time.Date(2019, 7, 6, 0, 0, 0, 0, time.UTC), "Round to day (expressed in seconds)"},
	{"now/d", time.Date(2019, 7, 6, 0, 0, 0, 0, time.UTC), "Round to day"},
	{"now/M", time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC), "Round to month"},
	{"now-M/3M", time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC), "Round to quarter"},
	{"now+3M/6M", time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC), "Round to semesters"},
	{"now/w", time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC), "Round to week"},
	{"now/w/w", time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC), "Round to week (double)"},
	{"now/4w", time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC), "Round to week"},
	{"now/y", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), "Round to year"},
	{"start/10m - m", time.Date(2019, 7, 6, 15, 59, 0, 0, time.UTC), "Round to 10m, use an identifier"},
	{"'2019-02-01'/3M", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC), "Explicit anchor date, round to quarter"},
	{"'2019-07-07 12:03:05'/w", time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC), "Explicit anchor date, round to quarter"},
}

func TestSuccess(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range datemathSuccess {
		t.Run(tt.in, func(t *testing.T) {
			ts, err := datemath.EvalWithEnv(tt.in, map[string]time.Time{"now": now, "start": start})
			Ω.Expect(err).To(Succeed(), tt.name+": '"+tt.in+"'")
			Ω.Expect(ts).To(BeTemporally("~", tt.out), tt.name+": '"+tt.in+"'")
		})
	}
}

var datemathError = []struct {
	in  string
	err string
}{
	{"!@#H", ":1:0:token recognition error at: '!'"},
	{"now-now", ":1:4:mismatched input 'now' expecting {EsDuration, GoDuration}"},
	{"now- 1 h", ":1:5:extraneous input '1' expecting {EsDuration, GoDuration}"},
	{"unknown-1h", "Unknown identifier: 'unknown'"},
}

func TestError(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range datemathError {
		t.Run(tt.in, func(t *testing.T) {
			_, err := datemath.Eval(tt.in)
			Ω.Expect(err).NotTo(Succeed())
			Ω.Expect(err.Error()).To(Equal(tt.err))
		})
	}
}
