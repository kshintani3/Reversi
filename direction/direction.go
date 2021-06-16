package direction

//探索方向の変化率の数値化
func Direction(search_dire int) (float32,float32){
    switch search_dire{
        case 0: return -1,0 //上
        case 1: return -1,1 //右上
        case 2: return 0,1 //右
        case 3: return 1,1 //右下
        case 4: return 1,0 //下
        case 5: return 1,-1 //左下
        case 6: return 0,-1 //左
        case 7: return -1,-1 //左上
        default: return 0,0
    }
}