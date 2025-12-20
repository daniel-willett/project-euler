package main
import("fmt";"math")

func main() {
	for c:=1000; c>1; c--{
		for b:=1000-c; b>1; b--{
			for a:=b; a>1; a--{
				if a+b+c==1000{
					aSquared := int(math.Pow(float64(a),2))
					bSquared := int(math.Pow(float64(b),2))
					cSquared := int(math.Pow(float64(c),2))
					if aSquared+bSquared==cSquared{
						fmt.Println(a*b*c)
					}
				}
			}
		}
	}
}
