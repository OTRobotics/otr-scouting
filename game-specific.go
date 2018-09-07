package otrscouting

type PowerUpRobot struct {
	TeamNumber   int
	Autonomous   PowerUpAuton
	TeleOperated PowerUpTeleOp
	Endgame      PowerUpEndgame
	OverallRobot RobotPerformance
}

type PowerUpAuton struct {
	Moved             bool
	SwitchAttempted   int
	SwitchScored      int
	ExchangeAttempted int
	ExchangeScored    int
	ScaleAttempted    int
	ScaleScored       int
}

type PowerUpTeleOp struct {
	ExchangeAttempted  int
	ExchangeScored     int
	OwnSwitchAttempted int
	OwnSwitchScored    int
	ScaleAttempted     int
	ScaleScored        int
	ScaleDropped       int
	OppSwitchAttempted int
	OppSwitchScored    int
	CubesDropped       int
}

type PowerUpEndgame struct {
	ClimbSetupTime    int
	LiftedSelf        bool
	PartnersLifted    int
	Parked            bool
	DroppedPartner    bool
	LiftedByPartner   bool
	PartnersAttempted int
}

type RobotPerformance struct {
	StoppedMoving bool
	NeverMoved    bool
	NoShow        bool
	Defense       int
}
