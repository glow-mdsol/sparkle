package umlsrest

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestDecodeTGTMessage(t *testing.T) {
	message := `<!DOCTYPE HTML PUBLIC \"-//IETF//DTD HTML 2.0//EN\">
<html>
	<head>
		<title>201 Created</title>
	</head>
	<body>
		<h1>TGT Created</h1>
		<form action="https://utslogin.nlm.nih.gov/cas/v1/api-key/TGT-646472-MbAwCn1fPd26bBIBoetyF1YyxQMleRlUEyzDrD7FfdTaicJDjX-cas" method="POST">Service:
			<input type="text" name="service" value="">
			<br>
			<input type="submit" value="Submit">
		</form>
	</body>
</html>
`
	reader := ioutil.NopCloser(bytes.NewReader([]byte(message)))
	decoded, err := DecodeTGTMessage(reader)
	if err != nil {
		t.Error("Decoding failed", err)
	}
	expected := "https://utslogin.nlm.nih.gov/cas/v1/api-key/TGT-646472-MbAwCn1fPd26bBIBoetyF1YyxQMleRlUEyzDrD7FfdTaicJDjX-cas"
	if expected != decoded.Action {
		t.Error("Unable to get expected Action")
	}

}