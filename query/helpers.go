package query

import (
	"github.com/knakk/sparql"
	ld "github.com/piprate/json-gold/ld"
	log "github.com/sirupsen/logrus"
)

func toJSONLD(results *sparql.Results)(interface{}, error) {
	// Rewrite the results to add a redirect for the dereference of the element
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	doc, err := proc.FromRDF(results, options)
	if err != nil {
		log.Error("Unable to serialise results: ", err)
		return nil, err
	}
	return doc, nil
}
