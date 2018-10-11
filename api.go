package mjlib

const (
	MJ_TYPE_13 = 13
	MJ_TYPE_16 = 16
)

var (
	TableMod *tableMod
	HuMod    *huMod
	UtilMod  *utilMod
)

func Init(mjType int) {
	TableMod = &tableMod{}
	TableMod.Init()
	HuMod = &huMod{}
	UtilMod = &utilMod{}

	var maxLevel int

	switch mjType {
	case MJ_TYPE_13:
		maxLevel = 4
	case MJ_TYPE_16:
		maxLevel = 5
	default:
		panic("mjType wrong!")
	}

	// start := time.Now().Unix()
	// println("generate hu tpl begin...")

	gen_table(maxLevel)
	gen_zi_table(maxLevel)

	// println("generate hu tpl end, use time =", time.Now().Unix()-start, "second")
}
