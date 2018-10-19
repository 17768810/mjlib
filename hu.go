package mjlib

const INVALID_CARD = -1

type huMod struct {
}

func (this *huMod) GetHuInfo(cards []int32, cur_card int, gui_1 int, gui_2 int) bool {
	if cur_card != INVALID_CARD {
		cards[cur_card]++
	}

	var (
		gui_num_1 int32
		gui_num_2 int32
	)

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

func check(gui int32, eye_num int32, gui_num int32, gui_sum int32) (bool, int32) {
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

func (this *huMod) split(cards []int32, gui_num int32) bool {
	var (
		eye_num int32
		gui_sum int32
		gui     int32
		ret     bool
	)

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

func (this *huMod) _split(cards []int32, gui_num int32, min int, max int, xu bool, eye_num int32) (int32, int32) {
	var (
		key int32
		b   bool
		num int32
	)

	if xu {
		for i := min; i <= max; i++ {
			num += cards[i]

			if cards[i] > 4 {
				key = 0
				break
			}
			if cards[i] > 0 {
				b = true
				key = key*10 + cards[i]
			} else {
				if b {
					key = key*10 + cards[i]
					b = false
				}
			}
		}
		if !b {
			key = key / 10
		}
	} else {
		for i := min; i <= max; i++ {
			num += cards[i]

			if cards[i] > 4 {
				key = 0
				break
			}
			if cards[i] > 0 {
				key = key*10 + cards[i]
			}
		}
	}

	if num == 0 {
		return 0, eye_num
	}

	for i := int32(0); i <= gui_num; i++ {
		yu := (num + i) % 3
		if yu == 1 {
			continue
		}
		eye := (yu == 2)
		if TableMod.check(key, i, eye, xu) {
			if eye {
				eye_num++
			}
			return i, eye_num
		}
	}

	return -1, 0
}

func (this *huMod) Check7Dui(cards []int32, gui_num int32) bool {
	var need int32
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

func (this *huMod) Check8DuiBan(cards []int32) bool {
	keNum := 0

	for i := 0; i < 34; i++ {
		if cards[i] == 3 {
			keNum++
			cards[i] -= 3
			if !this.Check7Dui(cards, 0) {
				cards[i] += 3
				return false
			}
			cards[i] += 3
		}
	}

	if keNum != 1 {
		return false
	}

	return true
}
