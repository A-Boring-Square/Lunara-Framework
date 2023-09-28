package Lunara

import (
	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/scene"

)

func InitLunara() {
	oak.AddScene("myApp", scene.Scene{Start: func(ctx *scene.Context) {

	}})
	oak.Init("myApp")
}
