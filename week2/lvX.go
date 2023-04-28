package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var t, s string
var ch = make(chan int, 1)
var (
	myHp    = 100
	enemyHp = 100
)

func main() {
	hurt := map[int]int{
		1: 10,
		2: 14,
		3: 16,
		4: 20,
		5: 25,
	}
	m := map[string]string{
		"1":  "龙卷风摧毁停车场",
		"2":  "黑虎掏心",
		"3":  "乌鸦坐飞机",
		"4":  "暴龙振翅飞翔",
		"5":  "笨驴踢腿",
		"6":  "仓鼠上车轮",
		"7":  "超行星燃烧",
		"8":  "大象踢腿",
		"9":  "二龙戏珠",
		"10": "飞鹤捕虾",
		"11": "飞龙在天",
		"12": "飞天猴巧夺宝盒",
		"13": "飞鹰展翅",
		"14": "弗拉明戈舞步",
		"15": "黑虎捕食困小羊",
		"16": "黑虎掠过秃鹰",
		"17": "火山烧农场",
		"18": "脑袋砸核桃",
	}
	n := map[string]string{
		"1": "尝尝我的厉害吧",
		"2": "哒咩",
		"3": "无胆懦夫",
		"4": "抱头鼠窜吧",
		"5": "感受神的愤怒",
		"6": "你的毁灭将让我喜不自禁",
		"7": "汝之手段，只与田间老农无异",
		"8": "你做得还不够好",
	}
	for i := 1; ; i++ {
		fmt.Println("Round", i)
		templateChoice(n)
		skillChoice(m)
		ReleaseSkill(m[s], func(skillName string) {
			fmt.Println("hero:"+n[t], skillName)
		})
		hp(hurt)
		if enemyHp < 0 {
			fmt.Println("我是无敌的")
			break
		}
		x1 := rand.Intn(7)
		s1 := strconv.Itoa(x1 + 1)
		x2 := rand.Intn(7)
		t1 := strconv.Itoa(x2 + 1)
		ReleaseSkill(m[s1], func(skillName string) {
			fmt.Println("enemy:"+n[t1], skillName)
		})
		if myHp < 0 {
			fmt.Println("你倒在了前进的路上")
			break
		}
		fmt.Println("hero's Hp:", myHp, "enemy's Hp:", enemyHp)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		time.Sleep(time.Second)
	}
}
func templateChoice(n map[string]string) {
	println("请选择释放技能时输出文字的模板")

	var b string
	for i := 1; i < 9; i++ {
		b = strconv.Itoa(i)
		fmt.Println(b + ":" + n[b])
	}

	fmt.Scanln(&t)

}
func skillChoice(m map[string]string) {
	println("请选择释放的技能")
	for i := 1; i < 19; i++ {
		a := strconv.Itoa(i)
		fmt.Println(a + ":" + m[a])
	}
	fmt.Scanln(&s)

}
func hp(hurt map[int]int) {
	rand.Seed(time.Now().UnixNano())
	Hurt := rand.Intn(4)
	myHp -= hurt[Hurt+1]
	enemyHp -= (hurt[Hurt+1] + 20) / 2

}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
