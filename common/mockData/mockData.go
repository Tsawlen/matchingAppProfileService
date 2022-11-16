package mockData

import (
	"app/matchingAppProfileService/common/dataStructures"
)

var ProfileData = []dataStructures.Profile{
	{Id: "1", Name: "Tomke Müller"},
	{Id: "2", Name: "Babett Müller"},
	{Id: "3", Name: "Mathis Neunzig"},
}

var UserData = []dataStructures.User{
	{City: "Homberg (Ohm)", Email: "jost-tomke-mueller@t-online.de",
		First_name: "Tomke", Name: "Müller", Password: "Test1234", Street: "Lichtenau", HouseNumber: "5",
		Username: "Seyna"},
}
