package test

import (
	"log"
	"time"
)

type Test struct {
	//ID			int			`json:"afspraakId"`
	Status		string		`json:"afspraakStatus"`
	FinalResult	string		`json:"uiteindelijkeUitslag"`
	//Date		time.Time	`json:"afnamedatum"`
	//Results		[]Result	`json:"labuitslagen"`
}

type Result struct {
	OrderNumber	string		`json:"ordernummer"`
	Result		string		`json:"uitslag"`
	TestType	string		`json:"coronatestType"`
	Accuracy	int			`json:"betrouwbaarheidScore"`
	ResultDate	time.Time	`json:"uitslagdatum"`
	SentDate	time.Time	`json:"verstuurddatum"`
	Sent		bool		`json:"verstuurd"`
}

func GetInterestingMessage(tests []Test) string {
	var message string

	if len(tests) > 0 {
		lastTest := tests[len(tests)-1]
		if lastTest.Status == "AFGEROND" {
			message = "The result from your last test is: " + lastTest.FinalResult
		}
	} else {
		message = "There are no tests visible. Maybe something is wrong?"
	}

	log.Printf("Got new message: %s", message)

	return message
}