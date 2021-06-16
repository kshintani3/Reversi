package output_board

import "fmt"

var N int = 9
//現在の盤面の状況を表示
func OutputBoard(board [][]string){
    for i := 0; i < N; i++ {
	    for j := 0; j < N; j++ {
            fmt.Printf(" " + board[i][j])
        }
        fmt.Println()
    }
}