package umlsrest

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type UMLSConfiguration struct {
	APIKey        string
	UMLSTGTURL    string
	GrantDate     time.Time
	TGTExpiryDate time.Time
}

// Response
//<!DOCTYPE HTML PUBLIC \"-//IETF//DTD HTML 2.0//EN\">
//<html>
//	<head>
//		<title>201 Created</title>
//	</head>
//	<body>
//		<h1>TGT Created</h1>
//		<form action="https://utslogin.nlm.nih.gov/cas/v1/api-key/TGT-646472-MbAwCn1fPd26bBIBoetyF1YyxQMleRlUEyzDrD7FfdTaicJDjX-cas" method="POST">Service:
//			<input type="text" name="service" value="">
//			<br>
//			<input type="submit" value="Submit">
//		</form>
//	</body>
//</html>

func (u *UMLSConfiguration) GenerateTicket() {
	const UTSLogin = "https://utslogin.nlm.nih.gov"
	ur, _ := url.Parse(UTSLogin)
	ur.Path = "/cas/v1/tickets"
	requestURL := ur.String()
	response, err := http.PostForm(requestURL, url.Values{"apikey": {u.APIKey}})
	if err != nil {
		log.Error("Error requesting TGT: ", err)
	}
	form, err := DecodeTGTMessage(response.Body)
	u.UMLSTGTURL = form.Action
	u.GrantDate = time.Now()
	u.TGTExpiryDate = u.GrantDate.Add(8 * time.Hour)
	if err != nil {
		log.Error("Error getting content")
	}
}

func (u *UMLSConfiguration) getServiceTicket()([]byte){
	const SKSSvc = "http://umlsks.nlm.nih.gov"
	// check that the ticket has been
	if u.UMLSTGTURL == ""{
		log.Warn("Need to get TGT before Service Ticket")
		u.GenerateTicket()
	}
	// check for ticket expiry
	if time.Now().After(u.TGTExpiryDate) {
		log.Info("TGT has expired, reissuing")
		u.GenerateTicket()
	}
	response, err := http.PostForm(u.UMLSTGTURL, url.Values{"service": {SKSSvc}})
	if err != nil {

	}
	var ticket []byte
	_, err = response.Body.Read(ticket)
	return ticket
}
