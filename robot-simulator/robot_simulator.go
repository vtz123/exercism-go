package robot

import (
	"fmt"
	"sync"
	"time"
)

//var tmp = []byte{'E','N','W','S'}
//var list = [][]int{ {0,1}, {1,0}, {0,-1}, {-1,0}}


const (
	N = iota
	E
	S
	W
)
// //   N
//// W   E
////   S
func (d Dir) String() string {
	return fmt.Sprintf("%d", d)
}

// step1
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y += 1
	case E:
		Step1Robot.X += 1
	case S:
		Step1Robot.Y -= 1
	case W:
		Step1Robot.X -= 1
	}
}

func Right() {
	Step1Robot.Dir += 1
	Step1Robot.Dir %= 4
}

func Left() {
	Step1Robot.Dir += 3
	Step1Robot.Dir %= 4
}

// step2
func Room(rect Rect, step2robot Step2Robot, act chan Action, rep chan Step2Robot) {

	for a := range act {
		switch a {
		case 'A':
			step2robot.Advance(rect)
		case 'R':
			step2robot.Right()
		case 'L':
			step2robot.Left()
		default:

		}
	}

	rep <- step2robot
}

func StartRobot(cmd chan Command, act chan Action) {
	for command := range cmd {
		act <- Action(command)
	}

	close(act)
}

type Action byte

func (stp2 *Step2Robot) Advance(rect Rect) {

	switch stp2.Dir {
	case N:
		if stp2.Pos.Northing < rect.Max.Northing {
			stp2.Pos.Northing += 1
		}
	case E:
		if stp2.Pos.Easting < rect.Max.Easting{
			stp2.Pos.Easting += 1
		}
	case S:
		if stp2.Pos.Northing > rect.Min.Northing {
			stp2.Pos.Northing -= 1
		}
	case W:
		if stp2.Pos.Easting > rect.Min.Easting {
			stp2.Pos.Easting -= 1
		}
	}

}

func (stp2 *Step2Robot ) Right() {
	stp2.Dir += 1
	stp2.Dir %= 4
}

func (stp2 *Step2Robot ) Left() {
	stp2.Dir += 3
	stp2.Dir %= 4
}

// step3
func (stp2 *Step2Robot) Advance2(rect Rect, log chan string, grid [][]int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()

	grid[stp2.Easting][stp2.Northing]--

	poscopy := stp2.Pos
	switch stp2.Dir {
	case N:
		if stp2.Pos.Northing < rect.Max.Northing -rect.Min.Northing{
			stp2.Pos.Northing += 1
		}else {
			log <- "*  A robot attempting to advance into a wall"
		}
	case E:
		if stp2.Pos.Easting < rect.Max.Easting - rect.Min.Easting {
			stp2.Pos.Easting += 1
		} else {
			log <- "*  A robot attempting to advance into a wall"
		}
	case S:
		if stp2.Pos.Northing > 0 {
			stp2.Pos.Northing -= 1
		} else {
			log <- "*  A robot attempting to advance into a wall"
		}
	case W:
		if stp2.Pos.Easting > 0 {
			stp2.Pos.Easting -= 1
		} else {
			log <- "*  A robot attempting to advance into a wall"
		}
	}

	if grid[stp2.Easting][stp2.Northing] >= 1 {
		stp2.Pos = poscopy
		grid[poscopy.Easting][poscopy.Northing]++
		log <- "*  A robot attempting to advance into another robot"
	}else {
		//grid[poscopy.Easting][poscopy.Northing]--
		grid[stp2.Easting][stp2.Northing]++
	}

}

func (stp2 *Step2Robot ) Control (rect Rect ,script string, grid [][]int, log chan string ){
	mu := &sync.Mutex{}

	for i := range script{
		switch script[i] {
		case 'A':
			stp2.Advance2(rect, log, grid, mu)
		case 'R':
			stp2.Right()
		case 'L':
			stp2.Left()
		default:
			log <- "*  An undefined command in a script"
			return
		}
	}
}

func StartRobot3(name, script string, action chan Action3, log chan string) {

	// script -> action
	if name == "" {
		//log <- "*  An action from an unknown robot"
	} else {
		action <- Action3{
			Name: name,
			Script:  script,
		}
	}

	//go func() {
	//	time.Sleep(500*time.Millisecond)
	//	if _, ok := <- action ; ok {
	//		close(action)
	//	}
	//
	//}()
	//
	//close(action)
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	tmp := make(map[string]int) // name used
	list := make(map[string]*Step3Robot) //
	grid := make([][]int, extent.Max.Easting+1 - extent.Min.Easting)
	for i := range grid{
		grid[i] = make([]int, extent.Max.Northing+1 - extent.Min.Northing)
	}
	//fmt.Println(len(grid), len(grid[0]))
	for i := range robots{
		// name
		robots[i].Step2Robot.Northing -= extent.Min.Northing
		robots[i].Step2Robot.Easting  -= extent.Min.Easting

		if robots[i].Name == "" {
			log <- "*  A robot without a name"
			//list["-"] = append(list["-"], &robots[i])
		}else {
			if _,ok := tmp[robots[i].Name]; ok{
				log <- "*  Duplicate robot names"
				continue
			} else {
				tmp[robots[i].Name]++
				list[robots[i].Name] = &robots[i]
			}
			//list[robots[i].Name] = append(list[robots[i].Name], &robots[i])
		}
		// pos
		if helper(extent, robots[i].Pos) == false {
			log <- "*  A robot placed outside of the room"
			continue
		}
		//fmt.Println(robots[i].Easting, robots[i].Northing )
		if grid[robots[i].Easting][robots[i].Northing] >= 1 {
			log <- "*  Robots placed at the same place"
		}
		//grid[robots[i].Northing][robots[i].Easting]++
		grid[robots[i].Easting][robots[i].Northing]++
	}


	//
	go func() {
		time.Sleep(1*time.Microsecond)
		close(action)
	}()


	for act := range action{
		//var count int
		if count,ok := tmp[act.Name]; count == 1 && ok {
			list[act.Name].Step2Robot.Control(extent, act.Script, grid, log)
		} else if count > 1 {
			continue
		} else if !ok {
			log <- "*  An action from an unknown robot"
			continue
		}


	}

	// close action

	for i := range robots {
		robots[i].Easting += extent.Min.Easting
		robots[i].Northing += extent.Min.Northing
	}
	report <- robots
}



type Action3 struct{
	Name 	string
	Script 	string
}

func helper(extent Rect, pos Pos) bool {
	if pos.Easting >= 0 &&
		pos.Easting <= extent.Max.Easting &&
		pos.Northing >= 0 &&
		pos.Northing <= extent.Max.Northing {
		return true
	}
	return false
}
