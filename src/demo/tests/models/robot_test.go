package models

import (
	"demo/models/robot"
	"testing"
)

func TestRobot(t *testing.T) {
	enrobot := robot.NewEnRobot()
	t.Fatal(enrobot.Name())
	t.Logf(enrobot.Name())
	t.Log("A")
}
