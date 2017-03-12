package play

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
		case "MP3":
			p = &MP3Player{}
		case "MAV":
			p = &MAVPlayer{}
		default:
			fmt.Println("Unsupport music type", mtype);
			return
	}
	p.Play(source)
}