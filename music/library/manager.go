package library

import (
	"errors"
)

type Music struct {
	Id int
	Name string 
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index > m.Len() {
		return nil, errors.New("index out of range");
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *Music {
	if m.Len() == 0 {
		return nil
	}

	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}

	return nil
}

func (m *MusicManager) Add(music *Music)() {				//指针声明没有意义
	m.musics = append(m.musics, *music)					//
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index > m.Len() {
		return nil
	}

	removeMusic := &m.musics[index]

	if index < m.Len() - 1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)		//(type []Music) as type Music in append
	} else if index == 0 {
		m.musics = make([]Music, 0)
	} else {
		m.musics = m.musics[:index-1]
	}

	return removeMusic
}

/*func main() {
	music_1 := Music{Id:1, Name:"music_1"}			//cannot use 1 (type int) as type string in field value
	music_2 := Music{Id:2, Name:"music_2"}

	musicManager := MewMusicManager()
	musicManager.Add(&music_1)						//cannot use music_1 (type Music) as type *Music in argument to musicManager.Add
	musicManager.Add(&music_2)

	fmt.Printf("%v\n", musicManager)

	music_2.Artist = "neojos"
	fmt.Printf("%v\n", music_2)
	fmt.Printf("%v\n", musicManager)

	music_get_2, _ := musicManager.Get(1)				//multiple-value musicManager.Get() in single-value context, err declared and not used
	music_get_2.Artist = "neojos_get_2"					//non-name music_get_2.Artist on left side of :=
	fmt.Printf("%v\n", musicManager)
}*/