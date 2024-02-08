package log

import (
	"sync"

	"github.com/nats-io/nuid"
	"github.com/petermattis/goid"
)

var traceMap sync.Map

func getTrace(gid int64) []byte {
	tid, ok := traceMap.Load(gid)
	if !ok {
		return nil
	}
	return tid.([]byte)
}

func setTrace(gid int64, traceId string) {
	if traceId == "" {
		traceId = GenTraceId()
	}
	traceMap.Store(gid, []byte(traceId))
}

func delTrace(gid int64) {
	traceMap.Delete(gid)
}

func GetTrace() []byte {
	return getTrace(goid.Get())
}

func SetTrace(traceId ...string) {
	if len(traceId) > 0 {
		setTrace(goid.Get(), traceId[0])
		return
	}
	setTrace(goid.Get(), "")
}

func DelTrace() {
	delTrace(goid.Get())
}

var nid = nuid.New()

func GenTraceId() string {
	return nid.Next()[16:]
}
