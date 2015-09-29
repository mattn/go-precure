package precure

import (
	"github.com/mattn/go-precure/role"
)

type CureBeauty struct {
	role.Base
}

func (p *CureBeauty) HumanName() string {
	return `青木れいか`
}

func (p *CureBeauty) PrecureName() string {
	return `キュアビューティ`
}

func (p *CureBeauty) Age() int {
	return 14
}

func (p *CureBeauty) Challenge() string {
	return `しんしんと降りつもる清き心! キュアビューティ!`
}

func (p *CureBeauty) MemberOf(b role.Buddy) bool {
	return b.String() == `スマイル`
}

func init() {
	RegisterPrecure(&CureBeauty{})
}
