package dto

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type MetricDto struct {
	Tags     map[string]string `json:"tags" validate:"gte=0,dive,keys,min=1,endkeys,min=1"`
	Name     string            `json:"name" validate:"required,min=1"`
	Value    *float64          `json:"value" validate:"required"`
	Duration *time.Duration    `json:"duration" validate:"omitempty,min=1s,max=2h"`
}

func (m *MetricDto) UniqueId() string {
	return m.Name + m.tagsToString()
}

func (m *MetricDto) ToString() string {
	return m.UniqueId() + " " + strconv.FormatFloat(*m.Value, 'G', 10, 64)
}

func (m *MetricDto) tagsToString() string {
	var arr []string
	for _, key := range m.getSortKeys() {
		arr = append(arr, key+"=\""+m.Tags[key]+"\"")
	}

	str := strings.Join(arr, ",")
	if str != "" {
		return "{" + str + "}"
	}

	return ""
}

func (m *MetricDto) getSortKeys() []string {
	var arr []string
	for tag := range m.Tags {
		arr = append(arr, tag)
	}

	sort.Strings(arr)

	return arr
}
