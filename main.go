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

type Brick struct {
	x int 
	y int 
	symbol string
}

//Bricks settings
var BRICKS = []Brick{}
var BRICKS_COUNT = 1

//Map setttings
var MAP_W, MAP_H = consolesize.GetConsoleSize()
var MAP_BACKGROUND = " "

//Drops settings
var DROPS = []Drop{}
var DROPS_COUNT = MAP_W*3

//Hand settings
var HAND_POS_X = MAP_W/2
var HAND_POS_Y = MAP_H-1

func get_map() [][]string{
	MAP := make([][]string, MAP_H)
	for y := 0; y < MAP_H; y++{
    	MAP[y] = make([]string, MAP_W)

    	for x := 0; x < MAP_W; x++ {
    		MAP[y][x] = MAP_BACKGROUND
		}
	}
	return MAP
}

func add_brick(x int, y int, symbol string){
	brick := []Brick{Brick{x: x, y: y, symbol: symbol},}
	BRICKS = append(BRICKS, brick...)
}

func render_map() string{
	MAP := get_map()

	MAP_output := ""

	for _, drop := range DROPS{
		for i, sym := range drop.symbol {
			MAP[drop.y][drop.x+i] = sym
		}
		
	}
	for _, brick := range BRICKS{
		MAP[brick.y][brick.x] = brick.symbol
	}

	MAP[HAND_POS_Y][HAND_POS_X] = "O"


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

func hand_handler(){
	tty, _ := tty.Open()
    defer tty.Close()
    char, _ := tty.ReadRune()
    switch string(char){
     case "d":
     	if (HAND_POS_X < MAP_W-1){
	     	HAND_POS_X += 1
			add_brick(HAND_POS_X, HAND_POS_Y, "#" )
     	}

     case "w":
     	if !(HAND_POS_Y <= 0){
        	HAND_POS_Y -= 1
        	add_brick(HAND_POS_X, HAND_POS_Y, "#" )
     	}
     
     case "a":
     	if !(HAND_POS_X <= 0){
	     	HAND_POS_X -= 1
			add_brick(HAND_POS_X, HAND_POS_Y, "#" )	
     	}

     
     case "s":
     	if (HAND_POS_Y < MAP_H-1){
     	     HAND_POS_Y += 1
			add_brick(HAND_POS_X, HAND_POS_Y, "#" )
     	}

     } 
    
}

func broke_drops(){
	for _, brick := range BRICKS{
		for drop_i, drop := range DROPS{
			if (brick.y-1 == drop.y && brick.x == drop.x || drop.y >= MAP_H-1){
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
func spawn_new_brick(){
	y := rand.Intn(MAP_H - 0) + 0
	x := rand.Intn(MAP_W - 0) + 0
	add_brick(x, y, "-")

}
func init_drops(){
	for i := 0; i < DROPS_COUNT; i++ {
    	spawn_new_drop(MAP_H)
	}
}
func init_bricks(){
	for i := 0; i < BRICKS_COUNT; i++ {
    	spawn_new_brick()
	}
}

func fps_pause(){
	time.Sleep(20 * time.Millisecond)
}

func raining(){
	//Drops moving and broking
	move_drops()
	broke_drops()
    
    //Map render
	rendered_map := render_map()
	clear()
	print_rendered_map(rendered_map)

	//Hand drawing 
	hand_handler()

	fps_pause() 
	delete_broken_drops()
	raining() 
}

func main() {
	init_drops()
	init_bricks()
	raining()

}