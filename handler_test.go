package urlshort


import(
	 "testing"
	 
)

type Page struct {
	Path []string
	URL []string
}


func TestBuildMap(t *testing.T) {
	var testPage = InstanceConfig{[]string{"/urlshort"},[]string{"https://github.com/gophercises/urlshort"}}
	pageMap := buildMap(testPage)
	urlReceived := pageMap["/urlshort"]
	actualURL := "https://github.com/gophercises/urlshort"

	if(urlReceived != actualURL) {
		t.Errorf("URL Expected: %s\n URL Received: %s\n",actualURL,urlReceived)
	}
}