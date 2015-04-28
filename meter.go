package nolo

type MetricMetadata map[string]string

type Metric struct {
	Identifier string
	Value      string
	Metadata   MetricMetadata
}

type Meter struct {
	Identifier string
	Metrics    []Metric
}

type MeterList []*Meter
type MeterMap map[string][]map[string]string

func (m Metric) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range m.Metadata {
		result[k] = v
	}
	result["value"] = m.Value
	result["identifier"] = m.Identifier
	return result
}

func (ml MeterList) ToMeterMap() MeterMap {
	result := make(map[string][]map[string]string)
	for _, mr := range ml {
		length := len(mr.Metrics)
		values := make([]map[string]string, length)
		for i, mc := range mr.Metrics {
			values[i] = mc.ToMap()
		}
		result[mr.Identifier] = values
	}
	return result
}
