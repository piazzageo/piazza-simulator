package geo

type FuncType int

const (
	AddLastFunc FuncType = iota
	RemoveFirstFunc
	AddRandomFunc
	RemoveRandomFunc
	ScanFunc
	ReprojectFunc
)

func AddLast_Line1(line *Line1, v float64) {
	pt := Point1{
		x: v,
		y: v,
		z: v,
	}
	line.pts = append(line.pts, pt)
}

func AddLast_Line2(line *Line2, v float64) {
	pt := Point2{
		xyz: [3]float64{float64(v), float64(v), float64(v)},
	}
	line.pts = append(line.pts, pt)
}

func AddLast_Line3(line *Line3, v float64) {
	line.x = append(line.x, v)
	line.y = append(line.y, v)
	line.z = append(line.z, v)
}

func AddLast_Line4(line *Line4, v float64) {
	x, y, z := v, v, v
	line.xyz = append(line.xyz, x)
	line.xyz = append(line.xyz, y)
	line.xyz = append(line.xyz, z)
}

func RemoveFirst_Line1(line *Line1) {
	line.pts = line.pts[1:len(line.pts)]
}

func RemoveFirst_Line2(line *Line2) {
	line.pts = line.pts[1:len(line.pts)]
}

func RemoveFirst_Line3(line *Line3) {
	n := len(line.x)
	line.x = line.x[1:n]
	line.y = line.y[1:n]
	line.z = line.z[1:n]
}

func RemoveFirst_Line4(line *Line4) {
	n := len(line.xyz)
	line.xyz = line.xyz[3:n]
}
