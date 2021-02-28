package Controllers

import {
	"./models/types.go"
}

var UserList [2]User{
	{
		Username: "Cyrlax", 
		Password: "LaVieEnVioletPrunelle",
		Articles: [],
		Comments: [],	
	},
	{
		Username: "sdfsdf", 
		Password: "LaVieEnVioletPsdfqsdfelle",
		Articles: [],
		Comments: [],
	}
}

func UserHandler(responseWriter http.ResponseWriter, request *http.Request) {
	UserListJSON, err := json.Marshal(UserList)

	if err != nil {
		panic("Could not marshal json.")
	}

	fmt.Fprintf(responseWriter, string(UserListJSON))
}

