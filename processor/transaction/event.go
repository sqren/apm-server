package transaction

import (
	m "github.com/elastic/apm-server/processor/model"
	"github.com/elastic/apm-server/utility"
	"github.com/elastic/beats/libbeat/common"
)

type Event struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Type      string        `json:"type"`
	Result    *string       `json:"result"`
	Duration  float64       `json:"duration"`
	Timestamp string        `json:"timestamp"`
	Context   common.MapStr `json:"context"`
	Traces    []Trace       `json:"traces"`
}

func (t *Event) DocType() string {
	return "transaction"
}

func (t *Event) Transform() common.MapStr {
	enh := utility.NewMapStrEnhancer()
	tx := common.MapStr{"id": t.Id}
	enh.Add(tx, "name", t.Name)
	enh.Add(tx, "duration", utility.MillisAsMicros(t.Duration))
	enh.Add(tx, "type", t.Type)
	enh.Add(tx, "result", t.Result)
	return tx
}

func (t *Event) Mappings(pa *Payload) (string, []m.SMapping, []m.FMapping) {
	return t.Timestamp,
		[]m.SMapping{
			{Key: "processor.name", Value: processorName},
			{Key: "processor.event", Value: t.DocType()},
		}, []m.FMapping{
			{Key: t.DocType(), Apply: t.Transform},
			{Key: "context", Apply: func() common.MapStr { return t.Context }},
			{Key: "context.app", Apply: pa.App.Transform},
			{Key: "context.system", Apply: pa.System.Transform},
		}
}
