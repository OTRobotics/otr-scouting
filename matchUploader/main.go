package main

import "github.com/otrobotics/otr-scouting"

type ClooneySheet2018 struct {
	MatchId  string
	Position string
}

func (s ClooneySheet2018) ToRobotTemplate() otrscouting.RobotTemplate {
	var upload otrscouting.RobotTemplate
	return upload
}

var (
	currentMatch *otrscouting.MatchUpload
)

func main() {

}

func parseClooneySheet(sheet ClooneySheet2018) {
	if sheet.MatchId == currentMatch.MatchId {
		switch sheet.Position {
		case "R1":
			currentMatch.MatchData.Red1 = sheet.ToRobotTemplate()
			break
		case "R2":
			currentMatch.MatchData.Red2 = sheet.ToRobotTemplate()
			break
		case "R3":
			currentMatch.MatchData.Red3 = sheet.ToRobotTemplate()
			break
		case "B1":
			currentMatch.MatchData.Blue1 = sheet.ToRobotTemplate()
			break
		case "B2":
			currentMatch.MatchData.Blue2 = sheet.ToRobotTemplate()
			break
		case "B3":
			currentMatch.MatchData.Blue3 = sheet.ToRobotTemplate()
			break
		}
	}
}
