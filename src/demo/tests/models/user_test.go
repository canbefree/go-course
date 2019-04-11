package models

import (
	"demo/models"
	"testing"
)

func Test(t *testing.T) {
	u := models.User{"小明", 18}
	if u.GetName() != "小明" {
		t.Errorf("name error %s", u.GetName())
	}
}
