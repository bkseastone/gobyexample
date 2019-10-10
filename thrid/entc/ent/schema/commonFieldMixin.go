package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

type CommonFieldMixin struct{}

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
