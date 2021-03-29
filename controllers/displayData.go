package controllers

import (
	"github.com/gobuffalo/packr/v2"
	"projet_wiki/models"
	"html/template"
	"log"
	"net/http"
)
// This is an example of a controller creating and potentially treating data. It would be in the same  

//Packr allows us to embed static files and not having to worry about relative paths
var box = packr.New("templateBox", "../views")

func DisplayData(w http.ResponseWriter, r *http.Request) {
	//Here we create datas but we could fetch some treated ones
	message := models.Message{Description: "Description message, can be long, short, treeated or untreated.", Longer_description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse viverra mollis metus, quis consequat lectus lacinia vel. Donec et velit id magna iaculis facilisis. Pellentesque vehicula ultricies justo nec lobortis. Aliquam dignissim egestas sapien eget volutpat. Nunc luctus venenatis risus ultrices hendrerit. ",
	}

	//Here we create the template
	htmlTemplate := template.New("displayData.html") 
	displayData, error := box.FindString("displayData.html")

	//Parse the template
	template, error := htmlTemplate.Parse(displayData)
	if error != nil {
		log.Print(error)
		return
	}

	//Merge it and it's done!
	error = template.Execute(w, message)
}