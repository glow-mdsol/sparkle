package query

import (
	"github.com/knakk/sparql"
	log "github.com/sirupsen/logrus"
)

const ontology_properties  = `
# tag: get-datatype-properties
SELECT ?propuri ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/") AS ?symbol) 
WHERE {
	?s a owl:DatatypeProperty .
    FILTER(STRSTARTS(STR(?propuri), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/"))
  	OPTIONAL { ?propuri rdfs:label ?label ;
		rdfs:comment ?comment .
  	} 
}

# tag: get-objecttype-properties
SELECT ?propuri ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/") AS ?symbol) 
WHERE {
	?s a owl:ObjectProperty .
    FILTER(STRSTARTS(STR(?propuri), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/"))
  	OPTIONAL { ?propuri rdfs:label ?label ;
		rdfs:comment ?comment .
  	} 
}

# tag: get-annotation-properties
SELECT ?propuri ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/") AS ?symbol) 
WHERE {
	?s a owl:AnnotationProperty .
    FILTER(STRSTARTS(STR(?propuri), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/"))
  	OPTIONAL { ?propuri rdfs:label ?label ;
      	rdfs:comment ?comment .
  	} 
}

# tag: get-functional-properties
SELECT ?s ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/") AS ?symbol) 
WHERE {
	?propuri a owl:FunctionalProperty .
    FILTER(STRSTARTS(STR(?propuri), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/"))
  	OPTIONAL { ?propuri rdfs:label ?label ;
      	rdfs:comment ?comment .
  	} 
}

# tag: get-ontologies
SELECT ?ontology ?label ?comment ?version
WHERE {
	?ontology a owl:Ontology .
    OPTIONAL { ?ontology rdfs:label ?label ;
      rdfs:comment ?comment ;
      owl:versionInfo ?version .
  } 
}

# tag: get-ontology
SELECT ?s ?label ?comment (STRAFTER(STR(?s), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/") AS ?symbol) 
WHERE {
	?propuri a owl:FunctionalProperty .
    FILTER(STRSTARTS(STR(?propuri), "http://purl.bioontology.org/ontology/{{.ONTOLOGY}}/"))
  	OPTIONAL { ?propuri rdfs:label ?label ;
      	rdfs:comment ?comment .
  	} 
}

`

type Ontology struct {
	URI	string
	Label string
	Comment	string
	Version string
	Name 	string
}


// Get the Ontologies
func (c *SPARQLConnection) GetOntologies() (*sparql.Results) {
	log.Info("GETTING ONTOLOGIES FROM", c.ServerName)
	bank := c.GetBank(ontology_properties)
	q, err := bank.Prepare("get-ontologies")
	if err != nil {
		log.Fatal(err)
	}
	repo, err := c.getRepo()
	if err != nil {
		log.Fatal(err)
	}
	res, err := repo.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	return res
}


// Get the Properties of an Ontology by name
func (c *SPARQLConnection) BootstrapOntology(ontology string) (*sparql.Results) {
	log.Info("BOOTSTRAPPING ONTOLOGY: ", ontology, "FROM", c.ServerName)
	bank := c.GetBank(ontology_properties)
	q, err := bank.Prepare("datatype-query", struct {
		ONTOLOGY string
	}{ontology})
	if err != nil {
		log.Fatal(err)
	}
	repo, err := c.getRepo()
	if err != nil {
		log.Fatal(err)
	}
	res, err := repo.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
