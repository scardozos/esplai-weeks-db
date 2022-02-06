package jsondates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	weeks_pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"
)

var (
	ErrorDateAlreadyPresent error = fmt.Errorf("date already present")
)

type WeeksJsonList struct {
	Weeks []*WeekJsonModel `json:"weeks"`
}

type WeekJsonModel struct {
	*DateJsonModel `json:"date"`
}

type DateJsonModel struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

// Convert DateJsonModel to proto Date
func (d *DateJsonModel) ToGrpcDate() *weeks_pb.Date {
	return &weeks_pb.Date{Year: d.Year, Month: d.Month, Day: d.Day}
}

// Convert DateJsonModel to WeekJsonModel
func (d *DateJsonModel) ToWeekJsonModel() *WeekJsonModel {
	return &WeekJsonModel{d}
}

// Convert WeekJsonModel to proto Date
func (w *WeekJsonModel) ToGrpcWeek() *weeks_pb.Date {
	return w.ToGrpcDate()
}

// Unmarshal data from json file (database) and store data in memory
func (w *WeeksJsonList) FromJsonFile(jsonData []byte) error {
	// initialize data variable in which we'll be storing
	// the unmarshalled data from the json file
	data := WeeksJsonList{}

	// Unmarshal data from json file; if an error is encountered,
	// return an unmarshalling error
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("error unmarshalling: %v", err)
	}

	// If no errors, set WeeksJsonlist.Weeks to data.Weeks
	// and return nil
	w.Weeks = data.Weeks
	return nil
}

func (w *WeeksJsonList) AddToList(week WeekJsonModel) error {
	// check if date already exists in database
	if _, ok := w.IsInList(week); ok {
		// return error if it already exists
		log.Printf("Could not add week to list, date '%v' already in database", week.JsonWeekToString())
		return ErrorDateAlreadyPresent
	}
	// otherwise, add date to list
	// this change is only persisted in memory
	// to save it to the json file, it must be committed with weeksList.ToJsonFile()
	w.Weeks = append(w.Weeks, &week)
	log.Printf("Added week '%v' to database", week.JsonWeekToString())
	return nil
}

func (w *WeeksJsonList) RemoveFromList(week WeekJsonModel) error {
	// check if week is in database
	if idx, ok := w.IsInList(week); ok {
		// if week is present, remove week from list
		// Got this here: https://stackoverflow.com/a/37335777
		w.Weeks[idx] = w.Weeks[len(w.Weeks)-1]
		w.Weeks = w.Weeks[:len(w.Weeks)-1]
		log.Printf("Removed week '%v' from database", week.JsonWeekToString())
		return nil
	}
	// if not present, return error and log
	log.Printf("Failed to remove week from database, date '%v' not present", week.JsonWeekToString())
	return fmt.Errorf("date does not exist")
}

// check if the provided week is present in Weeks list
// and return its index. Otherwise return -1 as index
func (w *WeeksJsonList) IsInList(weekReq WeekJsonModel) (int, bool) {
	for idx, week := range w.Weeks {
		// we have to check the underlying DateJsonModel in order
		// to be able to compare the values themselves
		// rather than a pointer
		if *weekReq.DateJsonModel == *week.DateJsonModel {
			// return the found index, and true as the provided
			// week is present in the database
			return idx, true
		}
	}
	// return -1 as index doesn't exist
	// and return false as provided week isn't in database
	return -1, false
}

// Converts local WeeksJsonList to proto dates
func (w *WeeksJsonList) ToGrpcWeeksList() []*weeks_pb.Date {
	// initialize returnSlice variable with the length of w.Weeks
	returnSlice := make([]*weeks_pb.Date, len(w.Weeks))

	// iterate through w.Weeks and append to returnSlice
	// every week in proto format
	for idx, week := range w.Weeks {
		returnSlice[idx] = week.ToGrpcWeek()
	}
	// when done iterating return slice with proto weeks
	return returnSlice
}

// Persist changes from local memory to JSON file
func (w *WeeksJsonList) ToJsonFile(jsonFilename string) error {
	// store marshalled json in jsonFileContent
	//log.Printf("Marshalling list to json")
	jsonFileContent, err := json.Marshal(w)
	if err != nil {
		return fmt.Errorf("failed to marshal: %v", err)
	}

	// write files from jsonFileContent to json file jsonFilename
	//log.Printf("Persisting changes in '%v'", jsonFilename)
	if err = ioutil.WriteFile(jsonFilename, jsonFileContent, 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}

// Return string date-formatted WeekJsonModel
func (week WeekJsonModel) JsonWeekToString() string {
	return fmt.Sprintf("%02d-%02d-%d", week.Day, week.Month, week.Year)
}
