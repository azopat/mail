package health

import (
	"encoding/json"
	"fmt"
	"io"
	"bytes"
)

type JsonWriterSink struct {
	io.Writer
}

func (j *JsonWriterSink) EmitEvent(job string, event string, kvs map[string]string) {

	b, err := json.Marshal(struct {
		Job       string
		Event     string
		Timestamp string
		Kvs       map[string]string
	}{job, event, timestamp(), kvs})

	if err != nil {
		return
	}
	var bb bytes.Buffer
	bb.Write(b)
	bb.WriteRune('\r')
	bb.WriteRune('\n')
	j.Write(bb.Bytes())
}

func (j *JsonWriterSink) EmitEventErr(job string, event string, err error, kvs map[string]string) {

	b, err := json.Marshal(struct {
		Job       string
		Event     string
		Timestamp string
		Err       string
		Kvs       map[string]string
	}{job, event, timestamp(), fmt.Sprint(err), kvs})

	if err != nil {
		return
	}
	var bb bytes.Buffer
	bb.Write(b)
	bb.WriteRune('\r')
	bb.WriteRune('\n')
	j.Write(bb.Bytes())
}

func (j *JsonWriterSink) EmitTiming(job string, event string, nanoseconds int64, kvs map[string]string) {

	b, err := json.Marshal(struct {
		Job         string
		Event       string
		Timestamp   string
		Nanoseconds int64
		Kvs         map[string]string
	}{job, event, timestamp(), nanoseconds, kvs})

	if err != nil {
		return
	}
	var bb bytes.Buffer
	bb.Write(b)
	bb.WriteRune('\r')
	bb.WriteRune('\n')
	j.Write(bb.Bytes())
}

func (j *JsonWriterSink) EmitComplete(job string, status CompletionStatus, nanoseconds int64, kvs map[string]string) {

	b, err := json.Marshal(struct {
		Job         string
		Status      string
		Timestamp   string
		Nanoseconds int64
		Kvs         map[string]string
	}{job, status.String(), timestamp(), nanoseconds, kvs})

	if err != nil {
		return
	}
	var bb bytes.Buffer
	bb.Write(b)
	bb.WriteRune('\r')
	bb.WriteRune('\n')
	j.Write(bb.Bytes())
}
