package role

import ()

type Precure interface {
	MemberOf(Buddy) bool
	HumanName() string
	PrecureName() string
	Challenge() string
	ImageURL() string
}

type Buddy interface {
	String() string
}

type Series interface {
	BeforeTransform()
	Transform()
	AfterTransform()
}

type Base struct {
	Precure
}

func (b *Base) String() string {
	return b.HumanName()
}
