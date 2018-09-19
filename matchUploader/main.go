package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/otrobotics/otr-scouting"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ClooneySheet2018 struct {
	AutoNoMove            int    `json:"auto_no_move"`
	AutoScoredExchange    int    `json:"auto_scored_exchange"`
	AutoScoredScale       int    `json:"auto_scored_scale"`
	AutoScoredSwitch      int    `json:"auto_scored_switch"`
	AutoWrongSwitch       int    `json:"auto_wrong_switch"`
	Disabled              int    `json:"disabled"`
	Filename              string `json:"filename"`
	Match                 int    `json:"match"`
	NoMove                int    `json:"no_move"`
	NoShow                int    `json:"no_show"`
	Notes                 string `json:"notes"`
	Pos                   int    `json:"pos"`
	MatchCode             string `json:"match_code"`
	RobotDrawing          string `json:"robot_drawing"`
	TeamNum               int    `json:"team_num"`
	TeamNumberEncoded     int    `json:"team_number_encoded"`
	TeleClimbedSelf       int    `json:"tele_climbed_self"`
	TeleDefense           int    `json:"tele_defense"`
	TeleDescoredScale     int    `json:"tele_descored_scale"`
	TeleDroppedCubes      int    `json:"tele_dropped_cubes"`
	TeleLiftedByPartner   int    `json:"tele_lifted_by_partner"`
	TeleParked            int    `json:"tele_parked"`
	TelePartnerDropped    int    `json:"tele_partner_dropped"`
	TelePartnersAttempted int    `json:"tele_partners_attempted"`
	TelePartnersLifted    int    `json:"tele_partners_lifted"`
	TeleScoredExchange    int    `json:"tele_scored_exchange"`
	TeleScoredOppSwitch   int    `json:"tele_scored_opp_switch"`
	TeleScoredOwnSwitch   int    `json:"tele_scored_own_switch"`
	TeleScoredScale       int    `json:"tele_scored_scale"`
	TeleSetupClimbTime    string `json:"tele_setup_climb_time"`
}

func (s ClooneySheet2018) ToRobotTemplate() otrscouting.RobotTemplate {
	var upload otrscouting.RobotTemplate
	if s.TeamNum == 0 && s.TeamNumberEncoded != 0 {
		upload.Team = s.TeamNumberEncoded
	} else {
		upload.Team = s.TeamNum
	}

	upload.MatchRef = s.MatchCode + "_" + strconv.Itoa(upload.Team)
	upload.PowerUp = otrscouting.PowerUpRobot{
		TeamNumber: upload.Team,
		Autonomous: otrscouting.PowerUpAuton{
			Moved:             s.AutoNoMove == 1,
			SwitchAttempted:   s.AutoScoredSwitch,
			SwitchScored:      s.AutoScoredSwitch,
			ExchangeAttempted: s.AutoScoredExchange,
			ExchangeScored:    s.AutoScoredExchange,
			ScaleAttempted:    s.AutoScoredScale,
			ScaleScored:       s.AutoScoredScale,
		},
		TeleOperated: otrscouting.PowerUpTeleOp{
			ScaleAttempted:     s.TeleScoredScale,
			ScaleScored:        s.TeleScoredScale,
			ExchangeAttempted:  s.TeleScoredExchange,
			ExchangeScored:     s.TeleScoredExchange,
			OwnSwitchAttempted: s.TeleScoredOwnSwitch,
			OwnSwitchScored:    s.TeleScoredOwnSwitch,
			OppSwitchAttempted: s.TeleScoredOppSwitch,
			OppSwitchScored:    s.TeleScoredOppSwitch,
			CubesDropped:       s.TeleDroppedCubes,
		},
		Endgame: otrscouting.PowerUpEndgame{
			ClimbSetupTime:    len(s.TeleSetupClimbTime),
			Parked:            s.TeleParked == 1,
			PartnersAttempted: s.TelePartnersAttempted,
			PartnersLifted:    s.TelePartnersLifted,
			LiftedByPartner:   s.TeleLiftedByPartner == 1,
			LiftedSelf:        s.TeleClimbedSelf == 1,
			DroppedPartner:    s.TelePartnerDropped == 1,
		},
		OverallRobot: otrscouting.RobotPerformance{
			StoppedMoving: s.NoMove == 1,
			NoShow:        s.NoShow == 1,
			Defense:       s.TeleDefense,
			NeverMoved:    s.NoMove == 1,
		},
	}
	return upload
}

func main() {
	// Create RobotTemplates from clooneysheets
	// Convert to matchUploads
	// UploadMatch via POST request to /upload with the data as an array of matchuploads

	dat, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic("Could not find data.json in the current directory.")
	}

	var sheets []ClooneySheet2018
	var matches = make(map[string][]ClooneySheet2018)
	var uploads []otrscouting.MatchUpload
	var toUpload []string

	json.Unmarshal(dat, &sheets)
	for _, sheet := range sheets {
		matchSheets := matches[sheet.MatchCode]
		matchSheets = append(matchSheets, sheet)
	}
	for matchCode, sheets := range matches {
		if len(sheets) != 6 {
			fmt.Printf("ERROR: Detected match %s with invalid number of sheets for upload. Counted %d sheets "+
				"instead of 6.", matchCode, len(sheets))
			continue
		}
		toUpload = append(toUpload, matchCode)

		match := otrscouting.MatchUpload{}
		match.Level = "qualifications"
		split := strings.Split(matchCode, "_")
		match.MatchId = split[2]
		match.EventId = split[1]
		year, _ := strconv.Atoi(split[0])
		match.Year = year
		for _, sheet := range sheets {
			switch sheet.Pos {
			case 1:
				match.MatchData.Red1 = sheet.ToRobotTemplate()
				break
			case 2:
				match.MatchData.Red2 = sheet.ToRobotTemplate()
				break
			case 3:
				match.MatchData.Red3 = sheet.ToRobotTemplate()
				break
			case 4:
				match.MatchData.Blue1 = sheet.ToRobotTemplate()
				break
			case 5:
				match.MatchData.Blue2 = sheet.ToRobotTemplate()
				break
			case 6:
				match.MatchData.Blue3 = sheet.ToRobotTemplate()
				break
			}
		}
		uploads = append(uploads, match)
	}
	data, err := json.Marshal(uploads)

	url := "http://otr-scouting.appspot.com/upload"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.PostForm.Add("json", string(data))
	req.PostForm.Add("pass", "OTR3474")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}
