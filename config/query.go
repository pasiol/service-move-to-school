package config

import (
	"log"
	"strconv"
	"time"

	"github.com/beevik/etree"
	pq "github.com/pasiol/gopq"
)

var (
	schoolName = "Ammattiopisto"
)

// RemoveUserAccountsQuery func
func RemoveUserAccountsQuery() pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = ""
	pq.Data = "#DATA{V1}"
	pq.Footer = ""

	return pq
}

// RemoveUserAccounts func
func MoveToSchoolQuery() pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = ""
	pq.Data = "#DATA{V1}"
	pq.Footer = ""

	return pq
}

// ArchieveApplicantQuery func
func ArchieveApplicantQuery() pq.PrimusQuery {
	pq := pq.PrimusQuery{}
	pq.Charset = "UTF-8"
	pq.Database = "opphenk"
	pq.Sort = ""
	pq.Search = ""
	pq.Data = "#DATA{V1}"
	pq.Footer = ""

	return pq
}

// RemoveUserAccountsXML generator
func RemoveUserAccountsXML(id string) (string, error) {
	xml := etree.NewDocument()
	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := xml.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "1")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "0")
	identity := xml.CreateElement("IDENTITY")
	identity.CreateText("some identity")
	card := xml.CreateElement("CARD")
	card.CreateAttr("FIND", id)
	waccount := card.CreateElement("WTUNNUS")
	waccount.CreateAttr("CMD", "DELETE")
	waccount.CreateAttr("LINE", "1")
	waccount.CreateText("")
	newTypeAccount := card.CreateElement("UUSITUNNUSKAYTOSSA")
	newTypeAccount.CreateText("Ei")
	xmlAsString, _ := xml.WriteToString()
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	if pq.Debug {
		log.Printf("\n%s", xmlAsString)
	}
	return filename, nil
}

// MoveToSchoolXML generator
func MoveToSchoolXML(id string) (string, error) {
	year := strconv.Itoa(time.Now().Year())[2:] // remember update this end of the year 9999
	xml := etree.NewDocument()
	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := xml.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "0")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "0")
	identity := xml.CreateElement("IDENTITY")
	identity.CreateText("some identity")
	card := xml.CreateElement("CARD")
	card.CreateAttr("FIND", id)
	school := card.CreateElement("KOULU")
	school.CreateAttr("CMD", "MODIFY")
	school.CreateAttr("LINE", "1")
	school.CreateText(schoolName)
	studentNumber := card.CreateElement("OPISKELIJANUMERO")
	studentNumber.CreateText(string(year + id))
	studentType := card.CreateElement("OPISKELIJALAJI")
	studentType.CreateText("Opiskelupaikan vastaanottanut")
	group := card.CreateElement("RYHMÄ")
	group.CreateAttr("CMD", "MODIFY")
	group.CreateText("123") // card number of group

	xmlAsString, _ := xml.WriteToString()
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	if pq.Debug {
		log.Printf("\n%s", xmlAsString)
	}
	return filename, nil
}

// ArchieveApplicantXML generator
func ArchieveApplicantXML(id string) (string, error) {
	xml := etree.NewDocument()
	xml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	primusquery := xml.CreateElement("PRIMUSQUERY_IMPORT")
	primusquery.CreateAttr("ARCHIVEMODE", "0")
	primusquery.CreateAttr("CREATEIFNOTFOUND", "0")
	identity := xml.CreateElement("IDENTITY")
	identity.CreateText("some identity")
	card := xml.CreateElement("CARD")
	card.CreateAttr("FIND", id)
	archieve := card.CreateElement("ARKISTO")
	archieve.CreateText("1")
	xmlAsString, _ := xml.WriteToString()
	filename, err := pq.CreateTMPFile(pq.StringWithCharset(128)+".xml", xmlAsString)
	if err != nil {
		return "", err
	}
	if pq.Debug {
		log.Printf("\n%s", xmlAsString)
	}
	return filename, nil
}
