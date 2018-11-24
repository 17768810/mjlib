package mjlib

const (
	MJ_TYPE_13 = 13
	MJ_TYPE_16 = 16
)

var (
	TableMod *tableMod
	HuMod    *huMod
	UtilMod  *utilMod
	XuMod    *xuMod
	ZiMod    *ziMod
)

func Init(mjType int32, maxGuiNum int32) {
	TableMod = &tableMod{}
	TableMod.Init()
	HuMod = &huMod{}
	UtilMod = &utilMod{}

	var maxLevel int32

	switch mjType {
	case MJ_TYPE_13:
		maxLevel = 4
	case MJ_TYPE_16:
		maxLevel = 5
	default:
		panic("mjType wrong!")
	}

	XuMod = NewXuMod(maxLevel, maxGuiNum)
	ZiMod = NewZiMod(maxLevel, maxGuiNum)
	// start := time.Now().Unix()
	// println("generate hu tpl begin...")

	XuMod.gen_table()
	ZiMod.gen_zi_table()

	XuMod = nil
	ZiMod = nil

	// println("generate hu tpl end, use time =", time.Now().Unix()-start, "second")
}
