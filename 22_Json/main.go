package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"The_name_of_the_course"`
	Price    int
	Platform string
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("This is the lecture on Json data")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	var slice = []course{{"Webd", 99, "online", []string{"a", "b", "c"}}, {"Aiml", 199, "online", []string{"b", "c"}}, {"ArVr", 599, "offline", nil}}
	finalJson, _ := json.MarshalIndent(slice,"","\t")
	convertedJson := string(finalJson)
	fmt.Printf("%s\n", convertedJson)
}
func decodeJson(){
	var my_new_course course; 
	// we create a dummy json data 
	jsonData := []byte(`
		{
			"The_name_of_the_course": "Webd",
			"Price": 99,
			"Platform": "online",
			"tags": [ "a", "b", "c"]
        }
	`)

	// ---------- 1. Consuming JSON data and printing it in the form of a strcuture -----------
	// we check if the json value is a valid json data 
	validJson := json.Valid(jsonData); 

	if (validJson){
		fmt.Println("JSON Valid!");
		// consuming json data into the my_new_course by passing it with reference. 
		json.Unmarshal(jsonData,&my_new_course);
		fmt.Printf("%#v\n",my_new_course); /// this special format is used to print a structure 
	}else{
		fmt.Println("INVALID JSON")
	}

	// ---------- 2. Consuming JSON Data and printing it in thr form of a map with KEY-VALUE PAIR

	var new_container map[string]interface{}; 
	
	json.Unmarshal(jsonData,&new_container); 
	fmt.Printf("%#v\n",new_container);
}
