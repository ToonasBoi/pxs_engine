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

func Launch(launch_func func(), update_func func(float32), render_func func()) {
	rl.InitWindow(config.win_w, config.win_h, config.title)

	var curr_time float32 = 0
	var prev_time float32 = 0

	launch_func()

	for !rl.WindowShouldClose() {
		curr_time = float32(rl.GetTime())
		update_func(curr_time - prev_time)
		prev_time = curr_time

		rl.BeginDrawing()
		render_func()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
