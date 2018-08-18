package string

//based on https://play.golang.org/p/Xcno8QurAn & http://dabase.com/e/15006/

func (l *Lines) DeleteEmptyLines() {
	var r Lines
	for _, byteStr := range *l {
		str := string(byteStr)
		if str != "" && str != "\n" {
			r = append(r, byteStr)
		}
	}
	*l = r
}
