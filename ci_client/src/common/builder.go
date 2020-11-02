package common

import "strconv"

type Doc struct {
	info string
}

func New() *Doc {
	return &Doc{info: ""}
}

func (doc *Doc) Start() {
	doc.addSplit()
	doc.info += "{"
}

func (doc *Doc) End() {
	runes := []rune(doc.info)
	r := runes[len(runes)-1]
	if ',' == r {
		runes[len(runes)-1] = '}'
		doc.info = string(runes)
	} else {
		doc.info += "}"
	}
}

func (doc *Doc) StartTitle(title string) {
	doc.addSplit()
	doc.info = doc.info + "\"" + title + "\"" + ":" + "{"
}

func (doc *Doc) StartArr(title string) {
	doc.addSplit()
	doc.info = doc.info + "\"" + title + "\"" + ":" + "["
}

func (doc *Doc) EndArr() {
	runes := []rune(doc.info)
	r := runes[len(runes)-1]
	if ',' == r {
		runes[len(runes)-1] = ']'
		doc.info = string(runes)
	} else {
		doc.info += "]"
	}
}

func (doc *Doc) addSplit() {
	runes := []rune(doc.info)
	if len(runes) <= 0 {
		return
	}
	r := runes[len(runes)-1]
	if '}' == r || ']' == r {
		doc.info += ","
	}
}
func (doc *Doc) AddStr(k string, v string) {
	doc.addSplit()
	doc.info = doc.info + "\"" + k + "\"" + ":" + "\"" + v + "\"" + ","
}
func (doc *Doc) AddInt(k string, v int) {
	doc.addSplit()
	doc.info = doc.info + "\"" + k + "\"" + ":" + strconv.Itoa(v) + ","
}
func (doc *Doc) AddFloat64(k string, v float64) {
	doc.addSplit()
	doc.info = doc.info + "\"" + k + "\"" + ":" + strconv.FormatFloat(v, 'f', -1, 64) + ","
}
func (doc *Doc) AddEntity(k string, v string) {
	doc.addSplit()
	doc.info = doc.info + "\"" + k + "\"" + ":" + v + ","
}

func (doc *Doc) String() string {
	return doc.info
}

func (doc *Doc) Bytes() []byte {
	return []byte(doc.info)
}
