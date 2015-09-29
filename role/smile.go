package role

import (
	"github.com/mattn/go-precure/util"
)

type Smile struct {
	Buddies []Precure
}

func (s Smile) BeforeTransform() {
}

func (s Smile) AfterTransform() {
	if len(s.Buddies) == 5 {
		util.Say("五つの心が導く未来!")
		util.Say("輝け! スマイルプリキュア!")
	}
}
