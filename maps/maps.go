package maps

type Dictonary map[string]string

func (d Dictonary) SearchDic(word string) string {
	return d[word]
}
