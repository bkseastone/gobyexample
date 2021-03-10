package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CommonFieldMixin struct {
	ent.Mixin
}

func (CommonFieldMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("ctime").
			Default(time.Now).Immutable(),
		field.Time("utime").
			Default(time.Now).UpdateDefault(time.Now),
		field.Time("dtime").Optional().
			Nillable(),
	}
}

// 索引字段列表
func (CommonFieldMixin) Index() []ent.Index {
	return nil
}
