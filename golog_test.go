package logbench

import (
	"fmt"
	"io"
	"time"

	"golang.org/x/exp/slog"
)

func init() {
	// zerolog.TimeFieldFormat = ""
	// zerolog.DurationFieldInteger = true
	// zerolog.MessageFieldName = "message"

	tests["golog"] = gologTester{}
}

// func (o obj) MarshalGologObject(e *golog.Event) {
// 	e.Str("name", o.Name).
// 		Int("count", o.Count).
// 		Bool("enabled", o.Enabled)
// }

type gologTester struct {
	l *slog.Logger
}

var (
	_ logTesterArray = (*gologTester)(nil)
)

func (gologTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := slog.LevelDebug
	if disabled {
		lvl = slog.Level(10)
	}
	return gologTester{slog.New(slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time from the output for predictable test output.
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))}

}

func (t gologTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t gologTester) logFormat(format string, v ...interface{}) bool {
	t.l.Info(fmt.Sprintf(format, v...))
	return true
}

func (t gologTester) withContext(context map[string]interface{}) (logTester, bool) {
	var values []interface{}
	for k, v := range context {
		values = append(values, k)
		values = append(values, v)
	}
	return gologTester{t.l.With(values...)}, true

}

func (t gologTester) logBool(msg, key string, value bool) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logInt(msg, key string, value int) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logFloat32(msg, key string, value float32) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logFloat64(msg, key string, value float64) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logTime(msg, key string, value time.Time) bool {
	// t.l.With(key, value).Info(msg)
	// return true
	return false
}

func (t gologTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logError(msg, key string, value error) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logString(msg, key string, value string) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logObject(msg, key string, value *obj) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logBools(msg, key string, value []bool) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logInts(msg, key string, value []int) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logFloats32(msg, key string, value []float32) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logFloats64(msg, key string, value []float64) bool {
	t.l.With(key, value).Info(msg)
	return true
}

func (t gologTester) logTimes(msg, key string, value []time.Time) bool {
	// t.l.With(key, value).Info(msg)
	// return true
	return false
}

func (t gologTester) logDurations(msg, key string, value []time.Duration) bool {
	// t.l.With(key, value).Info(msg)
	// return true
	return false
}

func (t gologTester) logErrors(msg, key string, value []error) bool {
	// t.l.With(key, value).Info(msg)
	// return true
	return false
}

func (t gologTester) logStrings(msg, key string, value []string) bool {
	// t.l.With(key, value).Info(msg)
	// return true
	return false
}
