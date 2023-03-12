package section3

// Vertex Q: 構造体の命名をpublicにしてフィールド名をprivateにしたらどうなるのか？
// A: 外部からフィールドにアクセスできない
// Q: この使い方は実装的にあり得るのか？
type Vertex struct {
	X, Y int
	S    string
}

func ChangeVertex(v Vertex) {
	v.X = 1000
}

func ChangeVertex2(v *Vertex) {
	v.X = 1000
}
