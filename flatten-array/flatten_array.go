// Package flatten returns a flattened array
package flatten

// Flatten returns a flattened array
func Flatten(a interface{}) []interface{} {
	f := make([]interface{}, 32)
	p := 0
	for {
		x := a.([]interface{})
		for i := 0; i < len(x); i++ {
			// switch depending on type of interface
			switch y := x[i].(type) {
			case int:
				f[p] = y
				p++
			case []interface{}:
				// fmt.Println("Array interface")
				y = Flatten(y)
				// fmt.Println(y)
				for j := 0; j < len(y); j++ {
					// fmt.Println(f)
					f[p] = y[j]
					p++
				}

			default:
				// fmt.Printf("Something else: %v", y)
			}
		}
		f = f[:p]
		break
	}
	return f
}
