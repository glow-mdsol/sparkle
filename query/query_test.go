package query

import (
	"testing"
)

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func TestSPARQLConnection_ByCode(t *testing.T) {
	// Can't really stub this out
	r := SPARQLConnection{
		ServerName: "http://127.0.0.1:3030/ncit/sparql",
	}
	result := r.ByCode("C120665")
	expected := "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#C120665"
	if expected != result.Subject.String() {
		t.Error("Expected ", expected, " got ", result.Subject.String())
	}
	if 12 != len(result.Properties) {
		t.Error("Expected ", expected, " edges got ", len(result.Properties))

	}
}

func TestSPARQLConnection_GetCodeListByID(t *testing.T) {
	// Can't really stub this out
	r := SPARQLConnection{
		ServerName: "http://127.0.0.1:3030/ncit/sparql",
	}
	// CMDOSFRQ
	result := r.GetCodeListByID("C78419")
	expected := "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#C78419"
	if expected != result.Term.String() {
		t.Error("Expected ", expected, " got ", result.Term.String())
	}
	if 8 != len(result.CodeListItems) {
		t.Error("Expected ", expected, " edges got ", len(result.CodeListItems))
	}
	var expectedCodes = []string{"C17998", "C25473",
		"C64496", "C64498", "C64499", "C64525", "C64527", "C64530",
	}

	for _, cli := range result.CodeListItems {
		found := Contains(expectedCodes, cli.Code)
		if found != true {
			t.Error("Expected code ", cli.Code, " to be found")
		}
	}
}
