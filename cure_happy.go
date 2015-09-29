package precure

import (
	"github.com/mattn/go-precure/role"
)

type CureHappy struct {
	role.Base
}

func (p *CureHappy) HumanName() string {
	return `星空みゆき`
}

func (p *CureHappy) PrecureName() string {
	return `キュアハッピー`
}

func (p *CureHappy) Age() int {
	return 14
}

func (p *CureHappy) Challenge() string {
	return `キラキラ輝く未来の光! キュアハッピー!`
}

func (p *CureHappy) Color() int {
	return 200
}

func (p *CureHappy) ImageURL() string {
	return `http://www.toei-anim.co.jp/tv/precure/images/character/c1_1.jpg`
}

func (p *CureHappy) MemberOf(b role.Buddy) bool {
	return b.String() == `スマイル`
}

func init() {
	RegisterPrecure(&CureHappy{})
}
