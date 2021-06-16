package main

import "fmt"
import "strconv"

import "mymodule/output_board"
import "mymodule/count_bw"
import "mymodule/check_put_place"
import "mymodule/reverse_bw"

var N int = 9 //盤面[8*8]とそれぞれの座標を定義

func main() {
    turn_play := 1
	color := "○" //現在の入力の色
	color_name := "White" //色の名前
	xi := 0 //入力されたオセロのX座標
	yi := 0 //入力されたオセロのY座標
	put_possible := true //入力した場所が置くことができるのか
	exist_put_place := false //置く場所があるのかというフラグ
	judge_winner := 3 //白と黒の数から勝者を特定(0:引き分け,1:白,2:黒)、続くなら3
	count_white := 0 //白の数
	count_black := 0 //黒の数
	place_white := [][]int{} //白のオセロの座標
    place_black := [][]int{} //黒のオセロの座標
    count_othello := 0 //判定する色のオセロの数
    place_othello := [][]int{} //判定する色のオセロの座標
    judge_game := 0 //どちらも置けない時にゲームを終了するフラグ
    finish_game := false //どちらも置く場所がない時に、その時点のそれぞれの数で勝敗を決める。

    //盤面の初期化（スタート時に白２個、黒２個を配置）
    board := make([][]string, N)
    for i := 0; i < N; i++ {
        board[i] = make([]string, N)
        for j := 0; j < N; j++ {
            switch i {
                case 0:
                    board[i][j] = strconv.Itoa(j)//string型に変換
                default:
                    switch j{
                        case 0:
                            board[i][j] = strconv.Itoa(i)
                        default:
                            board[i][j] = "-"
                    }
            }
        }
    }
    board[0][0] = "*"
    board[4][4] = "●"
    board[4][5] = "○"
    board[5][4] = "○"
    board[5][5] = "●"

    for{
        //盤面を出力する
        output_board.OutputBoard(board)
        //現在の白、黒の座標と数
        judge_winner, place_white, place_black, count_white, count_black = count_bw.CountBW(board, finish_game)
        if judge_winner == 0 {
            fmt.Println("引き分けです")
                break
        }else if judge_winner == 1 || judge_winner == 2 {
            var winner_color string
            if judge_winner == 1{
                winner_color = "白"
            }else{
                winner_color = "黒"
            }
            fmt.Printf("勝者は%sです\n", winner_color)
            break
        }
        switch turn_play % 2{
            case 0:
                color = "●"
                color_name = "White"
                place_othello, _ = place_white, place_black
                count_othello, _ = count_white, count_black
            default:
                color = "○"
                color_name = "Black"
                _, place_othello = place_white, place_black
                _, count_othello = count_white, count_black
        }
        exist_put_place = check_put_place.CheckPutPlace(board, color, place_othello, count_othello)

        if exist_put_place{
            judge_game = 0
            fmt.Printf("ターン%d:%s", turn_play, color_name)
            fmt.Println()
            fmt.Printf("オセロを置く場所を選んでください(縦 横)例.5 6: ")

            input_bw := make([]int, 2)
            fmt.Scanln(&input_bw[0], &input_bw[1])
            put_possible = false
            //[8,8]のボードより外に置こうとした時
            if input_bw[0] > 8 || input_bw[1] > 8 {
                fmt.Println("*******正しくありません*******")
            //すでにオセロが置かれている時
            }else if board[input_bw[0]][input_bw[1]] != "-"{
                fmt.Println("***すでにオセロが置かれてます***")
            }else{
                xi = input_bw[0]
                yi = input_bw[1]
                board, put_possible = reverse_bw.ResearchBW(xi, yi, board, color)

                if put_possible {
                    board[xi][yi] = color
                    turn_play += 1
                }else{
                    fmt.Println("***選んだ場所には置けません***")
                }
            }
        }else{
            fmt.Printf("****%sは置ける場所がありません****", color)
            fmt.Println()
            judge_game += 1
            turn_play += 1
            if judge_game == 2{
                finish_game = true
            }
        }
	}
}