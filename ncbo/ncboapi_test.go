package ncbo

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func loadFixtureFile(fixtureFileName string) (io.ReadCloser) {
	filename := filepath.Join("fixtures", fixtureFileName)
	file, _ := ioutil.ReadFile(filename)
	reader := bytes.NewReader(file)
	return ioutil.NopCloser(reader)
}

func TestExtractMapping(t *testing.T) {
	content := loadFixtureFile("mapping.json")
	coded := ExtractMapping(content)
	expected := 207
	if len(coded) != expected {
		t.Error("Expected ", expected, " entries, got ", len(coded))
	}
}

func TestExtractResponse(t *testing.T) {
	content := loadFixtureFile("loinc_query.json")
	coded := ExtractResponse(content)
	expected := 340
	if coded.PageCount != expected {
		t.Error("Expected PageCount to be ", expected, " got ", coded.PageCount)
	}
}

func TestExtractResponseInspect(t *testing.T) {
	content := loadFixtureFile("loinc_query.json")
	coded := ExtractResponse(content)
	firstObject := coded.Collection[0]
	cui := []string{"C0205455"}
	if cui[0] != firstObject.CUI[0] {
		t.Error("Expected CUI to be ", cui, " got ", firstObject.CUI)
	}
	id := "http://purl.bioontology.org/ontology/LNC/LA10141-2"
	if id != firstObject.ID{
		t.Error("Expected ID to be ", id, " got ", firstObject.ID)
	}
	if false != firstObject.Provisional {
		t.Error("Expected Provisional to be false got ", firstObject.Provisional)
	}
}
