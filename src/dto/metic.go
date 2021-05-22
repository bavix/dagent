package dto

import (
	"sort"
	"strconv"
	"strings"
)

type MetricDto struct {
	Tags  map[string]string `json:"tags"`
	Name  string            `json:"name" validate:"required"`
	Value int64             `json:"value" validate:"required"`
}

func (m *MetricDto) UniqueId() string {
	return m.Name + m.tagsToString()
}

func (m *MetricDto) ToString() string {
	return m.UniqueId() + " " + strconv.FormatInt(m.Value, 10)
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
