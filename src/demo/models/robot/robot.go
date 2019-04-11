package robot

type Robot interface {
	Name() string
}

type EnRobot struct {
	RobotID int
}

func NewEnRobot() Robot {
	return &EnRobot{
		RobotID: 1,
	}
}

func (r *EnRobot) Name() string {
	return "enrobot"
}
