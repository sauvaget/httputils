package httputils

import (
	"fmt"
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilters(t *testing.T) {
	testCases := []struct {
		url        string
		pagination map[string]int
		filters    map[string][]string
	}{
		{
			"http://test.domain/ressource?filter1=foo&filter2=bar&filter3=foo,bar",
			map[string]int{
				"limit": 100,
				"page":  1,
			},
			map[string][]string{
				"filter1": {"foo"},
				"filter2": {"bar"},
				"filter3": {"foo", "bar"},
			},
		},
		{
			"http://test.domain/ressource?limit=25&filter1=foo&filter2=bar&filter3=foo,bar",
			map[string]int{
				"limit": 25,
				"page":  1,
			},
			map[string][]string{
				"filter1": {"foo"},
				"filter2": {"bar"},
				"filter3": {"foo", "bar"},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("GetFilters %d", i), func(st *testing.T) {
			url, _ := url.Parse(tc.url)
			p, f, _ := GetFilters(url)
			assert.Equal(t, tc.pagination["limit"], p["limit"])
			assert.Equal(t, tc.pagination["page"], p["page"])
			// assert.Equal(t, len(tc.expected.filters), len(f.filters))
			log.Printf("%+v\n", f)
		})
	}
}
