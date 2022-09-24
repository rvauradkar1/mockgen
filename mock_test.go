package mock

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/rvauradkar1/fuse"
)

func Test_pop(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	fmt.Println("+v", info)
	if len(info.Fields) != 8 {
		t.Errorf("length of populateFields should have been %d, but was %d", 7, len(info.Fields))
	}
	fmt.Println(len(info.Imports))
	if len(info.Imports) != 14 {
		t.Errorf("length of imports should have been %d, but was %d", 14, len(info.Imports))
	}
	fmt.Println(len(info.Funcs))
	if len(info.Funcs) != 3 {
		t.Errorf("length of funcs should have been %d, but was %d", 2, len(info.Funcs))
	}
	fmt.Println(len(info.Deps))
	for i := 0; i < len(info.Deps); i++ {
		fmt.Println(info.Deps[i])
	}
	if len(info.Deps) != 16 {
		t.Errorf("length of deps should have been %d, but was %d", 14, len(info.Deps))
	}
	if info.Typ != reflect.TypeOf(L1{}) {
		t.Errorf("type should have been %s, but was %s", "lvl1", info.Typ)
	}
}

func Test_shouldAdd(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	types := make(map[reflect.Type]*typeInfo, 0)
	types[reflect.TypeOf(L1{})] = info
	b := shouldAdd(types, info)
	if b == true {
		t.Errorf("should NOT have been added for %T, same types cannot be added", info.Typ)
	}

	info = populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	info2 := populateInfo(Component{Instance: &L2{}, Basepath: "./lvl1"})
	types = make(map[reflect.Type]*typeInfo, 0)
	types[reflect.TypeOf(L2{})] = info
	b = shouldAdd(types, info2)
	if b == false {
		t.Errorf("should have been added for %T, same should be added", info.Typ)
	}
}

func Test_pkg(t *testing.T) {
	s := pkg("")
	if s != "" {
		t.Errorf("pkg name should have been blank, but instead was %s", s)
	}
	s = pkg("a.b")
	if s != "a" {
		t.Errorf("pkg name should have been blank, but instead was %s", s)
	}
}

func Test_printOutParams(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	s := printOutParams(info.Funcs[0].Params)
	if s != "(string,*int)" {
		t.Errorf("should have been '%s', but was '%s'", "(string,*int)", s)
	}
}

func Test_printInParams(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	s := printInParams(info.Funcs[0].Params)
	if s != "i1 int,f2 float32" {
		t.Errorf("should have been '%s', but was '%s'", "i1 int,f2 float32", s)
	}
	s = printInParams(info.Funcs[2].Params)
	if s != "pf1 *float32" {
		t.Errorf("should have been '%s', but was '%s'", "pf1 *float32", s)
	}
}

func Test_printInNames(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	info.Funcs[0].Params[1].InName = "p1"
	info.Funcs[0].Params[1].Input = true
	info.Funcs[0].Params[2].InName = "p2"
	info.Funcs[0].Params[2].Input = true
	fmt.Println(len(info.Funcs[0].Params))
	s := printInNames(info.Funcs[0].Params)
	fmt.Println(s)
	if s != " p1, p2" {
		t.Errorf("should have been '%s', but was '%s'", " p1, p2", s)
	}
}

func Test_paramSlice(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	info.Funcs[0].Params[1].InName = "p1"
	info.Funcs[0].Params[1].Input = true
	info.Funcs[0].Params[2].InName = "p2"
	info.Funcs[0].Params[2].Input = true
	fmt.Println(len(info.Funcs[0].Params))
	s := paramSlice(info.Funcs[0].Params)
	fmt.Println(s)
	if s != "[]interface{}{p1 ,p2 }" {
		t.Errorf("should have been '%s', but was '%s'", " p1, p2", s)
	}
}

func Test_printImports(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	types := make(map[reflect.Type]*typeInfo)
	types[reflect.TypeOf(L1{})] = info
	s := printImports(types)
	fmt.Println(s)
	if !strings.Contains(s, "time") {
		t.Errorf("should have contained '%s'", "time")
	}
}

func Test_receiver(t *testing.T) {
	s := receiver(nil)
	if s != "error" {
		t.Errorf("should have errored out with message %s, but was instead %s", "error", s)
	}

	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	s = receiver(info.Funcs[0])
	if s != "v " {
		t.Errorf("should have been '%s', but was '%s'", "v ", s)
	}
	s = receiver(info.Funcs[1])
	if s != "p *" {
		t.Errorf("should have been '%s', but was '%s'", "p *", s)
	}
}

func Test_printFields(t *testing.T) {
	info := populateInfo(Component{Instance: &L1{}, Basepath: "./lvl1"})
	s := printFields(info.Fields)
	fmt.Println(s)
	if !strings.Contains(s, "S1 string\ntime time.Duration\nTime2 time.Duration") {
		t.Errorf("should have contained '%s'", "S1 string\ntime time.Duration\nTime2 time.Duration")
	}
}

/*
func Test_gen(t *testing.T) {
	m := MockGen{}
	comps := make([]Component, 0)
	comps = append(comps, Component{Instance: &lvl1.L1{}, Basepath: "./lvl1"})
	comps = append(comps, Component{Instance: &lvl2.L2{}, Basepath: "./lvl1/lvl2"})
	comps = append(comps, Component{Instance: &lvl3.L3{}, Basepath: "./lvl1/lvl2/lvl3"})
	m.Comps = comps
	m.Generate()

	t1 := typeInfo{}
	fmt.Println(t1)

}

*/

func Test_findDeps(t *testing.T) {
	fi := fieldInfo{StructField: reflect.StructField{Tag: "_deps"}}
	deps := findDeps(&fi)
	if len(deps) > 0 {
		t.Errorf("length of deps should be 0")
	}

	fi = fieldInfo{StructField: reflect.StructField{Tag: `_deps:"comp1"`}}
	deps = findDeps(&fi)
	if len(deps) != 1 {
		t.Errorf("length of deps should be 1")
	}
	if deps[0] != "comp1" {
		t.Errorf("comp should have been comp1 but was %s", deps[0])
	}

	fi = fieldInfo{StructField: reflect.StructField{Tag: `_deps:"comp1,comp2"`}}
	deps = findDeps(&fi)
	if len(deps) != 2 {
		t.Errorf("length of deps should be 1")
	}
	if deps[0] != "comp1" {
		t.Errorf("comp should have been comp1 but was %s", deps[0])
	}
	if deps[1] != "comp2" {
		t.Errorf("comp should have been comp2 but was %s", deps[0])
	}
}

/*
func Test_register(t *testing.T) {
	m := New("mock")
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "lvl1", Instance: &lvl1.L1{}})
	entries = append(entries, fuse.Entry{Name: "lvl2", Instance: &lvl2.L2{}})
	entries = append(entries, fuse.Entry{Name: "lvl3", Instance: &lvl3.L3{}})
	errors := m.Register(entries)
	fmt.Println("errors = ", errors)
	m.Generate()
}
*/

/*
func Test_gen(t *testing.T) {
	m := MockGen{}
	comps := make([]Component, 0)
	//fuse.Entry{Name: "OrdCtrl", Instance: &ctrl.OrderController{}})

	comps = append(comps, Component{Instance: &L1{}, Basepath: "./"})
	comps = append(comps, Component{Instance: &L2{}, Basepath: "./"})
	comps = append(comps, Component{Instance: &L3{}, Basepath: "./"})
	m.Comps = comps
	m.Generate()

	t1 := typeInfo{}
	fmt.Println(t1)

}

*/

func Test_register(t *testing.T) {
	m := New("mock")
	entries := make([]fuse.Entry, 0)
	entries = append(entries, fuse.Entry{Name: "OrdCtrl", Instance: &L1{}})
	entries = append(entries, fuse.Entry{Name: "CartSvc", Instance: &L2{}})
	entries = append(entries, fuse.Entry{Name: "AuthSvc", Instance: &L3{}})
	errors := m.Register(entries)
	fmt.Println("errors = ", errors)
	m.Generate()
}
