package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glow-mdsol/sparkle/query"
	log "github.com/sirupsen/logrus"
)

// API is a defined as struct bundle
// for api. Feel free to organize
// your app as you wish.
type API struct{
	SPARQLServerName	string
}

// TODO: Add caching to the views

// Bind attaches api routes
func (api *API) Bind(group *gin.RouterGroup) {
	group.GET("/v1/conf", api.ConfHandler)
	group.GET("/concept/:concept", api.ConceptQuery)
	group.GET("/codelist/:codelist", api.CodeListQuery)
	group.GET("/property/:property", api.PropertyQuery)
	group.GET("/term/:term", api.TermQuery)
	group.GET("/code/:code", api.CodeQuery)
}

// ConfHandler handle the app config, for example
func (api *API) ConfHandler(c *gin.Context) {
	if app, ok := c.Get("app"); ok {
		c.IndentedJSON(200, app.(*App).Conf.AllSettings())
	} else {
		c.String(400, "False")
	}
}

// ConceptHandler
func (api *API) ConceptQuery(c *gin.Context){
	conceptCUI := c.Param("concept")
	log.Info("Calling Concept Query with", conceptCUI)
	server := query.SPARQLConnection{api.SPARQLServerName, "", ""}
	response := server.GetConceptByCUI(conceptCUI)
	c.IndentedJSON(200, response.Results)
}

// ConceptHandler
func (api *API) CodeListQuery(c *gin.Context){
	codelist := c.Param("codelist")
	log.Info("Calling Codelist Query with", codelist)
	server := query.SPARQLConnection{api.SPARQLServerName, "", ""}
	response := server.GetCodeListByID(codelist)
	c.IndentedJSON(200, response.CodeListItems)
}

// Query Term
func (api *API) TermQuery(c *gin.Context){
	term := c.Param("term")
	log.Info("Calling Term Query with", term)
	server := query.SPARQLConnection{api.SPARQLServerName, "", ""}
	response := server.TermBy(term)
	c.IndentedJSON(200, response.Results)
}

// Query Term
func (api *API) PropertyQuery(c *gin.Context){
	property := c.Param("property")
	log.Info("Calling Property Query with", property)
	server := query.SPARQLConnection{api.SPARQLServerName, "", ""}
	response := server.GetPropertyDetails(property)
	c.IndentedJSON(200, response.Results)
}

// Code Query
func (api *API) CodeQuery(c *gin.Context){
	code := c.Param("code")
	log.Info("Calling Code Query with", code)
	server := query.SPARQLConnection{api.SPARQLServerName, "", ""}
	concept := server.ByCode(code)
	c.IndentedJSON(200, concept)
}

