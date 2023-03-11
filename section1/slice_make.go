package section1

//func main() {
//n := make([]int, 3, 5)
//fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
//
//n = append(n, 0, 0)
//fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
//
//n = append(n, 1, 2, 3, 4, 5)
//fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

//// メモリ領域を確保する
//b := make([]int, 0)
//// 値がnillになるのでメモリ領域を確保しない
//var c []int
//
//fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
//fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

// for文の実行結果
// [0 0 0 0 0 0]
// [0 0 0 0 0 0 1]
// [0 0 0 0 0 0 1 2]
// [0 0 0 0 0 0 1 2 3]
// [0 0 0 0 0 0 1 2 3 4]
// len,capも5の為make時点でlength 5つ分の0が入る
//c = make([]int, 5)

// for文の実行結果
// [0]
// [0 1]
// [0 1 2]
// [0 1 2 3]
// [0 1 2 3 4]
// len,caoがそれぞれ指定されている為make時点でlengthは0,capは5になっている為上記の結果になる
//c = make([]int, 0, 5)

//for i := 0; i < 5; i++ {
//	c = append(c, i)
//	fmt.Println(c)
//}
//
//fmt.Println(c)
//}
