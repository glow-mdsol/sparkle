package ncbo

import (log "github.com/sirupsen/logrus")

const loinc_queries = `
# tag: loinc-properties
SELECT ?s ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/LNC/") AS ?symbol) 
WHERE {
	?s a owl:DatatypeProperty .
    FILTER(STRSTARTS(STR(?s), "http://purl.bioontology.org/ontology/LNC/"))
  OPTIONAL { ?s rdfs:label ?label ;
      rdfs:comment ?comment .
  } 
}

`

func GetLOINCCode(code, apiKey string)([]NCBOResult, error){
	ncbo := NCBOConfiguration{APIKey:apiKey}
	query := make(map[string]string)
	query["q"] = code
	query["require_exact_match"] = "true"
	query["ontologies"] = "LNC"
	response, err := ncbo.GetResult(query)
	if err != nil {
		log.Error(err)
		return nil, nil
	}
	return response, nil
}
