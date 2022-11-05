package models

import (
	"fmt"
	"sync"
)

type LayerInfo struct {
	// Selected
	MutexSelected *sync.Mutex `json:"-"`
	Selected      uint        `json:"selected"`

	// Time out
	MutexCuttedByTimeOut *sync.Mutex `json:"-"`
	CuttedByTimeOut      uint        `json:"cutted_by_time_out"`

	// Pruned in depth
	MutexPrunedInDepth *sync.Mutex `json:"-"`
	PrunedInDepth      uint        `json:"pruned_in_depth"`

	// Pruned in width
	MutexPrunedInWidth *sync.Mutex `json:"-"`
	PrunedInWidth      uint        `json:"pruned_in_width"`

	// Cutted by max width
	MutexCuttedByMaxWidth *sync.Mutex `json:"-"`
	CuttedByMaxWidth      uint        `json:"cutted_by_max_width"`

	// Saved by min depth
	MutexSavedByMinDepth *sync.Mutex `json:"-"`
	SavedByMinDepth      uint        `json:"saved_by_min_depth"`
}

func (l *LayerInfo) IncrementSelected() {
	l.MutexSelected.Lock()
	l.Selected++
	l.MutexSelected.Unlock()
}

func (l *LayerInfo) IncrementCuttedByTimeOut() {
	l.MutexCuttedByTimeOut.Lock()
	l.CuttedByTimeOut++
	l.MutexCuttedByTimeOut.Unlock()
}

func (l *LayerInfo) IncrementPrunedInDepth() {
	l.MutexPrunedInDepth.Lock()
	l.PrunedInDepth++
	l.MutexPrunedInDepth.Unlock()
}

func (l *LayerInfo) IncrementPrunedInWidth() {
	l.MutexPrunedInWidth.Lock()
	l.PrunedInWidth++
	l.MutexPrunedInWidth.Unlock()
}

func (l *LayerInfo) IncrementCuttedByMaxWidth() {
	l.MutexCuttedByMaxWidth.Lock()
	l.CuttedByMaxWidth++
	l.MutexCuttedByMaxWidth.Unlock()
}

func (l *LayerInfo) IncrementSavedByMinDepth() {
	l.MutexSavedByMinDepth.Lock()
	l.SavedByMinDepth++
	l.MutexSavedByMinDepth.Unlock()
}

type Analyzer struct {
	Layers []LayerInfo `json:"layers"`
}

func (a *Analyzer) Init(depth uint8) {
	a.Layers = make([]LayerInfo, depth)

	for depth := range a.Layers {
		a.Layers[depth].MutexSelected = &sync.Mutex{}
		a.Layers[depth].MutexCuttedByTimeOut = &sync.Mutex{}
		a.Layers[depth].MutexPrunedInDepth = &sync.Mutex{}
		a.Layers[depth].MutexPrunedInWidth = &sync.Mutex{}
		a.Layers[depth].MutexCuttedByMaxWidth = &sync.Mutex{}
		a.Layers[depth].MutexSavedByMinDepth = &sync.Mutex{}
	}
}

func (a *Analyzer) Print() {
	fmt.Println("============================= [Analyzer]")

	for depth, layer := range a.Layers {
		fmt.Printf("--------------- [layer %d]\n", depth)

		layer.MutexSelected.Lock()
		fmt.Println("selected:\t\t\t", layer.Selected)
		layer.MutexSelected.Unlock()

		layer.MutexCuttedByTimeOut.Lock()
		fmt.Println("cutted by time out:\t\t\t", layer.CuttedByTimeOut)
		layer.MutexCuttedByTimeOut.Unlock()

		layer.MutexPrunedInDepth.Lock()
		fmt.Println("pruned in depth:\t\t", layer.PrunedInDepth)
		layer.MutexPrunedInDepth.Unlock()

		layer.MutexPrunedInWidth.Lock()
		fmt.Println("pruned in width:\t\t", layer.PrunedInWidth)
		layer.MutexPrunedInWidth.Unlock()

		layer.MutexCuttedByMaxWidth.Lock()
		fmt.Println("cutted by max in width:\t", layer.CuttedByMaxWidth)
		layer.MutexCuttedByMaxWidth.Unlock()

		layer.MutexSavedByMinDepth.Lock()
		fmt.Println("saved by min in depth:\t", layer.SavedByMinDepth)
		layer.MutexSavedByMinDepth.Unlock()
	}
}
