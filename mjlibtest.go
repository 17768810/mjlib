package mjlib

import (
	"fmt"
	"time"
)

var g_cards = []int32{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, // 万
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, // 条
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // 筒
	0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, // 东南西北中发白
}

func value2index(value int32) int32 {
	if value < 0x31 {
		return ((value&0xF0)>>4)*9 + (value & 0x0F) - 1
	} else {
		return 27 + ((value & 0xF0) >> 4) - 3
	}
}

func count(cards []int32) []int32 {
	nums := make([]int32, 34)
	for _, v := range cards {
		nums[value2index(v)]++
	}
	return nums
}

func getPairs() [][]int32 {
	pairs := make([][]int32, 0, len(g_cards))

	for _, v := range g_cards {
		pair := []int32{v, v}
		pairs = append(pairs, pair)
	}

	return pairs
}

func getGroups() [][]int32 {
	groups := make([][]int32, 0, len(g_cards)+(9-2)*3)

	// find three identical tiles
	for _, v := range g_cards {
		group := []int32{v, v, v}
		groups = append(groups, group)
	}

	// find three sequence tiles
	for i := 2; i < len(g_cards); i++ {
		if g_cards[i-2]+1 == g_cards[i-1] && g_cards[i-1] == g_cards[i]-1 {
			group := []int32{g_cards[i-2], g_cards[i-1], g_cards[i]}
			groups = append(groups, group)
		}
	}

	return groups
}

func TestAll() {
	pairs := getPairs()
	groups := getGroups()
	println("len(pairs)", len(pairs))
	println("len(groups)", len(groups))

	start := time.Now().Unix()

	for _, p := range pairs {

		encode(p)

		for _, a := range groups {
			var a_temp []int32
			a_temp = append(a_temp, p...)
			a_temp = append(a_temp, a...)
			encode(a_temp)

			for _, b := range groups {
				var b_temp []int32
				b_temp = append(b_temp, a_temp...)
				b_temp = append(b_temp, b...)
				encode(b_temp)

				for _, c := range groups {
					var c_temp []int32
					c_temp = append(c_temp, b_temp...)
					c_temp = append(c_temp, c...)
					encode(c_temp)

					for _, d := range groups {
						var d_temp []int32
						d_temp = append(d_temp, c_temp...)
						d_temp = append(d_temp, d...)
						encode(d_temp)

						for _, e := range groups {
							var e_temp []int32
							e_temp = append(e_temp, d_temp...)
							e_temp = append(e_temp, e...)
							encode(e_temp)
						}

					}
				}
			}
		}
	}

	fmt.Println("use time=", time.Now().Unix()-start)
}

func encode(cards []int32) {
	nums := count(cards)

	if checkIsValid(nums) {
		if !MHuLib.GetHuInfo(nums, -1, -1, -1) {
			print_cards(nums)
			panic("hu err!")
		} else {
			// print_cards(nums)
		}
	} else {
		if MHuLib.GetHuInfo(nums, -1, -1, -1) {
			print_cards(nums)
			panic("no hu err!")
		} else {
			// print_cards(nums)
		}
	}
}

func checkIsValid(nums []int32) bool {

	var allNum int32

	for _, v := range nums {
		allNum += v
		if v > 4 {
			return false
		}
	}

	if allNum%3 != 2 {
		return false
	}

	return true
}

func print_cards(cards []int32) {
	for i := 0; i < 9; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 9; i < 18; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 18; i < 27; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")

	for i := 27; i < 34; i++ {
		fmt.Printf("%d,", cards[i])
	}
	fmt.Printf("\n")
}

var tested = map[int32]bool{}

func check_hu(cards []int32, max int) {
	for i := 0; i < max; i++ {
		if cards[i] > 4 {
			return
		}
	}

	var num int32
	for i := 0; i < 9; i++ {
		num = num*10 + cards[i]
	}

	_, ok := tested[num]
	if ok {
		return
	}

	tested[num] = true

	for i := 0; i < max; i++ {
		if !MHuLib.GetHuInfo(cards, -1, -1, -1) {
			fmt.Printf("测试失败 i=%d\n", i)
			print_cards(cards)
		}
	}
}

func gen_auto_table_sub(cards []int32, level int) {
	for i := 0; i < 32; i++ {
		index := -1
		if i <= 17 {
			cards[i] += 3
		} else if i <= 24 {
			index = i - 18
		} else {
			index = i - 16
		}

		if index >= 0 {
			cards[index] += 1
			cards[index+1] += 1
			cards[index+2] += 1
		}

		if level == 4 {
			check_hu(cards, 18)
		} else {
			gen_auto_table_sub(cards, level+1)
		}

		if i <= 17 {
			cards[i] -= 3
		} else {
			cards[index] -= 1
			cards[index+1] -= 1
			cards[index+2] -= 1
		}
	}
}

func test_two_color() {
	fmt.Println("测试两种花色")
	cards := []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	for i := 0; i < 18; i++ {
		cards[i] = 2
		fmt.Printf("将 %d\n", i+1)
		gen_auto_table_sub(cards, 1)
		cards[i] = 0
	}
}

func test_one_success() {
	// cards := []int{
	// 	0, 0, 0, 0, 0, 1, 0, 0, 0,
	// 	1, 1, 0, 0, 0, 0, 1, 0, 0,
	// 	0, 0, 1, 0, 0, 0, 0, 0, 0,
	// 	1, 0, 0, 0, 0, 4, 4,
	// }

	// [402 402 205 205 205 403 403 403 202 203 204 107 108 109 205 206 207]
	// cards := []int{
	// 	0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 	0, 0, 0, 0, 0, 0, 0, 0, 0,
	// 	0, 3, 3, 0, 3, 4, 3, 1, 0,
	// 	0, 0, 0, 0, 0, 0, 0,
	// }

	cards := []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 3, 0, 3, 0, 3, 0, 3, 3,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	fmt.Println("测试1种能胡的牌型")
	print_cards(cards)
	if MHuLib.GetHuInfo(cards, -1, -1, -1) {
		fmt.Println("测试通过：胡牌")
	} else {
		fmt.Println("测试失败：能胡的牌型判断为不能胡牌")
	}
}

func test_one_fail() {
	cards := []int32{
		0, 1, 1, 1, 0, 0, 1, 0, 1,
		0, 1, 1, 1, 0, 0, 2, 2, 2,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
	}

	fmt.Println("测试1种不能胡的牌型")
	print_cards(cards)
	if !MHuLib.GetHuInfo(cards, -1, -1, -1) {
		fmt.Println("测试通过：不能胡牌")
	} else {
		fmt.Println("测试失败：不能胡牌的牌型判断为胡了")
	}
}

func test_time(count int) {
	cards := []int32{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 3, 4, 1, 1, 0,
		2, 3, 3, 0, 0, 0, 0,
	}

	print_cards(cards)
	start := time.Now().Unix()
	for i := 0; i < count; i++ {
		MHuLib.GetHuInfo(cards, -1, -1, -1)
	}
	print_cards(cards)
	fmt.Println("count=", count, "use time=", time.Now().Unix()-start)
}

func Test() {
	fmt.Println("test hulib begin...")

	// Init()
	// MTableMgr.LoadTable()
	// MTableMgr.LoadziTable()

	test_one_success()
	// test_one_fail()
	// test_time(100000000)
	//    test_two_color()
}
