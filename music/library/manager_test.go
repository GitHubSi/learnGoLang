package library

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {											//expected boolean expression
		t.Error("New MusicManager failed")
	}
	
	if mm.Len() != 0 {
		t.Error("New MusicManager failed, no empty.")
	}

	m0 := &Music{Id:1, Name:"My Heart Will Go On"}				// too few values in struct initializer
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add failed")
	}

	m := mm.Find(m0.Name)				//(type *MusicManager has no field or method Find, but does have find)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.")
	}
}