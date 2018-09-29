package mjlib

const INVALID_CARD = -1

type HuLib struct {
}

func (this *HuLib) GetHuInfo(cards []int, cur_card int, gui_1 int, gui_2 int) bool {
	if cur_card != INVALID_CARD {
		cards[cur_card]++
	}

	gui_num_1 := 0
	gui_num_2 := 0
	if gui_1 != INVALID_CARD {
		gui_num_1 = cards[gui_1]
		cards[gui_1] = 0
	}

	if gui_2 != INVALID_CARD {
		gui_num_2 = cards[gui_2]
		cards[gui_2] = 0
	}

	hu := this.split(cards, gui_num_1+gui_num_2)

	if gui_1 != INVALID_CARD {
		cards[gui_1] = gui_num_1
	}

	if gui_2 != INVALID_CARD {
		cards[gui_2] = gui_num_2
	}

	if cur_card != INVALID_CARD {
		cards[cur_card]--
	}

	return hu
}

func check(gui int, eye_num int, gui_num int, gui_sum int) (bool, int) {
	if gui < 0 {
		return false, 0
	}

	gui_sum += gui
	if gui_sum > gui_num {
		return false, 0
	}

	if eye_num == 0 {
		return true, gui_sum
	}

	return gui_sum+(eye_num-1) <= gui_num, gui_sum
}

func (this *HuLib) split(cards []int, gui_num int) bool {
	eye_num := 0
	gui_sum := 0
	gui := 0
	ret := false

	gui, eye_num = this._split(cards, gui_num, 0, 8, true, eye_num)
	ret, gui_sum = check(gui, eye_num, gui_num, gui_sum)
	if ret == false {
		return false
	}

	gui, eye_num = this._split(cards, gui_num-gui_sum, 9, 17, true, eye_num)
	ret, gui_sum = check(gui, eye_num, gui_num, gui_sum)
	if ret == false {
		return false
	}

	gui, eye_num = this._split(cards, gui_num-gui_sum, 18, 26, true, eye_num)
	ret, gui_sum = check(gui, eye_num, gui_num, gui_sum)
	if ret == false {
		return false
	}

	gui, eye_num = this._split(cards, gui_num-gui_sum, 27, 33, false, eye_num)
	ret, gui_sum = check(gui, eye_num, gui_num, gui_sum)
	if ret == false {
		return false
	}

	if eye_num == 0 {
		return gui_sum+2 <= gui_num
	}

	return true
}

// 1→0
// 10→10
// 2→110
// 20→1110
// 3→11110
// 30→111110
// 4→1111110
// 40→11111110
func (this *HuLib) _split(cards []int, gui_num int, min int, max int, xu bool, eye_num int) (int, int) {
	key := 0
	p := -1
	b := false
	num := 0

	if xu {
		for i := min; i <= max; i++ {
			num = num + cards[i]

			if cards[i] == 0 {
				if b {
					b = false
					key |= 0x1 << uint32(p)
					p++
				}
			} else {
				if cards[i] > 4 {
					key = 0
					b = false
					break
				}
				p++
				b = true
				switch cards[i] {
				case 2:
					key |= 0x3 << uint32(p)
					p += 2
				case 3:
					key |= 0xF << uint32(p)
					p += 4
				case 4:
					key |= 0x3F << uint32(p)
					p += 6
				}
			}
		}
		if b {
			key |= 0x1 << uint32(p)
		}
	} else {
		for i := min; i <= max; i++ {
			num = num + cards[i]

			if cards[i] > 0 {
				if cards[i] > 4 {
					key = 0
					break
				}
				p++
				switch cards[i] {
				case 2:
					key |= 0x3 << uint32(p)
					p += 2
				case 3:
					key |= 0xF << uint32(p)
					p += 4
				case 4:
					key |= 0x3F << uint32(p)
					p += 6
				}
				key |= 0x1 << uint32(p)
				p++
			}
		}
	}

	if num == 0 {
		return 0, eye_num
	}

	for i := 0; i <= gui_num; i++ {
		yu := (num + i) % 3
		if yu == 1 {
			continue
		}
		eye := (yu == 2)
		if MTableMgr.check(key, i, eye, xu) {
			if eye {
				eye_num++
			}
			return i, eye_num
		}
	}

	return -1, 0
}

func (this *HuLib) check_7dui(cards []int, gui_num int) bool {
	need := 0
	for i := 0; i < 34; i++ {
		if cards[i]%2 != 0 {
			need = need + 1
		}
	}

	if need > gui_num {
		return false
	}

	return true
}
