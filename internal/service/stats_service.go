package service

import (
	"sort"

	"clipboard/internal/model"
)

type StatsOverview struct {
	TotalEntries  int            `json:"total_entries"`
	ByContentType map[string]int `json:"by_content_type"`
	ByStatus      map[string]int `json:"by_status"`
	TotalGroups   int            `json:"total_groups"`
	TopTags       []TagCount     `json:"top_tags"`
}

type TagCount struct {
	TagID string `json:"tag_id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (s *Service) StatsOverview() (*StatsOverview, error) {
	entries := s.store.ListEntries()
	groups := s.store.ListGroups()
	entryTags := s.store.ListEntryTags()

	byContentType := make(map[string]int)
	byStatus := make(map[string]int)
	for _, e := range entries {
		byContentType[e.ContentType]++
		byStatus[e.Status]++
	}

	tagCounts := make(map[string]int)
	for _, et := range entryTags {
		tagCounts[et.TagID]++
	}

	top := make([]TagCount, 0)
	for tagID, count := range tagCounts {
		tag, err := s.store.GetTag(tagID)
		name := tagID
		if err == nil && tag != nil {
			name = tag.Name
		}
		top = append(top, TagCount{TagID: tagID, Name: name, Count: count})
	}
	sort.Slice(top, func(i, j int) bool {
		if top[i].Count != top[j].Count {
			return top[i].Count > top[j].Count
		}
		return top[i].TagID < top[j].TagID
	})
	if len(top) > 5 {
		top = top[:5]
	}

	return &StatsOverview{
		TotalEntries:  len(entries),
		ByContentType: byContentType,
		ByStatus:      byStatus,
		TotalGroups:   len(groups),
		TopTags:       top,
	}, nil
}

func (s *Service) ListTagsByEntry(entryID string) ([]*model.Tag, error) {
	entryTags := s.store.ListEntryTagsByEntry(entryID)
	result := make([]*model.Tag, 0, len(entryTags))
	for _, et := range entryTags {
		tag, err := s.store.GetTag(et.TagID)
		if err != nil {
			continue
		}
		result = append(result, tag)
	}
	return result, nil
}
