package test

import (
	"fmt"
	"project02/model"
	"testing"
)

func TestHeroname(t *testing.T) {
	fmt.Println(model.Heroname)
}

func TestDbConnect(t *testing.T) {
	model.DbConnect()
}
