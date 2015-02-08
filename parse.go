package nolo

// fixme: option values like key=2.3.4
func Parse(name, input string) *Meter {
	meter := new(Meter)
	meter.Identifier = name

	l := Lex(name, input)

	current_option_id := ""

Loop:
	for {
		item := l.nextItem()
		switch item.typ {
		case itemIdentifier:
			metric := new(Metric)
			metric.Identifier = item.val
			meter.Metrics = append(meter.Metrics, *metric)
		case itemValue:
			metric := &meter.Metrics[len(meter.Metrics)-1]
			metric.Value = item.val
		case itemOptionIdentifier:
			current_option_id = item.val
		case itemOptionValue:
			metric := &meter.Metrics[len(meter.Metrics)-1]
			if metric.Metadata == nil {
				metric.Metadata = make(map[string]string)
			}
			metric.Metadata[current_option_id] = item.val
		case itemEOF, itemError:
			break Loop
		}
	}

	return meter
}
