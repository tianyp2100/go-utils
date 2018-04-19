package tsgutils

/*
 other utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

import "reflect"

/*
  exception handler, replace other language:

	try {
		...
	} catch (Exception e) {
		e.printStackTrace();
	} finally {
		...
	}

	Usage:
		Try(func() {
			panic("exe ...")
		}).Catch(1, func(e interface{}) {
			Stdout("iii" ,e)
		}).Catch("", func(e interface{}) {
			Stdout("sss", e)
		}).Finally(func() {
			Stdout("fff")
		})

*/

type tryStruct struct {
	catchs map[reflect.Type]ExceptionHandler
	hold   func()
}

type ExceptionHandler func(interface{})

func Try(f func()) *tryStruct {
	return &tryStruct{
		catchs: make(map[reflect.Type]ExceptionHandler),
		hold:   f,
	}
}

func (t *tryStruct) Catch(e interface{}, f ExceptionHandler) *tryStruct {
	t.catchs[reflect.TypeOf(e)] = f
	return t
}

func (t *tryStruct) Finally(f func()) {
	defer func() {
		if e := recover(); nil != e {
			if h, ok := t.catchs[reflect.TypeOf(e)]; ok {
				h(e)
			} else {
				f()
			}
		}
	}()
	t.hold()
}
