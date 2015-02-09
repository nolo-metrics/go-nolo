package nolo

type Metric struct {
	Identifier string
	Value      string
	Metadata   map[string]string
}

type Meter struct {
	Identifier string
	Metrics    []Metric
}

type MeterList []*Meter
type MeterStringMap map[string][]map[string]string

func (m Metric) ToMap() map[string]string {
	result := make(map[string]string)
	for k, v := range m.Metadata {
		result[k] = v
	}
	result["value"] = m.Value
	result["identifier"] = m.Identifier
	return result
}

func (p Meter) ToStringMap() MeterStringMap {
	result := make(map[string][]map[string]string)
	length := len(p.Metrics)
	values := make([]map[string]string, length)
	for i, m := range p.Metrics {
		values[i] = m.ToMap()
	}
	result[p.Identifier] = values
	return result
}

func (ml MeterList) ToStringMap() MeterStringMap {
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
