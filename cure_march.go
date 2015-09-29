package precure

import (
	"github.com/mattn/go-precure/role"
)

type CureMarch struct {
	role.Base
}

func (p *CureMarch) HumanName() string {
	return `緑川なお`
}

func (p *CureMarch) PrecureName() string {
	return `キュアマーチ`
}

func (p *CureMarch) Age() int {
	return 14
}

func (p *CureMarch) Challenge() string {
	return `勇気リンリン直球勝負! キュアマーチ!`
}

func (p *CureMarch) Color() int {
	return 34
}

func (p *CureMarch) ImageURL() string {
	return `http://www.toei-anim.co.jp/tv/precure/images/character/c4_1.jpg`
}

func (p *CureMarch) MemberOf(b role.Buddy) bool {
	return b.String() == `スマイル`
}

func init() {
	RegisterPrecure(&CureMarch{})
}
