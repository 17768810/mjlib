package mjlib

type ziMod struct {
	maxLevel          int32
	maxGuiNum         int32
	zi_gui_tested     [9]*map[int32]bool
	zi_gui_eye_tested [9]*map[int32]bool
}

func NewZiMod(maxLevel int32, maxGuiNum int32) *ziMod {
	return &ziMod{
		maxLevel:          maxLevel,
		maxGuiNum:         maxGuiNum,
		zi_gui_tested:     [9]*map[int32]bool{},
		zi_gui_eye_tested: [9]*map[int32]bool{},
	}
}

func (this *ziMod) zi_check_add(cards []int32, gui_num int32, eye bool) bool {
	var key int32

	for i := 0; i < 7; i++ {
		if cards[i] > 3 {
			key = 0
			break
		}
		if cards[i] > 0 {
			key = key*10 + cards[i]
		}
	}

	var m *map[int32]bool
	if !eye {
		m = this.zi_gui_tested[gui_num]
	} else {
		m = this.zi_gui_eye_tested[gui_num]
	}
	_, ok := (*m)[key]

	if ok {
		return false
	}

	(*m)[key] = true

	if key == 0 {
		return true
	}

	TableMod.Add(key, gui_num, eye, false)

	return true
}

func (this *ziMod) parse_zi_table_sub(cards []int32, num int32, eye bool) {
	for i := 0; i < 7; i++ {
		if cards[i] == 0 {
			continue
		}

		cards[i]--

		if !this.zi_check_add(cards, num, eye) {
			cards[i]++
			continue
		}

		if num < this.maxGuiNum {
			this.parse_zi_table_sub(cards, num+1, eye)
		}
		cards[i]++
	}
}

func (this *ziMod) parse_zi_table(cards []int32, eye bool) {
	if !this.zi_check_add(cards, 0, eye) {
		return
	}
	this.parse_zi_table_sub(cards, 1, eye)
}

func (this *ziMod) gen_3(cards []int32, level int32, eye bool) {
	for i := 0; i < 7; i++ {
		if cards[i] > 3 {
			continue
		}
		cards[i] += 3

		this.parse_zi_table(cards, eye)
		if level < this.maxLevel {
			this.gen_3(cards, level+1, eye)
		}

		cards[i] -= 3
	}
}

func (this *ziMod) gen_zi_table() {
	for i := 0; i < 9; i++ {
		this.zi_gui_tested[i] = &map[int32]bool{}
		this.zi_gui_eye_tested[i] = &map[int32]bool{}
	}

	cards := []int32{
		0, 0, 0, 0, 0, 0, 0,
	}

	// 无眼
	// fmt.Printf("无眼表生成开始\n")
	this.gen_3(cards, 1, false)
	// fmt.Printf("无眼表生成结束\n")

	// 有眼
	// fmt.Printf("有眼表生成开始\n")
	for i := 0; i < 7; i++ {
		cards[i] = 2
		// fmt.Printf("将 %d \n", i)
		this.parse_zi_table(cards, true)
		this.gen_3(cards, 1, true)
		cards[i] = 0
	}
	// fmt.Printf("有眼表生成结束\n")

	this.zi_gui_tested = [9]*map[int32]bool{}
	this.zi_gui_eye_tested = [9]*map[int32]bool{}

	// fmt.Printf("表数据存储开始\n")
	// TableMod.DumpziTable()
	// fmt.Printf("表数据存储结束\n")
}
