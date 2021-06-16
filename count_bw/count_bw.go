package count_bw

import "fmt"

var N int = 9

//現在の白と黒の数を表示
func CountBW(board [][]string, finish_game bool) (int, [][]int, [][]int, int, int){
    cn_w := 0 //白のオセロの数
    cn_b := 0 //黒のオセロの数
    cn_n := 0 //空白の数
    place_w := [][]int{}
    place_b := [][]int{}
    judge := 0
    //それぞれの個数をカウント
    for i := 0; i < N; i++ {
	    for j := 0; j < N; j++ {
            switch board[i][j]{
                case "●":
                    pw := []int{i,j}
                    place_w = append(place_w, pw)
                    cn_w += 1
                case "○":
                    pb := []int{i,j}
                    place_b = append(place_b, pb)
                    cn_b += 1
                case "-":
                    cn_n += 1
            }
        }
    }
    fmt.Printf("白の数:%d", cn_w)
    fmt.Println()
    fmt.Printf("黒の数:%d", cn_b)
    fmt.Println()
    if finish_game{cn_n = 0}
    //空白が残っているかを判定
    switch cn_n{
        //空白がない場合、白と黒の数から勝者を特定(0:引き分け,1:白,2:黒)
        case 0:
            if cn_w > cn_b{
                judge = 1
            }else if cn_w == cn_b {
                judge = 0
            }else{
                judge = 2
            }
        //空白があっても、白が0個なら黒の勝ち、黒が0個なら白の勝ち
        default:
            if cn_b == 0{
                judge = 1
            }else if cn_w == 0{
                judge = 2
            }else{
                judge = 3
            }
    }
    return judge, place_w, place_b, cn_w, cn_b
}