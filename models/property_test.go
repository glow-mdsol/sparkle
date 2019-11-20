package models

import "testing"

func TestProcessProperties(t *testing.T) {
	rr := processProperties()
	expected := 104
	if expected != len(rr) {
		t.Error("Expected ", expected, " records, got ", len(rr))
	}
	r := rr[0]
	if "A1" != r.ConceptID {
		t.Error("Expected ConceptID to be A1, got ", r.ConceptID)
	}
	if "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A1" != r.URI {
		t.Error("Expected URI to be http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A1, got ", r.URI)
	}
	if "Role_Has_Domain" != r.Label {
		t.Error("Expected Label to be Role_Has_Domain, got ", r.Label)
	}
}
