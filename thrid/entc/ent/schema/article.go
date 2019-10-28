package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// 文章模型的字段
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(512).Comment("发布时间"),
		field.Uint32("category_id").Comment("类别id").StructTag(`json:"categoryId,omitempty"`),
		field.Uint32("view_count").Comment("	阅读数").Default(0).
			StructTag(`json:"viewCount"`),
		field.Uint32("star_count").Comment("	喜欢数").Default(0).
			StructTag(`json:"starCount"`),
		field.String("excerpt").MaxLen(1024).Comment("简介,摘要").Optional().Nillable(),
		field.String("content").MaxLen((2 << 24) - 1).Comment("内容").Optional().Nillable(),
		field.Uint32("author_uid").Comment("	作者id").Default(0).StructTag(`json:"authorUid"`),
		field.Time("release_time").
			Default(time.Now).StructTag(`json:"releaseTime"`),
	}
}

// 文章模型的关联性
func (Article) Edges() []ent.Edge {
	return nil
}

// 索引字段列表
func (Article) Index() []ent.Index {
	return nil
}

// 配置
func (Article) Config() ent.Config {
	return ent.Config{
		Table: "buff_article",
	}
}

// 混用
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonFieldMixin{},
	}
}
