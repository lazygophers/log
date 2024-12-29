package log

import (
	"github.com/google/uuid"
	"strings"
	"sync"

	"github.com/petermattis/goid"
)

var traceMap sync.Map

var DisableTrance bool

func getTrace(gid int64) string {
	tid, ok := traceMap.Load(gid)
	if !ok {
		return ""
	}
	return tid.(string)
}

func setTrace(gid int64, traceId string) {
	if DisableTrance {
		return
	}
	if traceId == "" {
		traceId = GenTraceId()
	}
	traceMap.Store(gid, traceId)
}

func delTrace(gid int64) {
	traceMap.Delete(gid)
}

func GetTrace() string {
	return getTrace(goid.Get())
}

func GetTraceWithGID(gid int64) string {
	return getTrace(gid)
}

func SetTrace(traceId ...string) {
	if len(traceId) > 0 {
		setTrace(goid.Get(), traceId[0])
		return
	}
	setTrace(goid.Get(), "")
}

func SetTraceWithGID(gid int64, traceId ...string) {
	if len(traceId) > 0 {
		setTrace(gid, traceId[0])
		return
	}
	setTrace(gid, "")
}

func DelTrace() {
	delTrace(goid.Get())
}

func DelTraceWithGID(gid int64) {
	delTrace(gid)
}

func GenTraceId() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")[16:]
}
