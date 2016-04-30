package matrix

import (
	"log"
	"math/rand"
	"time"
)

type Matrix struct {
	data      [9][9]int
	validList map[Position][9]int
}

type Position struct {
	x, y int
}

func New() *Matrix {
	x := Matrix{}
	x.validList = make(map[Position][9]int)
	return &x
}

func (m *Matrix) initializeLine(a *[9]int) {
	for i := 0; i < 9; i++ {
		a[i] = i + 1
	}
}

func (m *Matrix) pickValidList(p *Position) [9]int {
	valid := [9]int{}
	m.initializeLine(&valid)

	//find valid values in horizontal
	for _, cur := range m.data[p.x] {
		if cur != 0 {
			valid[cur-1] = 0
		}
	}

	//find valid values in vertical
	for idx, _ := range valid {
		t := m.data[idx][p.y]
		if t != 0 {
			valid[t-1] = 0
		}
	}

	log.Printf("valid values at [%d,%d]:%v\n", p.x, p.y, valid)
	return valid
}

func (m *Matrix) GeneratePlay() {
	p := Position{}
	for {
		log.Printf("===begin:%v\n", p)
		lst, ok := m.validList[p]
		if ok {
			log.Printf("fetch saved valid list at[%d,%d]:%v\n", p.x, p.y, lst)
		} else {
			lst = m.pickValidList(&p)
			m.validList[p] = lst
		}

		rnd := pickRandom(&lst)
		//save the updated list back to map
		m.validList[p] = lst

		if rnd == 0 {
			log.Printf("can not find a valid one, will move back:%v\n", p)
			m.data[p.x][p.y] = 0
			delete(m.validList, p)
			old := p
			moveBack(&p)
			if old == p {
				log.Println("it must already be {0,0}")
				m.PrintMe()
			}
			log.Printf("move one step back:%v\n", p)
			continue
		}
		log.Printf("pick value:%d\n", rnd)
		m.data[p.x][p.y] = rnd
		if p.y < 8 {
			p.y++
		} else {
			p.y = 0
			p.x++
			if p.x > 8 {
				break
			}
		}
		log.Printf("row %d is now: %v\n", p.y, m.data[p.x])
	}
}

//move back a step
func moveBack(p *Position) {
	if p.y > 0 {
		p.y--
	} else {
		if p.x > 0 {
			p.y = 8
			p.x--
		} else {
			//it must be {0,0}
			//we do nothing in this case
		}
	}
}

//pick a non-zero value from an array, and the picked value will be removed from it
func pickRandom(a *[9]int) int {

	lastempty := -1
	for i := 0; i < 9; i++ {
		cur := a[i]
		if cur != 0 {
			if lastempty != -1 {
				a[i], a[lastempty] = 0, a[i]
				lastempty++
			}
		} else {
			if lastempty == -1 {
				lastempty = i
			}
		}
	}
	if lastempty == 0 {
		return 0
	}
	if lastempty == -1 {
		lastempty = 9
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	t := rand.Intn(lastempty)
	var res int
	res, a[t] = a[t], 0

	return res
}

func (m *Matrix) PrintMe() {
	for idx, a := range m.data {
		log.Printf("%d - %v\n", idx, a)
	}
}
