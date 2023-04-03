package interfaces

type Meta interface {
	Copy() Meta
	OverrideMeta(override Meta)
	GetParentMeta() Meta
	GetTags() string
}
