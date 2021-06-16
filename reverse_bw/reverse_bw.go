package reverse_bw

//import "fmt"
import "mymodule/direction"

var N int = 9

//置かれたオセロと囲むことができるオセロの座標の探索
func ResearchBW(xi int, yi int, board [][]string, color string) ([][]string, bool){
    var put_possible bool = false
    for i := 0; i < 8; i++{
        var inverse_bw int = 0 //反転するオセロの個数
        var coordinal_dire[2] float32
        xii := xi
        yii := yi
        //fmt.Printf("research=")
        //fmt.Println(search_dire)

        //探索方向の変化率を取得
        coordinal_dire[0],coordinal_dire[1] = direction.Direction(i)

        //置かれたオセロから任意の方向に同じ色のオセロを探す
        for {
            //オセロが置かれた座標に変化率を足す
            xii = int(float32(xii) + coordinal_dire[0])
            yii = int(float32(yii) + coordinal_dire[1])
            //盤面の外までいくと全て0で終了
            if xii == 0 || xii == 9 || yii == 0 || yii == 9{
                inverse_bw = 0
                break
            }else{
                    //同じ色があればその座標とその間にある反対色のオセロの数を返す。
                    if board[xii][yii] == color{
                        break
                    //同じ色が見つかるまでに空白があれば、全て0で終了
                    }else if board[xii][yii] == "-"{
                        inverse_bw = 0
                        break
                    //反対の色があれば、カウントし同じ方向に探索を進める
                    }else{
                        inverse_bw += 1
                    }
            }
        }
        if inverse_bw > 0{
            var ii int = (i + 4) % 8 //searchの逆方向
            var coordinal_redire[2] float32
            //fmt.Printf("reverse=")
            //fmt.Println(search_redire)
            //進む方向の変化率の取得
            coordinal_redire[0],coordinal_redire[1] = direction.Direction(ii)

            for i := 0; i < inverse_bw; i++{
                //挟むことのできるオセロの座標に変化率を足す
                xii = int(float32(xii) + coordinal_redire[0])
                yii = int(float32(yii) + coordinal_redire[1])
                //オセロの色を反転
                board[xii][yii] = color
            }
            put_possible = true
        }
    }
    if put_possible{
        return board, true
    }else{
        return board, false
    }
}

//挟まれたオセロの反転
