package lib

import (
	"sync"
	"sync/atomic"
)

type SignBase interface {
	Set()
	Unset()
	GetCount() uint64
}

type Sign struct {
	count uint64
	sign  chan uint64
}

func (self *Sign) Set() {
	if self.sign == nil {
		self.sign = make(chan uint64, 1)
	}
	self.sign <- atomic.AddUint64(&self.count, uint64(1))
}

func (self *Sign) Unset() {
	if self.sign == nil {
		return
	}
	if len(self.sign) == 0 {
		return
	}
	<-self.sign
}

func (self *Sign) GetCount() uint64 {
	return self.count
}

type RWSign struct {
	count uint64
	sign  *sync.RWMutex
}

func (self *RWSign) Set() {
	if self.sign == nil {
		self.sign = new(sync.RWMutex)
	}
	self.sign.Lock()
	atomic.AddUint64(&self.count, uint64(1))
}

func (self *RWSign) Unset() {
	if self.sign == nil {
		return
	}
	self.sign.Unlock()
}

func (self *RWSign) GetCount() uint64 {
	return self.count
}

func (self *RWSign) RSet() {
	if self.sign == nil {
		self.sign = new(sync.RWMutex)
	}
	self.sign.RLock()
}

func (self *RWSign) RUnset() {
	if self.sign == nil {
		return
	}
	self.sign.RUnlock()
}

func NewSign() *Sign {
	return &Sign{count: 0}
}

func NewRWSign() *RWSign {
	return &RWSign{count: 0}
}
