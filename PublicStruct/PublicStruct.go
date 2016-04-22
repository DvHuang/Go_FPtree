package PublicStruct

type MapPrefixPath struct {
	Path  map[string]int
	Count int
	Data string
}

/*
//嵌套map赋值方法
type EmStr struct  {
	firstmap *

}

func EmMap_(外围map的类型，map，内里map的类型) {
	m := map[string]map[string]string{}
	mm, ok := m["kkk"]
	if !ok {
		mm = make(map[string]string)
		m["kkk"] = mm
	}
	mm[k1k1k1] = "sssss"

}*/
