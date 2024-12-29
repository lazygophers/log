package log_test

import (
	"github.com/elliotchance/pie/v2"
	"github.com/lazygophers/log"
	"os"
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	// fs , _, err := zap.Open("./out.log")
	// if err != nil {
	//    log.Errorf("err:%v", err)
	//    return
	// }

	log.SetTrace("")
	log.Info("msg")
	log.Infof("msgf")
	log.Infof("%d", 1)

	log.Clone().Caller(true).Info("not caller")
	log.Info("has caller")

	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		log.Info("go")
	}()

	w.Wait()
}

func TestFilename(t *testing.T) {
	files, err := os.ReadDir("D:\\Cache\\Temp")
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	pie.Each([]int{1, 2, 3, 4, 5}, func(i int) {
		t.Log(i)
	})

	pie.Each(
		pie.SortUsing(
			pie.Map(
				files,
				func(file os.DirEntry) string {
					return file.Name()
				},
			),
			func(a, b string) bool {
				return a > b
			},
		),
		func(s string) {
			t.Log(s)
		},
	)
}
