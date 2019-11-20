package query

import (
	"bytes"
	"github.com/knakk/rdf"
	"github.com/knakk/sparql"
	log "github.com/sirupsen/logrus"
	"time"
)

const NAMESPACE = "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#"

const queries = `
# tag: concept-query
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

SELECT ?concept ?property ?label ?value
WHERE {
  ?concept ncit:P207 "{{.CUI}}" ;
     ?property ?value .
  OPTIONAL {?property rdfs:label ?label}
}

# tag: identify-property
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

SELECT ?property ?label ?value
WHERE {
  ncit:{{.PropertyURI}} ?property ?value .
  OPTIONAL {?property rdfs:label ?label} .
}

# tag: query-preferred-title
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

SELECT ?s ?o
WHERE {
  {
    ?s ncit:P108 ?o .
    FILTER REGEX(?o, "{{.QueryTerm}}") .
  }
  UNION
  {
    ?s ncit:P90 ?o .
    FILTER REGEX(?o, "{{.QueryTerm}}") . 
  }
}

# tag: describe-ncit-uri
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>

DESCRIBE ncit:{{.Concept}}

# tag: describe-uri
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>

DESCRIBE {{.URI}}

# tag: query-codelist
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

SELECT ?codelist ?codelistPreferredName ?codelistItem ?codelistItemPreferredName ?codelistItemCode
WHERE { ?codelist ncit:NHC0 "{{.Codelist}}" ;
            ncit:P108 ?codelistPreferredName .
  ?codelistItem ncit:A8 ?codelist ;
                ncit:P108 ?codelistItemPreferredName ;
				ncit:NHC0 ?codelistItemCode .
}

# tag: object-synonyms
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>

SELECT *
WHERE {
	ncit:{{.Term}} ncit:P90 ?synonyms .
}

# tag: code-query
PREFIX ncit: <http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#>
PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>

SELECT ?subject ?p ?label ?value ?valuelabel
WHERE {
	?subject ncit:NHC0 "{{.Code}}" ;
		?p ?value .
    OPTIONAL {?p rdfs:label ?label}
    OPTIONAL {?value ncit:P108 ?valuelabel }
}

# tag: mesh-term-to-gda
PREFIX mesh: <http://purl.bioontology.org/ontology/MESH/>	
PREFIX umls: <http://bioportal.bioontology.org/ontologies/umls/>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX sio:	<http://semanticscience.org/resource/>

SELECT DISTINCT *
WHERE {
	?s mesh:MN "C25.723.522.750" ;
    	umls:cui ?cui .
  BIND (CONCAT("umls:", ?cui) AS ?pfx ) .
  	?disease dcterms:identifier ?pfx .
    ?gda sio:SIO_000628 ?disease .    
}

# tag: mesh-term-to-ncbigene
PREFIX mesh: <http://purl.bioontology.org/ontology/MESH/>	
PREFIX umls: <http://bioportal.bioontology.org/ontologies/umls/>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX sio:	<http://semanticscience.org/resource/>

SELECT DISTINCT *
WHERE {
	?s mesh:MN "C25.723.522.750" ;
    	umls:cui ?cui .
  BIND (CONCAT("umls:", ?cui) AS ?pfx ) .
  	?disease dcterms:identifier ?pfx .
    ?gda sio:SIO_000628 ?disease, ?gene_or_variant . 
  	FILTER ( strstarts(str(?gene_or_variant), "http://identifiers.org/ncbigene/") )
}

# tag: mesh-term-to-variant
PREFIX mesh: <http://purl.bioontology.org/ontology/MESH/>	
PREFIX umls: <http://bioportal.bioontology.org/ontologies/umls/>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX sio:	<http://semanticscience.org/resource/>

SELECT DISTINCT *
WHERE {
	?s mesh:MN "C25.723.522.750" ;
    	umls:cui ?cui .
  BIND (CONCAT("umls:", ?cui) AS ?pfx ) .
  	?disease dcterms:identifier ?pfx .
    ?gda sio:SIO_000628 ?disease, ?gene_or_variant . 
  	FILTER ( strstarts(str(?gene_or_variant), "http://identifiers.org/dbsnp/") )
}

# tag: mesh-term-to-hp
PREFIX mesh: <http://purl.bioontology.org/ontology/MESH/>	
PREFIX umls: <http://bioportal.bioontology.org/ontologies/umls/>
PREFIX dcterms: <http://purl.org/dc/terms/>
PREFIX sio:	<http://semanticscience.org/resource/>

SELECT DISTINCT *
WHERE {
	?s mesh:MN "C25.723.522.750" ;
    	umls:cui ?cui .
  BIND (CONCAT("umls:", ?cui) AS ?pfx ) .
  	?disease dcterms:identifier ?pfx .
    ?gda sio:SIO_000628 ?disease, ?gene_or_variant . 
  	FILTER ( strstarts(str(?gene_or_variant), "http://purl.obolibrary.org/obo/HP_") )
}
`

type SPARQLConnection struct {
	ServerName string
	Username   string
	Password   string
}

func (c *SPARQLConnection) getRepo() (repo *sparql.Repo, err error) {
	repo, err = sparql.NewRepo(c.ServerName)
	repo.SetOption(sparql.Timeout(time.Millisecond * 60000))
	log.Info("ADDED SERVER: ", c.ServerName)
	// add auth if provided
	if c.Username != "" && c.Password != "" {
		repo.SetOption(sparql.DigestAuth(c.Username, c.Password))
	}
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (c *SPARQLConnection) GetBank(queries string) (bank sparql.Bank) {
	f := bytes.NewBufferString(queries)
	bank = sparql.LoadBank(f)
	log.Info("LOADED QUERYBANK FOR: ", c.ServerName)
	return
}

// Get a concept by CUI
func (c *SPARQLConnection) GetConceptByCUI(cui string) (*sparql.Results) {
	log.Info("LOOKING FOR CUI: ", cui, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("concept-query", struct {
		CUI string
	}{cui})
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

// Get the details for a Property
func (c *SPARQLConnection) GetPropertyDetails(property string) (*sparql.Results) {
	log.Info("LOOKING FOR PROPERTY: ", property, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("identify-property", struct {
		PropertyURI string
	}{property})
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

// Query a term
func (c *SPARQLConnection) TermBy(queryTerm string) (*sparql.Results) {
	log.Info("LOOKING FOR TERM: ", queryTerm, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("query-preferred-title", struct {
		QueryTerm string
	}{queryTerm})
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

// Describe a term
func (c *SPARQLConnection) DescribeTerm(term string) (*sparql.Results) {
	log.Info("DESCRIBING TERM: ", term, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("describe-uri", struct {
		Concept string
	}{term})
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

// Describe a ncit term
func (c *SPARQLConnection) DescribeNCITerm(ncit string) (*sparql.Results) {
	log.Info("DESCRIBING TERM: ", ncit, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("describe-ncit-uri", struct {
		Concept string
	}{ncit})
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

// Describe a Codelist
func (c *SPARQLConnection) GetCodeListByID(term string) (CodeList) {
	log.Info("DESCRIBING CODELIST: ", term, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("query-codelist", struct {
		Codelist string
	}{term})
	if err != nil {
		log.Fatal(err)
	}
	repo, err := c.getRepo()
	if err != nil {
		log.Fatal(err)
	}
	result, err := repo.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	var codelist CodeList
	for idx, res := range result.Solutions() {
		codelistTerm := res["codelist"]
		if idx == 0 {
			// create the Concept
			codelist = CodeList{
				Term:    codelistTerm,
				Code: term,
				Name: res["codelistPreferredName"].String(),
				CodeListItems: make(map[rdf.Term]CodeListItem),
			}
			log.Info("Created Codelist: ", codelist, "for", codelist)
		}
		var codelistitem CodeListItem
		if _, ok := codelist.CodeListItems[res["codelistItem"]]; ok {
			log.Info("Got existing CodeListItem for ", res["codelistItem"].String())
			codelistitem = codelist.CodeListItems[res["codelistItem"]]
		} else {
			log.Info("Creating new codelistItem for ", res["codelistItem"].String())
			codelistitem = CodeListItem{
				Term: res["codelistItem"],
				PreferredName: res["codelistItemPreferredName"].String(),
				Code: res["codelistItemCode"].String(),
			}
		}
		codelist.CodeListItems[res["codelistItem"]] = codelistitem
		//if res["p"].String() == "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#NHC0"{
		//	codelistitem.Code = res["o"].String()
		//} else if res["p"].String() == "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P108" {
		//	codelistitem.PreferredName = res["o"].String()
		//} else {
		//	var codelistItemProperty ConceptProperty
		//	if _, ok := codelistitem.Properties[res["p"]]; ok {
		//		log.Info("Got existing concept for ", res["p"].String())
		//		codelistItemProperty = codelistitem.Properties[res["p"]]
		//	} else {
		//		log.Info("Creating new concept for ", res["p"].String())
		//		codelistItemProperty = ConceptProperty{
		//			Property: res["p"],
		//			Label:    res["label"],
		//		}
		//	}
		//	value := res["value"]
		//	conceptValue := ConceptValue{
		//		Value: value,
		//	}
		//	if value.Type() == rdf.TermIRI {
		//		if res["valuelabel"] != nil{
		//			conceptValue.ValueLabel = res["valuelabel"].String()
		//		}
		//	}
		//
		//	conceptProperty.Values = append(conceptProperty.Values, conceptValue)
		//	concept.Properties[res["p"]] = conceptProperty
		//}
	}
	return codelist
}


// Describe a Code
func (c *SPARQLConnection) ByCode(code string) (Concept) {
	log.Info("DESCRIBING CODE: ", code, "FROM", c.ServerName)
	bank := c.GetBank(queries)
	q, err := bank.Prepare("code-query", struct {
		Code string
	}{code})
	if err != nil {
		log.Fatal(err)
	}
	repo, err := c.getRepo()
	if err != nil {
		log.Fatal(err)
	}
	result, err := repo.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	var concept Concept
	for idx, res := range result.Solutions() {
		subject := res["subject"]
		if idx == 0 {
			// create the Concept
			concept = Concept{
				Subject:    subject,
				Properties: make(map[rdf.Term]ConceptProperty),
			}
			log.Info("Created Concept: ", concept, "for", subject)
		}
		var conceptProperty ConceptProperty
		if _, ok := concept.Properties[res["p"]]; ok {
			log.Info("Got existing concept for ", res["p"].String())
			conceptProperty = concept.Properties[res["p"]]
		} else {
			log.Info("Creating new concept for ", res["p"].String())
			conceptProperty = ConceptProperty{
				Property: res["p"],
				Label:    res["label"],
			}
		}
		value := res["value"]
		conceptValue := ConceptValue{
			Value: value,
		}
		if value.Type() == rdf.TermIRI {
			if res["valuelabel"] != nil{
				conceptValue.ValueLabel = res["valuelabel"].String()
			}
		}

		conceptProperty.Values = append(conceptProperty.Values, conceptValue)
		concept.Properties[res["p"]] = conceptProperty
	}
	return concept
}
