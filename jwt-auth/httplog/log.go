package logging

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ApacheFormatPattern = "%s - - [%s] \"%s %d %d\" %f\n"
)

type ApacheLogRecord struct {
	http.ResponseWriter

	ip                    string
	time                  time.Time
	method, uri, protocol string
	status                int
	responseBytes         int64
	elapsedTime           time.Duration
}

func (r *ApacheLogRecord) Log(out io.Writer) {
	timeFormatted := r.time.Format("02/Jan/2006 03:04:05")
	requestLine := fmt.Sprintf("%s %s %s", r.method, r.uri, r.protocol)
	fmt.Fprintf(out, ApacheFormatPattern, r.ip, timeFormatted, requestLine, r.status, r.responseBytes,
		r.elapsedTime.Seconds())
}

func HumanizeDuration(t time.Duration) string {
	ti := t.Nanoseconds()
	if ti < 1000 {

		return strconv.FormatInt(ti, 10) + "ns"
	} else if ti > 1000 && ti < 1000000 {
		return strconv.FormatInt(ti/1000, 10) + "µs"
	} else if ti > 1000000 && ti < 1000000000 {
		return strconv.FormatInt(ti/1000000, 10) + "ms"
	} else {
		return strconv.FormatFloat(t.Seconds(), 'f', 6, 64) + "s"
	}
}

func (r *ApacheLogRecord) Write(p []byte) (int, error) {
	written, err := r.ResponseWriter.Write(p)
	r.responseBytes += int64(written)
	return written, err
}

func (r *ApacheLogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

type ApacheLoggingHandler struct {
	handler http.Handler
	out     io.Writer
}

func NewApacheLoggingHandler(handler http.Handler, out io.Writer) http.Handler {
	return &ApacheLoggingHandler{
		handler: handler,
		out:     out,
	}
}

func (h *ApacheLoggingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
		clientIP = clientIP[:colon]
	}

	record := &ApacheLogRecord{
		ResponseWriter: rw,
		ip:             clientIP,
		time:           time.Time{},
		method:         r.Method,
		uri:            r.RequestURI,
		protocol:       r.Proto,
		status:         http.StatusOK,
		elapsedTime:    time.Duration(0),
	}

	startTime := time.Now()
	h.handler.ServeHTTP(record, r)
	finishTime := time.Now()

	record.time = finishTime.UTC()
	record.elapsedTime = finishTime.Sub(startTime)

	record.Log(h.out)
}
