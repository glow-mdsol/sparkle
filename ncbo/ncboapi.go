package ncbo

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type NCBOConfiguration struct {
	APIKey  string
	BaseURL string
}

type Context struct {
	Vocab          string `json:"@vocab"`
	PreferredLabel string `json:"prefLabel"`
	Synonym        string `json:"synonym"`
	Obsolete       string `json:"obsolete"`
	SemanticType   string `json:"semanticType"`
	CUI            string `json:"cui"`
}

type Link struct {
	Self        string `json:"self"`
	Ontology    string `json:"ontology"`
	Children    string `json:"children"`
	Parents     string `json:"parents"`
	Descendants string `json:"descendants"`
	Ancestors   string `json:"ancestors"`
	Instances   string `json:"instances"`
	Tree        string `json:"tree"`
	Notes       string `json:"notes"`
	Mappings    string `json:"mappings"`
	UI          string `json:"ui"`
	Context		LinkContext	`json:"@context"`
}

type LinkContext struct {
	Self        string `json:"self"`
	Ontology    string `json:"ontology"`
	Children    string `json:"children"`
	Parents     string `json:"parents"`
	Descendants string `json:"descendants"`
	Ancestors   string `json:"ancestors"`
	Instances   string `json:"instances"`
	Tree        string `json:"tree"`
	Notes       string `json:"notes"`
	Mappings    string `json:"mappings"`
	UI          string `json:"ui"`
}

func (l *Link) GetOntologySymbol() (string) {
	split := strings.Split(l.Ontology, "/")
	return split[len(split)-1]
}

type NCBOResult struct {
	PreferredLabel string   `json:"prefLabel"`
	Synonyms       []string `json:"synonym"`
	CUI            []string `json:"cui"`
	SemanticTypes  []string `json:"semanticType"`
	Obsolete       bool     `json:"obsolete"`
	MatchType      string   `json:"matchType"`
	OntologyType   string   `json:"ontologyType"`
	Provisional    bool     `json:"provisional"`
	ID             string   `json:"@id"`
	Type           string   `json:"@type"`
	Links          Link     `json:"links"`
	Context        Context  `json:"@context"`
}

type NCBOResponse struct {
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
	PrevPage  int `json:"prevPage"`
	NextPage  int `json:"nextPage"`
	Links     struct {
		PrevPage string `json:"links.prevPage"`
		NextPage string `json:"links.nextPage"`
	}
	Collection []NCBOResult `json:"collection"`
}

type NCBOClass struct {
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	Links   Link   `json:"links"`
	Context struct {
		Vocab string `json:"@context.@vocab"`
	}
}

type NCBOMapping struct {
	MappingID string      `json:"id"`
	Source    string      `json:"source"`
	Process   string      `json:"process"`
	ID        string      `json:"@id"`
	Type      string      `json:"@type"`
	Classes   []NCBOClass `json:"classes"`
}

func ExtractMapping(closer io.ReadCloser) ([]NCBOMapping) {
	var mapping []NCBOMapping
	content, _ := ioutil.ReadAll(closer)
	err := json.Unmarshal(content, &mapping)
	if err != nil {
		log.Error("Decoding failed", err)
	}
	return mapping
}

func ExtractResponse(closer io.ReadCloser) NCBOResponse {
	var response NCBOResponse
	content, _ := ioutil.ReadAll(closer)
	err := json.Unmarshal(content, &response)
	if err != nil {
		log.Error("Decoding failed", err)
	}
	return response
}

// execute the query and parse the results
func getResult(request *http.Request) (*NCBOResponse, error) {
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		//proc := ld.NewJsonLdProcessor()
		//options := ld.NewJsonLdOptions("")
		//expanded, err := proc.Expand(resp.Body, options)

		if err != nil {
			log.Error("Error loading the JSON-LD; ", err)
			return nil, nil
		}
	} else {
		log.Error("Got ", resp.StatusCode, " when requesting ", request.URL)
	}
	return nil, nil

}

// collate a series of responses
func batchRequest(url, apiKey string) ([]NCBOResult, error) {
	var collection []NCBOResult

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Unable to build request")
	}
	req.Header.Add("Authorization", "apikey token="+apiKey)
	resp, err := getResult(req)
	// use the first as the step in
	for i := 1; i < resp.PageCount; i++ {
		if err != nil {
			break
		}
		// retrieve the elements
		for coll := range resp.Collection {
			collection = append(collection, resp.Collection[coll])
		}
		nextUrl := resp.NextPage
		if nextUrl == -1 {
			break
		}
		req, err := http.NewRequest("GET", resp.Links.NextPage, nil)
		if err != nil {
			panic("Unable to build request")
		}
		req.Header.Add("Authorization", "apikey token="+apiKey)
		resp, err = getResult(req)
	}
	return collection, nil
}

// Construct the query based on the input
func (n *NCBOConfiguration) GetResult(query map[string]string) ([]NCBOResult, error) {
	u, err := url.Parse(n.BaseURL)
	if err != nil {
		log.Error("Unable to parse BaseURL")
		return nil, err
	}
	q := u.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}
	results, err := batchRequest(u.String(), n.APIKey)
	return results, err
}
