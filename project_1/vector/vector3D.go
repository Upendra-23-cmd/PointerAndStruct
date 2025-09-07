package vector


type Vector3D struct {
	X float32
	Y float32
	Z float32
}

func (v1 *Vector3D) AddVector3D(v2 Vector3D) {

	v1.X += v2.X
	v1.Y += v2.Y
	v1.Z += v2.Z

}
func (v1 *Vector3D) SubtractVector3D(v2 Vector3D) {

	v1.X -= v2.X
	v1.Y -= v2.Y
	v1.Z -= v2.Z

}
func (v1 Vector3D) DotVector3D(v2 Vector3D) float32 {
	return (v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z)
}

