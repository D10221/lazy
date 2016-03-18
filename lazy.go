package lazy

import (
"sync"
"html/template"
)

type Lazy struct {
	value interface{}
	done bool
	fty func() interface{}
}

func NewLazy(fun func() interface{})  *Lazy {
	return &Lazy{fty: fun}
}

var mutex = &sync.Mutex{}

func (l *Lazy) Value() interface{} {
	if l.done {
		return l.value
	}
	mutex.Lock()
	l.done = true
	l.value = l.fty()
	mutex.Unlock()
	return l.value
}

type LazyTemplate struct {
	_value *template.Template
	_done  bool
	_fty   func() *template.Template
}


func NewLazyTemplate(fun func() *template.Template) *LazyTemplate {
	return &LazyTemplate{_fty: fun}
}

func (l *LazyTemplate) Value() *template.Template {
	if l._done {
		return l._value
	}
	mutex.Lock()
	l._done = true
	l._value = l._fty()
	mutex.Unlock()
	return l._value
}



