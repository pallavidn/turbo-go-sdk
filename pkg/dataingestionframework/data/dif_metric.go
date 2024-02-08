package data

import (
	"fmt"
)

type DIFMetric struct {
	MetricMap map[string]*DIFMetricVal
}

type DIFMetricVal struct {
	Average     *float64       `json:"average"`
	Min         *float64       `json:"min,omitempty"`
	Max         *float64       `json:"max,omitempty"`
	Capacity    *float64       `json:"capacity,omitempty"`
	Unit        *DIFMetricUnit `json:"unit,omitempty"`
	Key         *string        `json:"key,omitempty"`
	Resizable   *bool          `json:"resizable,omitempty"`
	Description *string        `json:"description,omitempty"`
	RawMetrics  interface{}    `json:"rawData,omitempty"`
}

type DIFMetricUnit string

const (
	COUNT DIFMetricUnit = "count"
	TPS   DIFMetricUnit = "tps"
	MS    DIFMetricUnit = "ms"
	MB    DIFMetricUnit = "mb"
	MHZ   DIFMetricUnit = "mhz"
	PCT   DIFMetricUnit = "pct"
)

type DIFMetricValKind string

const (
	KEY         DIFMetricValKind = "key"
	DESCRIPTION DIFMetricValKind = "description"
	RAWDATA     DIFMetricValKind = "rawData"
	AVERAGE     DIFMetricValKind = "average"
	MAX         DIFMetricValKind = "max"
	MIN         DIFMetricValKind = "min"
	CAPACITY    DIFMetricValKind = "capacity"
	UNIT        DIFMetricValKind = "unit"
)

const (
	UNSET_FLOAT  = 0.0
	UNSET_STRING = ""
)

func (m *DIFMetricVal) String() string {
	s := ""
	if m.Average != nil {
		s += fmt.Sprintf("Average:%v ", *m.Average)
	}
	if m.Capacity != nil {
		s += fmt.Sprintf("Capacity:%v ", *m.Capacity)
	}
	if m.Unit != nil {
		s += fmt.Sprintf("Unit:%v ", *m.Unit)
	}
	if m.Key != nil {
		s += fmt.Sprintf("Key:%v ", *m.Key)
	}
	if m.Resizable != nil {
		s += fmt.Sprintf("Resizable:%v", *m.Resizable)
	}
	return s
}

// copy creates a copy of this DIFMetricVal.
func (m *DIFMetricVal) copy() *DIFMetricVal {
	copy := DIFMetricVal{
		RawMetrics: m.RawMetrics,
	}
	if m.Average != nil {
		average := *m.Average
		copy.Average = &average
	}
	if m.Min != nil {
		min := *m.Min
		copy.Min = &min
	}
	if m.Max != nil {
		max := *m.Max
		copy.Max = &max
	}
	if m.Capacity != nil {
		capacity := *m.Capacity
		copy.Capacity = &capacity
	}
	if m.Unit != nil {
		unit := *m.Unit
		copy.Unit = &unit
	}
	if m.Key != nil {
		key := *m.Key
		copy.Key = &key
	}
	if m.Resizable != nil {
		resizable := *m.Resizable
		copy.Resizable = &resizable
	}
	if m.Description != nil {
		description := *m.Description
		copy.Description = &description
	}
	return &copy
}

// add adds the values of another DIFMetricVal to the current one.
// Currently, we only do adding the average and the capacity.  We also assume Unit/Key/Resizable are the same.
func (m *DIFMetricVal) add(another *DIFMetricVal) {
	// Average
	newAvg := *m.Average + *another.Average
	m.Average = &newAvg
	// Capacity
	newCapacity := *m.Capacity + *another.Capacity
	m.Capacity = &newCapacity
}
