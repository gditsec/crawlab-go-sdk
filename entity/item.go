package entity

type Item map[string]interface{}

type EntityType int

const (
	MEDIA_CONTENT_TYPE EntityType = 0
	PART_CONTENT_TYPE  EntityType = 1
	ARTICLE_TYPE       EntityType = 2
)
