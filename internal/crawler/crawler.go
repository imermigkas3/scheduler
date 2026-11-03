package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const baseURL = "https://registration.banner.gatech.edu/StudentRegistrationSsb/ssb/courseSearchResults/courseSearchResults"
const jSessionID = "E47C5EDC9BB1E8FE98A68E40D55458D7"
const bigIPCookie = "924551690.64288.0000" 

type CourseSearchResults struct {
	TotalCount int `json:"totalCount"` // this is called a struct tag. (a string of extra info attached to a struct field)
	Data []CourseData `json:"data"` // backticks are Go's syntax for a raw string meaning a string where nothing gets escaped
} 	

type CourseData struct {
	Subject string `json:"subject"`
	CourseNumber string `json:"courseNumber"`
	CourseTitle string `json:"courseTitle"`
	CourseDescription string `json:"courseDescription"`
	CreditHourLow *int `json:"creditHourLow"` // we use *int becuase the default for int is 0 but it might be nil and the default for *int is nil
	CreditHourHigh *int `json:"creditHourHigh"`
	CreditHourIndicator string `json:"creditHourIndicator"`
}

// what do i need:
// every course and its credit hours
// its pre and coreqs
// when each course is offered
// degree requirements per major
func FetchCourses(client *http.Client, offset int) (CourseSearchResults, error) {
// ?txt_term=202702&startDatepicker=&endDatepicker=&uniqueSessionId=b5r2z1793661376817&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
	params := url.Values{}
	params.Set("txt_term", "202702")
	params.Set("uniqueSessionId", "imcv41793666996191")
	params.Set("pageOffset", strconv.Itoa(offset))
	params.Set("pageMaxSize", "10")
	params.Set("sortColumn", "subjectDescription")
	params.Set("sortDirection", "asc")
	fullUrl := baseURL + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return CourseSearchResults{}, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("JSESSIONID=%s; BIGipServer~BANNER~registration.coda=%s",jSessionID, bigIPCookie))
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return CourseSearchResults{}, fmt.Errorf("sending bod: %w", err) 
	}

	defer resp.Body.Close() // we need to close the stream. defer is a keyword that guarantees this funciton will run after this current function finishes

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		snippet := string(body)
		if len(snippet) > 1000 {
			snippet = snippet[:1000]
		}

		return CourseSearchResults{}, fmt.Errorf("unexpected status: %s\nbody: %s", resp.Status, snippet)  // returns an error similar ot how fmt.Sprintf returns a string
	}

	var result CourseSearchResults
	err = json.NewDecoder(resp.Body).Decode(&result) // we need to pass a pointer to our struct
	if err != nil {
		return CourseSearchResults{}, fmt.Errorf("decoding response body: %w", err) 
	}

	return result, nil 
}