package main

import (
	"github.com/Jeffail/gabs"
)	

func GetAuthor() {

/*
"book":{
		"author": "J.R.R. Tolkien",
		"title": "Lord of the Rings"
	}

*/
			
	resp, err := client.Do(req)
	defer resp.Body.Close()
	mt.Println("response Status:", resp.StatusCode)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		json_queue, err := gabs.ParseJSON(body)
		if err != nil {
			panic(err)
		}
		build_url, build_found = json_queue.Path("book.author").Data().(string)
	}
}