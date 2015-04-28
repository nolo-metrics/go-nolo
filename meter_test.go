package nolo

import (
	"reflect"
	"testing"
)

type ToMapTest struct {
	name     string
	input    MeterList
	expected MeterMap
}

var ToMapTests = []ToMapTest{
	{"empty",
		MeterList{&Meter{"empty", []Metric{}}},
		MeterMap{"empty": []map[string]string{}}},

	{"single",
		MeterList{
			&Meter{"single", []Metric{Metric{"guage", "99", MetricMetadata{}}}}},
		MeterMap{
			"single": []map[string]string{
				map[string]string{
					"identifier": "guage",
					"value":      "99"}}}},

	{"single with options",
		MeterList{
			&Meter{"single-with-options", []Metric{Metric{"guage", "99", MetricMetadata{"key": "value"}}}}},
		MeterMap{
			"single-with-options": []map[string]string{
				map[string]string{
					"identifier": "guage",
					"value":      "99",
					"key":        "value"}}}},

	{"multiple",
		MeterList{
			&Meter{"multiple",
				[]Metric{
					Metric{"guage", "99", MetricMetadata{}},
					Metric{"counter", "5", MetricMetadata{}}}}},
		MeterMap{
			"multiple": []map[string]string{
				map[string]string{
					"identifier": "guage",
					"value":      "99"},
				map[string]string{
					"identifier": "counter",
					"value":      "5"}}}},
}

func TestToMap(t *testing.T) {
	for _, test := range ToMapTests {
		actual := test.input.ToMeterMap()
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, actual, test.expected)
		}
	}
}

func TestMeterEquiv(t *testing.T) {
	first := Meter{"multiple", []Metric{
		Metric{"guage", "99", MetricMetadata{}},
		Metric{"counter", "5", MetricMetadata{}}}}
	second := Meter{"multiple", []Metric{
		Metric{"guage", "99", MetricMetadata{}},
		Metric{"counter", "5", MetricMetadata{}}}}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("Plugin Equivalance: got\n\t%v\nexpected\n\t%v", first, second)
	}
}

func TestMetricEquiv(t *testing.T) {
	first := Metric{"guage", "99", MetricMetadata{}}
	second := Metric{"guage", "99", MetricMetadata{}}
	if !reflect.DeepEqual(first, second) {
		t.Errorf("Metric Equivalance: got\n\t%v\nexpected\n\t%v", first, second)
	}
}
