package geo

type LineType int
type AccessType int

const (
	Line1Type LineType = iota
	Line2Type
	Line3Type
	Line4Type

	DirectAccess AccessType = iota
	IndirectAccess
)

type Point1 struct {
	x float64
	y float64
	z float64
}

type Point2 struct {
	xyz [3]float64
}

type Line1 struct {
	pts []Point1
}

type Line2 struct {
	pts []Point2
}

type Line3 struct {
	x []float64
	y []float64
	z []float64
}

type Line4 struct {
	xyz []float64
}

// --------------------------------------------------------------------

type ILine interface {
	IFeature
	Append(x, y, z float64)
}

type IFeature interface {
	Length() int
}

// --------------------------------------------------------------------

func newPoint1(x, y, z float64) Point1 {
	pt := Point1{
		x: x,
		y: y,
		z: z,
	}
	return pt
}

func (pt Point1) Length() int {
	return 1
}

// --------------------------------------------------------------------

func NewPoint2(x, y, z float64) Point2 {
	pt := Point2{}
	pt.xyz[0] = x
	pt.xyz[1] = y
	pt.xyz[2] = z
	return pt
}

func (pt Point2) Length() int {
	return 1
}

// --------------------------------------------------------------------

func NewLine1() *Line1 {
	line := &Line1{
		pts: []Point1{},
	}
	return line
}

func (line *Line1) Length() int {
	return len(line.pts)
}

func (line *Line1) Append(x, y, z float64) {
	line.pts = append(line.pts, newPoint1(x, y, z))
}

// --------------------------------------------------------------------

func NewLine2() *Line2 {
	line := &Line2{
		pts: []Point2{},
	}
	return line
}

func (line *Line2) Length() int {
	return len(line.pts)
}

func (line *Line2) Append(x, y, z float64) {
	line.pts = append(line.pts, NewPoint2(x, y, z))
}

// --------------------------------------------------------------------

func NewLine3() *Line3 {
	line := &Line3{
		x: []float64{},
		y: []float64{},
		z: []float64{},
	}
	return line
}

func (line *Line3) Length() int {
	return len(line.x)
}

func (line *Line3) Append(x, y, z float64) {
	line.x = append(line.x, x)
	line.y = append(line.y, y)
	line.z = append(line.z, z)
}

// --------------------------------------------------------------------

func NewLine4() *Line4 {
	line := &Line4{
		xyz: []float64{},
	}
	return line
}

func (line *Line4) Length() int {
	return len(line.xyz) / 3
}

func (line *Line4) Append(x, y, z float64) {
	line.xyz = append(line.xyz, x)
	line.xyz = append(line.xyz, y)
	line.xyz = append(line.xyz, z)
}
