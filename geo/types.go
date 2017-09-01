package geo

//
// We define two kinds of points:
//   Point1:  x,y,z float
//   Point2:  xyz [3]float
//
// We then define four kinds of lines:
//   Line1:  pts []Point1
//   Line2:  pts []Point2
//   Line3:  x,y,z []float64
//   Line4:  xyz []float64
//
// We next define the Line interface with three methods:
//   AddLast:  add a point to end of line
//   RemoveFirst:  remove first point from line
//   Reproject:  modify each point in the line
//
// And then, using Go's benchmark support, we test the speed of each of the
// four kinds of line on each of the three methods.
//

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

type Line interface {
	Length() int

	// replace the current list with a list of N points of (0,0,0)
	Resize(n int)

	// add point to end of list
	AddLast(x, y, z float64)

	// remove first point from list
	RemoveFirst()

	// add (1,1,1) to every point in list
	Reproject()
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

// --------------------------------------------------------------------

func NewPoint2(x, y, z float64) Point2 {
	pt := Point2{}
	pt.xyz[0] = x
	pt.xyz[1] = y
	pt.xyz[2] = z
	return pt
}

// --------------------------------------------------------------------

func NewLine1() *Line1 {
	line := &Line1{
		pts: []Point1{},
	}
	return line
}

func (line *Line1) Resize(n int) {
	line.pts = make([]Point1, n)
}

func (line *Line1) Length() int {
	return len(line.pts)
}

func (line *Line1) AddLast(x, y, z float64) {
	line.pts = append(line.pts, newPoint1(x, y, z))
}

func (line *Line1) RemoveFirst() {
	line.pts = line.pts[1:len(line.pts)]
}

func (line *Line1) Reproject() {
	n := line.Length()
	for i := 0; i < n; i++ {
		line.pts[i].x += 1
		line.pts[i].y += 1
		line.pts[i].z += 1
	}
}

// --------------------------------------------------------------------

func NewLine2() *Line2 {
	line := &Line2{
		pts: []Point2{},
	}
	return line
}

func (line *Line2) Resize(n int) {
	line.pts = make([]Point2, n)
}

func (line *Line2) Length() int {
	return len(line.pts)
}

func (line *Line2) AddLast(x, y, z float64) {
	line.pts = append(line.pts, NewPoint2(x, y, z))
}

func (line *Line2) RemoveFirst() {
	line.pts = line.pts[1:len(line.pts)]
}

func (line *Line2) Reproject() {
	n := line.Length()
	for i := 0; i < n; i++ {
		line.pts[i].xyz[0] += 1
		line.pts[i].xyz[1] += 1
		line.pts[i].xyz[2] += 1
	}
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

func (line *Line3) Resize(n int) {
	line.x = make([]float64, n)
	line.y = make([]float64, n)
	line.z = make([]float64, n)
}

func (line *Line3) Length() int {
	return len(line.x)
}

func (line *Line3) AddLast(x, y, z float64) {
	line.x = append(line.x, x)
	line.y = append(line.y, y)
	line.z = append(line.z, z)
}

func (line *Line3) RemoveFirst() {
	n := len(line.x)
	line.x = line.x[1:n]
	line.y = line.y[1:n]
	line.z = line.z[1:n]
}

func (line *Line3) Reproject() {
	n := line.Length()
	for i := 0; i < n; i++ {
		line.x[i] += 1
		line.y[i] += 1
		line.z[i] += 1
	}
}

// --------------------------------------------------------------------

func NewLine4() *Line4 {
	line := &Line4{
		xyz: []float64{},
	}
	return line
}

func (line *Line4) Resize(n int) {
	line.xyz = make([]float64, n*3)
}

func (line *Line4) Length() int {
	return len(line.xyz) / 3
}

func (line *Line4) AddLast(x, y, z float64) {
	line.xyz = append(line.xyz, x)
	line.xyz = append(line.xyz, y)
	line.xyz = append(line.xyz, z)
}

func (line *Line4) RemoveFirst() {
	n := len(line.xyz)
	line.xyz = line.xyz[3:n]
}

func (line *Line4) Reproject() {
	n := line.Length()
	for i := 0; i < n; i++ {
		line.xyz[0] += 1
		line.xyz[1] += 1
		line.xyz[2] += 1
	}
}
