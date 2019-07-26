package urlshort


import(
	 "testing"
	 "os"
	 "io/ioutil"
	 
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

func TestParseYaml(t *testing.T) {
	dir, err := os.Getwd()
	data, err := ioutil.ReadFile(dir + "/main/urls.yml")
		if err != nil {
			t.Errorf("This is a readfile error " + dir + "/main/urls.yml")
		}
	var config InstanceConfig
	config,_ = parseYAML(data)

	actualConfigIndexOne := "/urlshort"
	receivedConfigIndexOne := config.Path[0]

	if(actualConfigIndexOne != receivedConfigIndexOne) {
		t.Errorf("Config Expected: %s\n Config Received: %s\n",actualConfigIndexOne,receivedConfigIndexOne)
	}

	actualConfigIndexTwo := "/urlshort-final"
	receivedConfigIndexTwo := config.Path[1]

	if(actualConfigIndexTwo != receivedConfigIndexTwo) {
		t.Errorf("Config Expected: %s\n Config Received: %s\n",actualConfigIndexTwo,receivedConfigIndexTwo)
	}

}