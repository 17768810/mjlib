package mjlib

import (
	"fmt"
)

type TableMgr struct {
	m_tbl        [9]*Table
	m_eye_tbl    [9]*Table
	m_zi_tbl     [9]*Table
	m_zi_eye_tbl [9]*Table
}

func (this *TableMgr) Init() {
	for i := 0; i < 9; i++ {
		this.m_tbl[i] = &Table{}
		this.m_tbl[i].init()
	}

	for i := 0; i < 9; i++ {
		this.m_eye_tbl[i] = &Table{}
		this.m_eye_tbl[i].init()
	}

	for i := 0; i < 9; i++ {
		this.m_zi_tbl[i] = &Table{}
		this.m_zi_tbl[i].init()
	}

	for i := 0; i < 9; i++ {
		this.m_zi_eye_tbl[i] = &Table{}
		this.m_zi_eye_tbl[i].init()
	}
}

func (this *TableMgr) getTable(gui_num int, eye bool, xu bool) *Table {
	var tbl *Table
	if xu {
		if eye {
			tbl = this.m_eye_tbl[gui_num]
		} else {
			tbl = this.m_tbl[gui_num]
		}
	} else {
		if eye {
			tbl = this.m_zi_eye_tbl[gui_num]
		} else {
			tbl = this.m_zi_tbl[gui_num]
		}
	}
	return tbl
}

func (this *TableMgr) Add(key int, gui_num int, eye bool, xu bool) {
	tbl := this.getTable(gui_num, eye, xu)
	tbl.add(key)
}

func (this *TableMgr) check(key int, gui_num int, eye bool, xu bool) bool {
	tbl := this.getTable(gui_num, eye, xu)
	return tbl.check(key)
}

func (this *TableMgr) LoadTable() {
	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/table_%d.tbl", i)
		this.m_tbl[i].load(name)
	}

	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/eye_table_%d.tbl", i)
		this.m_eye_tbl[i].load(name)
	}

}

func (this *TableMgr) DumpTable() {
	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/table_%d.tbl", i)
		this.m_tbl[i].dump(name)
	}

	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/eye_table_%d.tbl", i)
		this.m_eye_tbl[i].dump(name)
	}

}

func (this *TableMgr) LoadziTable() {
	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/zi_table_%d.tbl", i)
		this.m_zi_tbl[i].load(name)
	}

	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/zi_eye_table_%d.tbl", i)
		this.m_zi_eye_tbl[i].load(name)
	}
}

func (this *TableMgr) DumpziTable() {
	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/zi_table_%d.tbl", i)
		this.m_zi_tbl[i].dump(name)
	}

	for i := 0; i < 9; i++ {
		name := fmt.Sprintf("tbl/zi_eye_table_%d.tbl", i)
		this.m_zi_eye_tbl[i].dump(name)
	}
}
