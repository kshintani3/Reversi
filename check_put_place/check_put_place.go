package check_put_place

// import "fmt"
import "mymodule/direction"

//置ける場所があるのか探索
func CheckPutPlace(board [][]string, color string, place_othello [][]int, count_othello int) bool{
    var xi int
    var yi int
    var coordinal_dire[2] float32

    for j := 0; j < count_othello; j++{
        for i := 0; i < 8; i++{
            inverse_bw := 0
            xi, yi = place_othello[j][0], place_othello[j][1]
            coordinal_dire[0],coordinal_dire[1] = direction.Direction(i)
            //オセロが置かれた座標に変化率を足す
            for {
                xi = int(float32(xi) + coordinal_dire[0])
                yi = int(float32(yi) + coordinal_dire[1])
                if xi == 0 || xi == 9 || yi == 0 || yi == 9{
                    break
                }else{
                    if board[xi][yi] == color{
                        break
                    }else if board[xi][yi] == "-"{
                        if inverse_bw == 0{
                            break
                        }else{
                            return true
                        }
                    }else{
                        inverse_bw += 1
                    }
                }
            }
        }
    }
    return false
}