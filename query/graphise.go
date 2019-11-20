package query

import "github.com/knakk/rdf"

type ConceptValue struct {
	Value      rdf.Term
	ValueLabel string
}

type ConceptProperty struct {
	Property rdf.Term
	Label    rdf.Term
	Values   []ConceptValue
}

type Concept struct {
	Subject    rdf.Term
	Properties map[rdf.Term]ConceptProperty
}


type CodeListItem struct {
	Term rdf.Term
	Code string
	PreferredName string
}

type CodeList struct {
	Term rdf.Term
	Code	string
	Name	string
	CodeListItems map[rdf.Term]CodeListItem
}
