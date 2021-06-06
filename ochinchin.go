package main

import (
	"fmt"
	"math"
	"strings"
)

// サイコロの情報
var dice = []string{"お", "ま", "ち", "こ", "う", "ん"}

// 使用するサイコロの数
var diceCnt = 5

func main() {
	// 使用するサイコロをset
	diceList := setDice(diceCnt)

	// おちんちんパターン数取得
	op := ochinchinPatternCheck(diceList)
	fmt.Printf("op is %d\n", op)

	// サイコロの目の出方のパターン数取得
	dp := dicePattern(len(diceList))
	fmt.Printf("dp is %d\n", dp)

	// "お", "ち", "ん", "ち", "ん"となる確率を求める
	result := (float64(op) / float64(dp)) * 100
	fmt.Printf("サイコロを振って「おちんちん」になる確率は、%f％です\n", result)
}

// 使用するサイコロを配列に定義する
func setDice(num int) [][]string {
	list := [][]string{}
	for i := 0; i < num; i++ {
		list = append(list, dice)
	}
	return list
}

// サイコロを振った時 "お", "ち", "ん", "ち", "ん" になるのが何パターンあるかチェック
func ochinchinPatternCheck(dl [][]string) int {
	cnt := 0
	for i0 := 0; i0 < len(dl[0]); i0++ {
		one := dl[0][i0]
		for i1 := 0; i1 < len(dl[1]); i1++ {
			two := dl[1][i1]
			for i2 := 0; i2 < len(dl[2]); i2++ {
				three := dl[2][i2]
				for i3 := 0; i3 < len(dl[3]); i3++ {
					four := dl[3][i3]
					for i4 := 0; i4 < len(dl[4]); i4++ {
						five := dl[4][i4]
						succesStr := "おちんちん"
						s := fmt.Sprintf("%s%s%s%s%s", one, two, three, four, five)
						// 生成された文字列が"お", "ち", "ん", "ち", "ん"となるか
						if ochinchinCheck(s, &succesStr) {
							cnt++
						}
					}
				}
			}
		}
	}
	return cnt
}

// 渡ってきた文字列を分割し、「おちんちん」になるかチェックする
func ochinchinCheck(s string, pointerStr *string) bool {
	slice := strings.Split(s, "")
	for _, v := range slice {
		if strings.Contains(*pointerStr, v) {
			*pointerStr = strings.Replace(*pointerStr, v, "", 1)
		}
	}
	return len(*pointerStr) == 0
}

// サイコロの目の出方は何通りあるか
func dicePattern(diceNum int) int {
	// サイコロの要素数はいくつか
	dicePattern := len(dice)
	// べき乗：サイコロの要素数^サイコロの数
	return int(math.Pow(float64(dicePattern), float64(diceNum)))
}
