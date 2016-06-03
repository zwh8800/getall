package ui

import (
	"fmt"
	"sync"

	"github.com/gosuri/uiprogress"

	"github.com/zwh8800/getall/event"
)

type progressData struct {
	name string
	bar  *uiprogress.Bar
}

var (
	progress        *uiprogress.Progress
	progressDataMap map[int]*progressData
	dataMut         sync.RWMutex
)

func init() {
	progressDataMap = make(map[int]*progressData)
	progress = uiprogress.New()
	progress.Start()
	event.Server.Subscribe(event.ProgressStart, startProgress)
	event.Server.Subscribe(event.ProgressUpdate, updateProgress)
	event.Server.Subscribe(event.ProgressFinish, finishProgress)
	event.Server.Subscribe(event.ProgressRetry, retryProgress)
	event.Server.Subscribe(event.ProgressError, errorProgress)
}

func getProgressData(id int) *progressData {
	dataMut.RLock()
	defer dataMut.RUnlock()
	return progressDataMap[id]
}

func setProgressData(id int, data *progressData) {
	dataMut.Lock()
	defer dataMut.Unlock()
	progressDataMap[id] = data
}

func deleteProgressData(id int) {
	dataMut.Lock()
	defer dataMut.Unlock()
	delete(progressDataMap, id)
}

func startProgress(id int, name string) {
	bar := progress.AddBar(100).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("Task %s", name)
	})
	setProgressData(id, &progressData{
		name: name,
		bar:  bar,
	})
}

func updateProgress(id int, progress int) {
	data := getProgressData(id)
	data.bar.Incr()
	data.bar.Set(progress)
}

func finishProgress(id int) {
	data := getProgressData(id)
	data.bar.Set(100)
	fmt.Fprintf(progress.Bypass(), "%s finished\n", data.name)

	deleteProgressData(id)
}

func retryProgress(id int) {

}

func errorProgress(id int) {

}
