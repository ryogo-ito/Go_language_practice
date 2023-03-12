package section4

type Vertex struct {
	x, y int
}

type Vertex3D struct {
	Vertex
	z int
}

// Area  値レシーバー
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

// Scale ポインターレシーバー
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}
