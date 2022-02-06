package grpcdates

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	local "github.com/scardozos/esplai-weeks-db/cmd/jsondates"
)

// Creates new dates server taking into account the database json filename
// Returns a model of a `DatesServer`
func NewDatesDBServer(jsonFilename string, startTime time.Time) *DatesServer {
	// get jsonFileMetadata. err is not nil if the file doesn't exist, or if
	// other errors arise when getting info from the file
	//log.Printf("Getting database file '%v'", jsonFilename)
	jsonFileMetadata, err := os.Stat(jsonFilename)
	// Catch errors
	if err != nil {
		// If error isn't file doesn't exist, print log with entire err msg
		if err != os.ErrNotExist {
			log.Printf("error getting file '%v': %v", jsonFilename, err)
			return nil
		}
		// Otherwise, print err msg stating that file wasn't found
		log.Printf("database json file '%v' not found", jsonFilename)
		return nil
	}

	// get jsonFilename's filesize. If it's empty, return nil and error log
	if jsonFileMetadata.Size() == 0 {
		log.Printf("error retrieving data from file: '%v' is empty", jsonFilename)
		return nil
	}
	jsonFile, err := ioutil.ReadFile(jsonFilename)
	//log.Print("Reading jsonFile")
	// if there's an error reading the json database file
	// return nil as the server could not be inititialized
	if err != nil {
		log.Printf("error reading file: %v", err)
		return nil
	}

	// set return object as models.DatesServer
	//log.Print("Initializing DatesServer")
	retObj := &DatesServer{
		WeeksList:    local.WeeksJsonList{},
		JsonFilename: jsonFilename,
	}

	// append items from database to local memory list
	err = retObj.WeeksList.FromJsonFile(jsonFile)
	// if there's an error adding items to WeeksList from jsonFile
	// return nil as the server could not be initialized
	if err != nil {
		log.Printf("error loading json file: %v", err)
		return nil
	}
	log.Printf("Started DatesServer successfully. Took %v", time.Since(startTime))
	return retObj
}
