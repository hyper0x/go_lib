package go_lib

import (
	"sync"
)

type SignBase interface {
	Set()
	Unset()
	GetCount() uint64
}

type Sign struct {
	count      int64
	sign       *sync.Mutex
	innerMutex *sync.Mutex
}

func (self *Sign) Set() {
	if self.sign == nil {
		self.sign = new(sync.Mutex)
	}
	self.sign.Lock()
	if self.innerMutex == nil {
		self.innerMutex = new(sync.Mutex)
	}
	self.innerMutex.Lock()
	self.count += 1
	self.innerMutex.Unlock()
}

func (self *Sign) Unset() {
	if self.sign == nil {
		return
	}
	self.sign.Unlock()
	if self.innerMutex == nil {
		self.innerMutex = new(sync.Mutex)
	}
	self.innerMutex.Lock()
	if self.count > 0 {
		self.count -= 1
	}
	self.innerMutex.Unlock()
}

func (self *Sign) GetCount() uint64 {
	return uint64(self.count)
}

type RWSign struct {
	count      int64
	sign       *sync.RWMutex
	innerMutex *sync.Mutex
}

func (self *RWSign) Set() {
	if self.sign == nil {
		self.sign = new(sync.RWMutex)
	}
	self.sign.Lock()
	if self.innerMutex == nil {
		self.innerMutex = new(sync.Mutex)
	}
	self.innerMutex.Lock()
	self.count += 1
	self.innerMutex.Unlock()
}

func (self *RWSign) Unset() {
	if self.sign == nil {
		return
	}
	self.sign.Unlock()
	if self.innerMutex == nil {
		self.innerMutex = new(sync.Mutex)
	}
	self.innerMutex.Lock()
	if self.count > 0 {
		self.count -= 1
	}
	self.innerMutex.Unlock()
}

func (self *RWSign) GetCount() uint64 {
	return uint64(self.count)
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
