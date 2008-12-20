// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Element interface {
}

type Vector struct {
	nelem int;
	elem []Element;
}

func New() *Vector {
	v := new(*Vector);
	v.nelem = 0;
	v.elem = new([10]Element);
	return v;
}

func (v *Vector) At(i int) Element {
	return v.elem[i];
}

func (v *Vector) Insert(e Element) {
	v.elem[v.nelem] = e;
	v.nelem++;
}

type I struct { val int; };  // BUG: can't be local;

func main() {
	i0 := new(*I); i0.val = 0;
	i1 := new(*I); i1.val = 11;
	i2 := new(*I); i2.val = 222;
	i3 := new(*I); i3.val = 3333;
	i4 := new(*I); i4.val = 44444;
	v := New();
	print("hi\n");
	v.Insert(i4);
	v.Insert(i3);
	v.Insert(i2);
	v.Insert(i1);
	v.Insert(i0);
	for i := 0; i < v.nelem; i++ {
		var x *I;
		x = v.At(i);
		print(i, " ", x.val, "\n");  // prints correct list
	}
	for i := 0; i < v.nelem; i++ {
		print(i, " ", v.At(i).(*I).val, "\n");
	}
}
/*
bug027.go:50: illegal types for operand
	(<Element>I{}) CONV (<I>{})
bug027.go:50: illegal types for operand
	(<Element>I{}) CONV (<I>{})
*/
