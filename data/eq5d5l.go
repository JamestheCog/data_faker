// A module created to house the original EQ5D5L responses that were in a JSON file initially.  Dumba*s me completely forgot that the
// moment ./send.exe is moved out of this project's folder, the binary is going to start throwing errors because it has no idea where
// to find those said JSON files. So, I've decided that the data will just be stored in the form of a typed map:
//
// -- Kevin

package data

var EQ5D5L = map[string][]string{
	"mobility": {
		"I have no problems in walking about",
		"I have slight problems in walking about",
		"I have moderate problems in walking about",
		"I have severe problems in walking about",
		"I am unable to walk about",
	},
	"self_care": {
		"I have no problems washing or dressing myself",
		"I have slight problems washing or dressing myself",
		"I have moderate problems washing or dressing myself",
		"I have severe problems washing or dressing myself",
		"I am unable to wash or dress myself",
	},
	"usual_activities": {
		"I have no problems doing my usual activities",
		"I have slight problems doing my usual activities",
		"I have moderate problems doing my usual activities",
		"I have severe problems doing my usual activities",
		"I am unable to do my usual activities",
	},
	"pain_discomfort": {
		"I have no pain or discomfort",
		"I have slight pain or discomfort",
		"I have moderate pain or discomfort",
		"I have severe pain or discomfort",
		"I have extreme pain or discomfort",
	},
	"anxiety_depression": {
		"I am not anxious or depressed",
		"I am slightly anxious or depressed",
		"I am moderately anxious or depressed",
		"I am severely anxious or depressed",
		"I am extremely anxious or depressed",
	},
}
