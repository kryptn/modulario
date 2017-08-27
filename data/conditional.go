package data

import (
	"math/rand"
)

func (e *Engine) BuildDecider(post Post) func() Link {
	deciderMap := map[string]func(Post) func() Link{
		"random":     e.BuildRandomDecider,
		"popularity": e.BuildPopularityDecider,
	}

	decider := deciderMap[post.DeciderType]
	if decider == nil {
		decider = e.BuildRandomDecider
	}

	return decider(post)
}

func (e *Engine) BuildRandomDecider(post Post) func() Link {
	return func() Link {
		link := post.Links[rand.Intn(len(post.Links))]

		link.Accesses += 1
		e.db.Save(&link)
		return link
	}
}

func (e *Engine) BuildPopularityDecider(post Post) func() Link {
	var lastLink *Link
	var result *Link

	return func() Link {
		for _, link := range post.Links {
			lastLink = &link
			if link.Condition.Threshold > e.TotalPostVisits(&post) {
				result = &link
				break
			}
		}
		if result == nil {
			result = lastLink
		}

		result.Accesses += 1
		e.db.Save(&result)
		return *result
	}
}
