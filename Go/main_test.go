package main

import (
	"fmt"

)

type Feature struct{
	response *http.response
	body io.Reader
}

func (f *Feature) Initialize(sc *godog.ScenarioContext) {
sc.beforeScenario(func(*godog.ScenarioContext){
	f.response = nil
	f.body = nil
})
sc.Step()

}