package main

import (
	"log"
	"service-move-to-school/config"
	"strings"
	"time"

	pq "github.com/pasiol/gopq"
)

var (
	// Version for build
	Version string
	// Build for build
	Build                   string
	jobName                 = "service-move-to-school"
	debugState              = false
	accountToRemove         = ""
	accountToMoveConfig     = ""
	accountToArchieveConfig = ""
)

func removeAccounts() ([]string, error) {
	query := config.RemoveUserAccountsQuery()
	c := config.GetPrimusConfig()
	query.Host = c.PrimusHost
	query.Port = c.PrimusPort
	query.User = c.PrimusUser
	query.Pass = c.PrimusPassword
	output, err := pq.ExecuteAndRead(query, 30)
	if err != nil {
		return nil, err
	}
	if output == "" {
		log.Print("No data, nothing to do.")
	}
	return strings.Fields(output), nil
}

func moveToSchool() ([]string, error) {
	query := config.MoveToSchoolQuery()
	c := config.GetPrimusConfig()
	query.Host = c.PrimusHost
	query.Port = c.PrimusPort
	query.User = c.PrimusUser
	query.Pass = c.PrimusPassword
	output, err := pq.ExecuteAndRead(query, 30)
	if err != nil {
		return nil, err
	}
	if output == "" {
		log.Print("No data, nothing to do.")
	}
	return strings.Fields(output), nil
}

func archieveApplicant() ([]string, error) {
	query := config.ArchieveApplicantQuery()
	c := config.GetPrimusConfig()
	query.Host = c.PrimusHost
	query.Port = c.PrimusPort
	query.User = c.PrimusUser
	query.Pass = c.PrimusPassword
	output, err := pq.ExecuteAndRead(query, 30)
	if err != nil {
		return nil, err
	}
	if output == "" {
		log.Print("No data, nothing to do.")
	}
	return strings.Fields(output), nil
}

func main() {

	start := time.Now()
	c := config.GetPrimusConfig()

	log.Printf("Starttime: %v", start.Format(time.RFC3339))
	log.Printf("Starting job: %s", jobName)
	log.Println("Version: ", Version)
	log.Println("Build Time: ", Build)
	pq.Debug = debugState

	accountsToRemove, _ := removeAccounts() // TODO err
	for i, account := range accountsToRemove {
		if debugState {
			log.Printf("Processing row %d: %v.", i, account)
		}
		id := account
		outputFile, err := config.RemoveUserAccountsXML(id)
		if err != nil {
			log.Printf("removing user error: %s", err)
		}
		pq.ExecuteImportQuery(outputFile, c.PrimusHost, c.PrimusPort, c.PrimusUser, c.PrimusPassword, accountToRemove)
	}

	accountsToMove, _ := moveToSchool() // TODO err

	for i, account := range accountsToMove {
		if debugState {
			log.Printf("Processing row %d: %v.", i, account)
		}
		id := account
		outputFile, err := config.MoveToSchoolXML(id)
		if err != nil {
			log.Printf("moving user error: %s", err)
		}
		pq.ExecuteImportQuery(outputFile, c.PrimusHost, c.PrimusPort, c.PrimusUser, c.PrimusPassword, accountToMoveConfig)
	}

	accountsToArchieve, _ := archieveApplicant() // TODO err
	for i, account := range accountsToArchieve {
		log.Printf("Processing row %d: %v.", i, account)
		id := account
		outputFile, err := config.ArchieveApplicantXML(id)
		if err != nil {
			log.Printf("archieve user error: %s", err)
		}
		pq.ExecuteImportQuery(outputFile, c.PrimusHost, c.PrimusPort, c.PrimusUser, c.PrimusPassword, accountToArchieveConfig)
	}

	t := time.Now()
	elapsed := t.Sub(start)

	log.Printf("Ending succesfully %s.", jobName)
	log.Printf("Endtime: %v", t)
	log.Printf("Elapsed processing time %d.", elapsed)
}
