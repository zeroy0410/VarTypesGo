/**
 * Introduction 域敏感-路径长度
 * Level 3
 */

// evaluation information start
// real case = true
// evaluation item = 准确度->域敏感->接口/类->域敏感-路径长度
// bind_url = case/accuracy/field_sensitive/interface_class/field_len_005_T/field_len_005_T.go
// evaluation information end

package main

type A struct {
	data string
}
type B struct {
	a A
}
type C struct {
	b B
}
type D struct {
	c C
}
type E struct {
	d D
}
type F struct {
	e E
}

func main() {
	__taint_src := GetSensitiveData()
	pa := A{
		data: __taint_src,
	}
	var b B
	b.a = pa
	var c C
	c.b = b
	var d D
	d.c = c
	var e E
	e.d = d
	var f F
	f.e = e
	p := f.e.d.c
	q := p.b.a.data
	__taint_sink(q)
}

func __taint_sink(o interface{}) {
}

func GetSensitiveData() string {
	return "sensitive data"
}
