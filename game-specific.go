package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type PowerUpRobot struct {
	TeamNumber int
	Autonomous PowerUpAuton
	TeleOperated PowerUpTeleOp
	Endgame PowerUpEndgame
	OverallRobot RobotPerformance
}

type PowerUpAuton struct {
	Moved bool
	SwitchAttempted int
	SwitchScored int
	ExchangeAttempted int
	ExchangeScored int
	ScaleAttempted int
	ScaleScored int
}

type PowerUpTeleOp struct {
	ExchangeAttempted int
	ExchangeScored int
	OwnSwitchAttempted int
	OwnSwitchScored int
	ScaleAttempted int
	ScaleScored int
	ScaleDropped int
	OppSwitchAttempted int
	OppSwitchScored int
	CubesDropped int
}

type PowerUpEndgame struct {
	ClimbSetupTime int
	LiftedSelf bool
	PartnersLifted int
	Parked bool
	DroppedPartner bool
	LiftedByPartner bool
	PartnersAttempted int
}

type RobotPerformance struct {
	StoppedMoving bool
	NeverMoved bool
	NoShow bool
	Defense int
}

func writeScoutedRobot(robot *PowerUpRobot, c *gin.Context) {

	ctx := appengine.NewContext(c.Request)

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "robot", nil), &robot)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 Employee
	if err = datastore.Get(ctx, key, &e2); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(c.Writer, "Stored and retrieved the Employee named %q", e2.Name)
}