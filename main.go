package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
	"math/rand"
	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/mattn/go-tty"
)

type Drop struct {
	x int
	y int
	symbol []string
	is_broken bool
}

type Land struct {
	x int 
	y int 
	symbol string
}

var LANDS = []Land{}
var DROPS = []Drop{}
var LANDS_COUNT = 1
var MAP_W, MAP_H = consolesize.GetConsoleSize()
var DROPS_COUNT = MAP_W*3
var HAND_POS_X = MAP_W/2
var HAND_POS_Y = MAP_H-1

func get_map() [][]string{
	MAP := make([][]string, MAP_H)
	for y := 0; y < MAP_H; y++{
    	MAP[y] = make([]string, MAP_W)

    	for x := 0; x < MAP_W; x++ {
    		MAP[y][x] = " "
		}
	}
	return MAP
}
func render_map() string{
	MAP := get_map()

	MAP_output := ""

	for i1, row := range MAP {
        for i2,_ := range row{
        	MAP[i1][i2] = MAP[i1][i2]
        }
	}
	for _, drop := range DROPS{
		for i, sym := range drop.symbol {
			MAP[drop.y][drop.x+i] = sym
		}
		
	}
	for _, land := range LANDS{
		MAP[land.y][land.x] = land.symbol
	}


	for i1, row := range MAP {
        for i2,_ := range row{
        	MAP_output += MAP[i1][i2]
        }
        MAP_output += "\n"
	}
	return MAP_output
}

func print_rendered_map(MAP string){
	fmt.Print(MAP)
}

func move_drops(){
	for i, drop := range DROPS{
		if !(drop.y >= MAP_H-1){
		DROPS[i].y = drop.y + 1
			
		}
	}
}

func lands_create_handler(){
	tty, _ := tty.Open()
    defer tty.Close()
    char, _ := tty.ReadRune()
    switch string(char){
     case "d":
     	if (HAND_POS_X < MAP_W-1){
	     	HAND_POS_X += 1
	        LANDS = append(LANDS, []Land{Land{x: HAND_POS_X, y: HAND_POS_Y, symbol: "#"}}...)
     	}

     
     case "w":
     	if !(HAND_POS_Y <= 0){
        	HAND_POS_Y -= 1
        	LANDS = append(LANDS, []Land{Land{x: HAND_POS_X, y: HAND_POS_Y, symbol: "#"}}...)
     	}
     
     case "a":
     	if !(HAND_POS_X <= 0){
	     	HAND_POS_X -= 1
	        LANDS = append(LANDS, []Land{Land{x: HAND_POS_X, y: HAND_POS_Y, symbol: "#"}}...)	
     	}

     
     case "s":
     	if (HAND_POS_Y < MAP_H-1){
     	     HAND_POS_Y += 1
        	LANDS = append(LANDS, []Land{Land{x: HAND_POS_X, y: HAND_POS_Y, symbol: "#"}}...)
     	}

     } 
    
}

func broke_drops(){
	for _, land := range LANDS{
		for drop_i, drop := range DROPS{
			if (land.y-1 == drop.y && land.x == drop.x || drop.y >= MAP_H-1){
				DROPS[drop_i].symbol = []string{"*", "'"}
				DROPS[drop_i].is_broken = true
			}
		}
	} 
}

func delete_broken_drops(){
	for i, drop := range DROPS{
		if (drop.is_broken){
			DROPS = DROPS[:i+copy(DROPS[i:], DROPS[i+1:])] 
			spawn_new_drop(3)
		}
	}
}
func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
    cmd.Run()
}

func spawn_new_drop(start_y int){
	y := rand.Intn(start_y - 0) + 0
	x := rand.Intn(MAP_W - 1) + 0
	DROPS = append(DROPS, Drop{x: x, y: y, symbol: []string{"|"}})

}
func spawn_new_land(){
	y := rand.Intn(MAP_H - 0) + 0
	x := rand.Intn(MAP_W - 0) + 0
	land := []Land{Land{x: x, y: y, symbol: "-"},
	Land{x: x+1, y: y, symbol: "-"},
	// Land{x: x+2, y: y, symbol: "-"},
	// Land{x: x+3, y: y, symbol: "-"},
	// Land{x: x+4, y: y, symbol: "-"},
	// Land{x: x+5, y: y, symbol: "-"},
	// Land{x: x+6, y: y, symbol: "-"},
	// Land{x: x+7, y: y, symbol: "-"},
	// Land{x: x+8, y: y, symbol: "-"},
	// Land{x: x+9, y: y, symbol: "-"},
	// Land{x: x+10, y: y, symbol: "-"},
	// Land{x: x+11, y: y, symbol: "-"},
	 }
	LANDS = append(LANDS, land...)

}
func init_drops(){
	for i := 0; i < DROPS_COUNT; i++ {
    	spawn_new_drop(MAP_H)
	}
}
func init_lands(){
	for i := 0; i < LANDS_COUNT; i++ {
    	spawn_new_land()
	}
}

func fps_pause(){
	time.Sleep(20 * time.Millisecond)
}

func raining(){
	move_drops()
	broke_drops()
	rendered_map := render_map()
	clear()
	print_rendered_map(rendered_map)
	lands_create_handler()
	fps_pause()
	delete_broken_drops()
	


	raining()
}
func main() {
	init_drops()
	init_lands()
	raining()

}