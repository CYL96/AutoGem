package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type MyConfigExt struct {
	AnYing   int `json:"an_ying" type:"int(11)" default:"0" comment:""`
	GuangHui int `json:"guang_hui" type:"int(11)" default:"0" comment:""`
	XingChen int `json:"xing_chen" type:"int(11)" default:"0" comment:""`
}

var MyConfig MyConfigExt
var lastTime int64

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(data, &MyConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("config:", MyConfig)
}

func main() {

	fmt.Println("欢迎使用[自动刷宝石粉]学习软件")
	fmt.Println("本软件由kkyy制造，原汁原味")
	fmt.Println("本软件仅供学习交流请勿用于非法用途，严禁用于商业用途，请于24小时内删除")
	fmt.Println("ctrl+1 : 启动")
	fmt.Println("ctrl+2 : 停止")
	fmt.Println("ctrl+0 : 退出")
	str := ""
	if MyConfig.AnYing == 1 {
		str += "暗影"
	}
	if MyConfig.GuangHui == 1 {
		if len(str) != 0 {
			str += ","
		}
		str += "光辉"
	}
	if MyConfig.XingChen == 1 {
		if len(str) != 0 {
			str += ","
		}
		str += "星辰"
	}
	fmt.Println("当前配置: ", str)

	add()

	// s := robotgo.Start()
	// // end hook
	// defer robotgo.End()
	//
	// for ev := range s {
	// 	fmt.Println("hook: ", ev)
	// }
}
func checkTime() bool {
	now := time.Now().Unix()
	if now-lastTime >= 2 {
		lastTime = now
		return true
	}
	return false
}

const (
	ayStr   = "b40,15,eNqVVUtvG1UUvn7PjO3xZMaPGdsZPybxa+w4fqXyI409de3YVhKDQ1SQioKEgAUrfgJLEEIVVEXisUBlRQFVSAiBVMECqapYsQEVsUFNCqkgAbUKjmtfjmewqWgX5tOVdX3ud893zrnn3tHrx0CPAtjNZrPNZuM4zuv1hsPhRCKRzWZLpZKiKK1Wa2trq9frbW9vd7vddrsNxnK5nMvlZFkGMmxhWRa2WywW/QQGgwF8UhRFkiRMHpbW6XTAAQLsFQQhFAqBaCaTKRaLtVptfX19c3OzpwLUK5VKJBJJp9MQFdBAlOd5mqZBEZyAKy0F0IIw7Ha7w+GAVZhYrVaQIAgCVo1G44NhwEageTweURTBeSqVyufzoA6pQeKFQgGMU7LJZAJvmiuLCtCCv+Af7KAFrhgVbpaGX4eKaQyaLmO3m+20Jg1LQIOCu91uyMU7AcQDWv8hUyoe1JrKzanwOBmfi3bN2WBw7NiirWrBy/F4NBqFKmmJaDWZQksE8DAZisZMoAlNJxCGjSIFF1PMJeVoOOBz806Hh6NhAMZhI7TzyqvnvrnOTmpongBi0AoC89nJMNH6UwoF88vyRrParJXWitlCOiYviqLXZbOSQKP0+tZHV3sYo0uXeYJ4RLcbx+Ul1KUZydBy0BJ1pbbd7bTrq1stpbfReKxTbyplKeAjCQtwOF5Y/Wl/CWMzxsH3PvALvIvjwn7/0kJIzOfnRZFkWO2uzU7O57KtZqPVUJq1yuMdBUZ5JR0SfVaK1MIzIrSyc+45jPVHI9PhMIbxy/t3vr1584ffDl66N7Bd/srPcdNcZid3zlY3mkq7vvbERqNSSGaS/14Es07HMczu7tOf7N3qDjHa76O9/srh/WsYf4nxmQFGNzCn7JiMhnGykOnM5IBf6HXqm421Vq1ULS6nokHoarBDO/tO15596+KVn3/8FOPzcFi/nKDb/dLh4Gp/9PZfI+HPEfoOk8+/xpPGcYS0Y/fC6zOSWYY+u3rqTDlfyadOpWPJSCAdD1OEGZa8HBd45oXqHj6NseUIoztjP9XDwecno3ePh77D++gPrL/y/fy8pFXGpTRnJ0eCvmxCSkeDqUVRlvyQ7HJCcnH/XFgaXmOvhF58E319jPZx/e7oi5PRpeOhB9weDFAfG9+/4WXm/i/ZShDyYjAZCaWi4aWYlI4vZOSFxQAP18fFOgwGvdqfOp4yW2tPxq/f/gzji/0RdzSE4NE9jOC8bmFh/SmTXjf5cMxKlgL+XCqWkSMgCupxSXSzdngx5gWn6HPDU6I5ZHTo/DsfXsCYuztCewP0K0YHGP2OEcbUGx+71HMxqt0yI9lus8YXQpLo83ucgnMOXqe/AVaOrWI="
	ghStr   = "b40,20,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"
	xcStr   = "b40,20,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"
	lockStr = "b100,20,eNrtl/dLI0EUgP9oL1Y8Ti/2grFwGsUuRmNFiUbxFBFjJVFEDIiKJgdRxG5sifeRxy1Diq7LCnfcvp923rxp37wyGwqFwuHwL0t0yMzMTHd39w9LdMjw8HBNTU1hYeFXS94Tt9tdW1tbXFz83ZL3ZGBgoLq62mKlR/r6+iorKy1WeoTEXlZW9neyqq+vn52dpfpUVVWZOzNp52dCKioq9I/q6Oiw2+0Gluvq6vL7/UtLS6YfRJPe3t6np6eHh4eGhgZzZ6aovSakrq5O/6jW1lZjy01MTLDW+fn5h5b7p1k5nU5jAfgfsmpubrZY6RS28amsmpqahoaGyKJTU1NSRzKZ8SpeWFiYm5tzuVxilsqqpKSkp6env7+fbFlaWqoOHxsbm5+fHxkZcTgcqfPz3uZ1tLi4ODo62tjYaIwVtcZuSCYnJ4UVe0trwLNtdXWVw77+kXg8HgqF2tvbVTNeLGtray8vL6+KwI0ujRWnE2OYY/n8/AwZ0UAVROoq9/f33KO6BJTu7u40A4Yzs3xn2nxaAfg3QzI+Pi6s0s6Ar+7s7GDw+Pi4srLCdVP6z87O0FxcXHCbYlZUVLS/v48yFoutr69z6Rxzb2+P09GL8wgrsUfJbAD3er0MlOGAYvjNzY3H4+ns7MQzZUhbW5ssQZIRLMfHx9jgt2xA4/bR438xJIODg6zF8XHO1F6OyfWxbfZss9lECcBIJMKo5eXl7OxsNIRbLCEtLS2aGR/Si5IZotEo4caL6Pr6mrHAycnJEUueKzgMZLDUxgq9ra0t0Wxvb9M8OjpSt3dyciKs2JL+I/PXnJWVZToroo/e3d3d3NxcVS+RSyTyK0ozGAzS3NjY0ECporEiBsPhMJb4Xl5enmYgvn1wcKAO5xWE8vT0FODohQlZTp0ZBzbAKj8/33RW7PDw8FB8IKmLpy9BdHV1havQvLy8lOyUdglhhQgowHKzqoHP50NPBosqgpuhZAk48CwXJnyoA8vLyw2wwp9NZ8WFyulILEldhCQRRzTBCpeTnMw/6RusyFGkGnEVMrl6I/w4CJZgigQCAdKLVu+SUg0FxQArVvyMGJQssbm5mRRcbrdbjUFJX6T9N1jhKlQrwUXaUU9NokZJdbClEwwoXsKEcq/OPD09bYAV8hmsyEvEGuWJG9SUOJJUPbKZZG/yD00iUdBlYoUTksalflFeCwoKxIAshOb29pbeVB+QD9wYG6JV7eWyDLD6DQanrUY="
	bsStr   = "b150,30,eNrtmYtPU1ccx/+SMmjpg2Bky8zcsmRbllkKlJauc2RGtGVLqGZTNsp4BMcWF0oL7W3RDiSDhYrDDVDQ8QwvEY1MkDCgBBV0PCQwmTwcrbTQ/W7vKH3c1mslYYdc8snNub/z+53z+51vz7m3hRUZwsJhOHG1Q9xuSWGQNRhkPoyAUSF+5mJQCGT4SYwRcHaG/1mQXIqkdHGS0o10t6uSrIsK3mMmkBkpj0mEKymEU0+bihGRpUgzyNIM8q8Msv8xcj/tIELkfq6uXjlZm/q8wVUnf5liFZjkmF6SgolfkAQChZPNW/F24D7g1vgKH7tvAgryrLx9FJ7Ogf09LeJthaScFJJySAO3biWaV4EEzV4aROGrQvn5YXwVDaocUL1CK4g2+WHRaiYuIg2aROczBWoWXGlQRc0UaFhwpUEUgZoZo2EJaFAmpiAcRKRBl9jCcFxEGmSJ1bJBxF3D4eK39M2ZHeb6gYc3+h901/dXZFcfSdBHuvtoG9Ov3DF9eUHqsiSVvH3x1g91fT+lVn6IXMlxOg4uIvrEaTlZ1Un350Y2NtbX7M8Wns4vWxahbbWtNgz+nGjc5/LsMNc5HA5dcwZxKy2Kah+5DJ7XxxoPnn0dscILw4U6ThxePvKkVkrmlqbXN+y94x2napNTyvknzouNbbl/LT/acGzU/F4qwniEJ6Eg1pwBbTCaenQ2+9rITP/R0neQqxpEFGJc2IY+sMmMpL3sFwlnUxvTy85+Xj4csSGyd7wddOkea/zYuM/dP6dWvmx5AvsxrSqRsMAZiyvYkinUcfOufvHUujz994PPz4sCTxEwPSoVsamULMRhC6kPqGXHY1zhf4EIc7JSAhoBqRckXl0iLAIOSZDsyoCJsLgUBGfYtov/PIY9i2jhIGK8ngciOuFswvWxUITrSRAOgRPwG4i1ZMBROTx9++DZKN8ptE1KkOzOwx5J0R4wdo7iCppu6Mwz/Za1VUNrlghfhOBqDFxRAAdSO5Wl87DjH1FIXs/FwRCmokcLovTcbYLnmm9vTo0M3m0m5kcPFe93KTi7OAmiw9uL9MxedAsXYhyRYVNBlCnt+h5EuXm/VWSIcLPzCHIvf2Zft43Pmw+V7Ac7oeD1u40r1iU4Qr+5lPzSCRAT7UDhxO+iAufPa07c276wnudAGsKkPH7wI5+qTrav28ceDUr1UZ7Z4hS15IBkfRPX4gt50NU+jD8WsabMsq58eAudWhg/Vh5LuQR/y0J9cbZxEZgH8kLx/xLmhaLOp+c+WFiZs66tZl084tUlLIjom+gGyco61YSlbQhXsOA3JXS1DFbDN8E/Jns/OfMmorXvDgX5Kuavt0pBi5GpPvm597fk0/CK205bbZbZJ5NHi9/1UhDaiUVvDE3dhgdi+3CdWLuHVnAHgU00OjMAWvz5+J6pG/u2JkXflH1ttOGZzbpiWVLVp/LzwnwVBI6Xx4O+cAhX3TTGqjm0gjuIrOS9LvPVFcsi6Ohw/sEr6L3Zoe9qFYJ8lsvNS0Hg9KXj8L0eDmGsMStaxaIV3EFgE8GeKmz4uqxTY2zNzag6/BH2mpdPyo8x2b/I3B98oO9JkxSMJyok7lrTCtLQCtLQCu5u/gVG3Qki"
)

var (
	stop chan int
)

func add() {
	stop = make(chan int, 1)
	on := false
	robotgo.EventHook(hook.KeyDown, []string{"1", "ctrl"}, func(e hook.Event) {
		// JoinWorld(MyConfig.World_1)
		// x, y := robotgo.FindBitmap(gh)
		// fmt.Println(x, y)
		if !on {
			on = true
			go FindCos(stop)
		}

	})
	robotgo.EventHook(hook.KeyDown, []string{"2", "ctrl"}, func(e hook.Event) {
		// JoinWorld(MyConfig.World_2)
		stop <- 1
		on = false

	})
	robotgo.EventHook(hook.KeyDown, []string{"0", "ctrl"}, func(e hook.Event) {
		fmt.Println(time.Now().Format("[2006-01-02 15:04:05]"), "ctrl+shift+0 退出")
		robotgo.EventEnd()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func FindCos(ch chan int) {

	time.Sleep(2 * time.Second)
	startX, startY := 1600, 200
	page := 1
	ay := robotgo.BitmapStr(ayStr)
	defer robotgo.FreeBitmap(ay)
	gh := robotgo.BitmapStr(ghStr)
	defer robotgo.FreeBitmap(gh)
	xc := robotgo.BitmapStr(xcStr)
	defer robotgo.FreeBitmap(xc)
	lock := robotgo.BitmapStr(lockStr)
	defer robotgo.FreeBitmap(lock)
	bs := robotgo.BitmapStr(bsStr)
	defer robotgo.FreeBitmap(bs)
	// 第一次 放到第一页
	fmt.Println("切换第", page, "页")
	x, y := GetPageXY(page)
	time.Sleep(200 * time.Millisecond)
	robotgo.MoveMouseSmooth(x, y, 1.0, 0.1)
	time.Sleep(200 * time.Millisecond)
	robotgo.MouseClick("left")
	time.Sleep(200 * time.Millisecond)
	var last bool
	for {
		select {
		case <-ch:
			return
		default:

			if startX == -1 || startY == -1 {
				if last {
					return
				}
				startX, startY = 1600, 200
				page++
				x, y := GetPageXY(page)
				if x == -1 || y == -1 {
					// 	分解装备
					Fenjie()
					page = 1
					x, y = GetPageXY(page)
					fmt.Println("切换第", page, "页")
					robotgo.MoveMouseSmooth(x, y, 1.0, 0.1)
					time.Sleep(200 * time.Millisecond)
					robotgo.MouseClick("left")
					time.Sleep(200 * time.Millisecond)
					// 	开启宝石
					fmt.Println("开启宝石")
					now := time.Now().Unix()
					for {
						yes := time.Now().Unix()
						if yes-now >= 60 {
							// 	超过一分钟了，说明已经开完了
							last = true
							robotgo.MoveMouseSmooth(960, 480, 1.0, 0.1)
							time.Sleep(200 * time.Millisecond)
							robotgo.MouseClick("left")
							time.Sleep(200 * time.Millisecond)
							break
						}
						robotgo.MoveMouseSmooth(950, 640, 1.0, 0.1)
						time.Sleep(200 * time.Millisecond)
						robotgo.MouseClick("left")
						time.Sleep(200 * time.Millisecond)
						x, y := robotgo.FindBitmap(bs)
						if x != -1 || y != -1 {
							// 关闭窗口
							robotgo.MoveMouseSmooth(960, 480, 1.0, 0.1)
							time.Sleep(200 * time.Millisecond)
							robotgo.MouseClick("left")
							time.Sleep(200 * time.Millisecond)
							break
						}
					}
				} else {
					fmt.Println("切换第", page, "页")
					robotgo.MoveMouseSmooth(x, y, 1.0, 0.1)
					time.Sleep(200 * time.Millisecond)
					robotgo.MouseClick("left")
					time.Sleep(200 * time.Millisecond)
				}
			}

			robotgo.MoveMouseSmooth(startX, startY, 1.0, 0.1)
			time.Sleep(200 * time.Millisecond)
			robotgo.MouseClick("right")
			time.Sleep(200 * time.Millisecond)

			// 识别宝石
			var x, y int
			x, y = func() (int, int) {
				var xtmp, ytmp = -1, -1
				// ---------------------
				if MyConfig.AnYing == 1 {
					xtmp, ytmp = robotgo.FindBitmap(ay)
				}
				if xtmp != -1 || ytmp != -1 {
					fmt.Println("找到暗影宝石,开始强化")
					return xtmp, ytmp
				}
				// ---------------------
				if MyConfig.GuangHui == 1 {
					xtmp, ytmp = robotgo.FindBitmap(gh)
				}
				if xtmp != -1 || ytmp != -1 {
					fmt.Println("找到光辉宝石,开始强化")
					return xtmp, ytmp
				}
				// ---------------------
				if MyConfig.XingChen == 1 {
					xtmp, ytmp = robotgo.FindBitmap(xc)
				}
				if xtmp != -1 || ytmp != -1 {
					fmt.Println("找到星辰宝石,开始强化")
					return xtmp, ytmp
				}
				return xtmp, ytmp
			}()
			if x != -1 || y != -1 {
				fmt.Println("找到宝石,开始强化")
				now := time.Now().Unix()

				for {
					robotgo.KeyTap("space")
					time.Sleep(100 * time.Millisecond)
					x, y := robotgo.FindBitmap(lock)
					if x != -1 || y != -1 || time.Now().Unix()-now > 10 {
						// 10秒内就能升级完毕
						break
					}
				}
			}
			startX, startY = GetNextPos(startX, startY)
		}
	}
}

// 1590,200                1840 200
//
// 1590,800                1840,800

func GetNextPos(x, y int) (x_1, y_2 int) {
	if x < 1850 {
		return x + 67, y
	} else if y < 780 {
		return 1600, y + 65
	} else {
		return -1, -1
	}
}

func GetPageXY(page int) (x_1, y_2 int) {
	if page > 3 {
		return -1, -1
	}
	switch page {
	case 1:
		return 1650, 860
	case 2:
		return 1730, 860
	case 3:
		return 1800, 860

	}
	return
}
func Fenjie() {
	robotgo.KeyTap("e")
	fmt.Println("分解宝石")
	// 点击分解
	time.Sleep(200 * time.Millisecond)
	robotgo.MoveMouseSmooth(1130, 770, 1.0, 0.1)
	time.Sleep(200 * time.Millisecond)
	robotgo.MouseClick("left")
	time.Sleep(500 * time.Millisecond)

	// 点击确认
	robotgo.MoveMouseSmooth(850, 480, 1.0, 0.1)
	time.Sleep(200 * time.Millisecond)
	robotgo.MouseClick("left")
	time.Sleep(500 * time.Millisecond)

	// 点击x
	robotgo.MoveMouseSmooth(1320, 240, 1.0, 0.1)
	time.Sleep(200 * time.Millisecond)
	robotgo.MouseClick("left")
	time.Sleep(200 * time.Millisecond)

}
