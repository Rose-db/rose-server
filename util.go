package main

func UtilHasString(s string, t []string) bool {
	for i := 0; i < len(t); i++ {
		if s == t[i] {
			return true
		}
	}

	return false
}

func Uint8Sort(*[]uint8) {

}
