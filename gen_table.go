package mjlib

type xuMod struct {
	maxLevel       int32
	maxGuiNum      int32
	gui_tested     [9]*map[int32]bool
	gui_eye_tested [9]*map[int32]bool
}

func NewXuMod(maxLevel int32, maxGuiNum int32) *xuMod {
	return &xuMod{
		maxLevel:       maxLevel,
		maxGuiNum:      maxGuiNum,
		gui_tested:     [9]*map[int32]bool{},
		gui_eye_tested: [9]*map[int32]bool{},
	}
}

func (this *xuMod) check_add(cards []int32, gui_num int32, eye bool) bool {
	var (
		key int32
		b   bool
	)

	for i := 0; i < 9; i++ {
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

	var m *map[int32]bool
	if !eye {
		m = this.gui_tested[gui_num]
	} else {
		m = this.gui_eye_tested[gui_num]
	}
	_, ok := (*m)[key]

	if ok {
		return false
	}

	(*m)[key] = true

	if key == 0 {
		return true
	}

	TableMod.Add(key, gui_num, eye, true)

	return true
}

func (this *xuMod) parse_table_sub(cards []int32, num int32, eye bool) {
	for i := 0; i < 9; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--

		if !this.check_add(cards, num, eye) {
			cards[i]++
			continue
		}

		if num < this.maxGuiNum {
			this.parse_table_sub(cards, num+1, eye)
		}
		cards[i]++
	}
}

func (this *xuMod) parse_table(cards []int32, eye bool) {
	if !this.check_add(cards, 0, eye) {
		return
	}
	this.parse_table_sub(cards, 1, eye)
}

func (this *xuMod) gen_111_3(cards []int32, level int32, eye bool) {
	for i := 0; i < 16; i++ {
		if i <= 8 {
			if cards[i] > 3 {
				continue
			}
			cards[i] += 3
		} else {
			index := i - 9
			if cards[index] > 5 || cards[index+1] > 5 || cards[index+2] > 5 {
				continue
			}
			cards[index]++
			cards[index+1]++
			cards[index+2]++
		}

		this.parse_table(cards, eye)
		if level < this.maxLevel {
			this.gen_111_3(cards, level+1, eye)
		}

		if i <= 8 {
			cards[i] -= 3
		} else {
			index := i - 9
			cards[index]--
			cards[index+1]--
			cards[index+2]--
		}
	}
}

func (this *xuMod) gen_table() {
	for i := 0; i < 9; i++ {
		this.gui_tested[i] = &map[int32]bool{}
		this.gui_eye_tested[i] = &map[int32]bool{}
	}

	cards := []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	// 无眼
	// fmt.Printf("无眼表生成开始\n")
	this.gen_111_3(cards, 1, false)
	// fmt.Printf("无眼表生成结束\n")

	// 有眼
	// fmt.Printf("有眼表生成开始\n")
	for i := 0; i < 9; i++ {
		cards[i] = 2
		// fmt.Printf("将 %d \n", i)
		this.parse_table(cards, true)
		this.gen_111_3(cards, 1, true)
		cards[i] = 0
	}
	// fmt.Printf("有眼表生成结束\n")

	this.gui_tested = [9]*map[int32]bool{}
	this.gui_eye_tested = [9]*map[int32]bool{}

	// fmt.Printf("表数据存储开始\n")
	// TableMod.DumpTable()
	// fmt.Printf("表数据存储结束\n")
}
