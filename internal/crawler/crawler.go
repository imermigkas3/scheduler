package crawler

import "net/http"

type CourseSearchResults struct {
	TotalCount int `json:"totalcount"` // this is called a struct tag. (a string of extra info attached to a struct field)
	Data []CourseData `json:"data"` // backticks are Go's syntax for a raw string meaning a string where nothing gets escaped
} 	

type CourseData struct {
	Subject string `json:"subject"`
	CourseNumber string `json:"courseNumber"`
	CourseTitle string `json:"courseTitle"`
	CourseDescription string `json:"courseDescription"`
	CreditHourLow *int `json:"creditHourLow"` // we use *int becuase the default for int is 0 but it might be nil and the default for *int is nil
	CreditHourHigh *int `json:"creditHourHigh"`
	CreditHourIndicator *int `json:"creditHourIndicator"`
}

// what do i need:
// every course and its credit hours
// its pre and coreqs
// when each course is offered
// degree requirements per major
func FetchCourses(client *http.Client, offset int) (CourseData, error) {

}