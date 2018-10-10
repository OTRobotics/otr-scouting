package otrscouting

type PowerUpRobot struct {
	TeamNumber   int              `json:"team_number"`
	Autonomous   PowerUpAuton     `json:"autonomous"`
	TeleOperated PowerUpTeleOp    `json:"tele_operated"`
	Endgame      PowerUpEndgame   `json:"endgame"`
	OverallRobot RobotPerformance `json:"overall_robot"`
}

func (r *PowerUpRobot) Add(robot PowerUpRobot) {
	r.Endgame.Add(robot.Endgame)
	r.TeleOperated.Add(robot.TeleOperated)
	r.Autonomous.Add(robot.Autonomous)
}

type PowerUpAuton struct {
	Moved             bool `json:"moved"`
	SwitchAttempted   int  `json:"switch_attempted"`
	SwitchScored      int  `json:"switch_scored"`
	ExchangeAttempted int  `json:"exchange_attempted"`
	ExchangeScored    int  `json:"exchange_scored"`
	ScaleAttempted    int  `json:"scale_attempted"`
	ScaleScored       int  `json:"scale_scored"`
}

func (a *PowerUpAuton) Add(auton PowerUpAuton) {
	a.ExchangeScored += auton.ExchangeScored
	a.ExchangeAttempted += auton.ExchangeAttempted
	a.ScaleScored += auton.ScaleScored
	a.ScaleAttempted += auton.ScaleAttempted
	a.SwitchScored += auton.SwitchScored
	a.SwitchAttempted += auton.SwitchAttempted
}

type PowerUpTeleOp struct {
	ExchangeAttempted  int `json:"exchange_attempted"`
	ExchangeScored     int `json:"exchange_scored"`
	OwnSwitchAttempted int `json:"own_switch_attempted"`
	OwnSwitchScored    int `json:"own_switch_scored"`
	ScaleAttempted     int `json:"scale_attempted"`
	ScaleScored        int `json:"scale_scored"`
	ScaleDropped       int `json:"scale_dropped"`
	OppSwitchAttempted int `json:"opp_switch_attempted"`
	OppSwitchScored    int `json:"opp_switch_scored"`
	CubesDropped       int `json:"cubes_dropped"`
}

func (t *PowerUpTeleOp) Add(tele PowerUpTeleOp) {
	t.ScaleAttempted += tele.ScaleAttempted
	t.ScaleScored += tele.ScaleScored
	t.ExchangeAttempted += tele.ExchangeAttempted
	t.ExchangeScored += tele.ExchangeScored
	t.CubesDropped += tele.CubesDropped
	t.OppSwitchScored += tele.OppSwitchScored
	t.OppSwitchAttempted += tele.OppSwitchAttempted
	t.OwnSwitchScored += tele.OwnSwitchScored
	t.OwnSwitchAttempted += tele.OwnSwitchAttempted
	t.ScaleDropped += tele.ScaleDropped
}

type PowerUpEndgame struct {
	ClimbSetupTime    int  `json:"climb_setup_time"`
	LiftedSelf        bool `json:"lifted_self"`
	PartnersLifted    int  `json:"partners_lifted"`
	Parked            bool `json:"parked"`
	DroppedPartner    bool `json:"dropped_partner"`
	LiftedByPartner   bool `json:"lifted_by_partner"`
	PartnersAttempted int  `json:"partners_attempted"`
}

func (e *PowerUpEndgame) Add(endgame PowerUpEndgame) {
	e.PartnersLifted += endgame.PartnersLifted
	e.PartnersAttempted += endgame.PartnersAttempted
}

type RobotPerformance struct {
	StoppedMoving bool `json:"stopped_moving"`
	NeverMoved    bool `json:"never_moved"`
	NoShow        bool `json:"no_show"`
	Defense       int  `json:"defense"`
}
