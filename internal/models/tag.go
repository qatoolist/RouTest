package models

import "strings"

type Tag struct {
	Name string
}

type Tags []Tag

// AddTag adds a new tag to the Tags slice.
func (t *Tags) AddTag(tag Tag) {
	*t = append(*t, tag)
}

// GetTagByName returns the tag with the specified name.
func (t Tags) GetTagByName(name string) *Tag {
	for _, tag := range t {
		if tag.Name == name {
			return &tag
		}
	}
	return nil
}

// LoadFromStr loads tags from a comma-separated string.
func (t *Tags) LoadFromStr(tagStr string) {
	tags := strings.Split(tagStr, ",")
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			t.AddTag(Tag{Name: tag})
		}
	}
}
