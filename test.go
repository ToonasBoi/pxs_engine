package pxs_engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var config struct {
	win_w int32
	win_h int32
	res_w int32
	res_h int32

	title string
}

func Config(win_w int32, win_h int32, res_w int32, res_h int32, title string) {
	config.win_w = win_w
	config.win_h = win_h
	config.res_w = res_w
	config.res_h = res_h
	config.title = title
}

func Launch(start_func func(), update_func func(float32), draw_func func()) {
	rl.InitWindow(config.win_w, config.win_h, config.title)

	start_func()

	for !rl.WindowShouldClose() {
		update_func(float32(rl.GetTime()))
		draw_func()
	}
}
