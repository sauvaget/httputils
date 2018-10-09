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
		url      string
		expected restFilters
	}{
		{
			"http://test.domain/ressource?filter1=foo&filter2=bar&filter3=foo,bar",
			restFilters{
				pagination{limit: 100, page: 1},
				map[string][]string{
					"filter1": {"foo"},
					"filter2": {"bar"},
					"filter3": {"foo", "bar"},
				},
			},
		},
		{
			"http://test.domain/ressource?filter1=foo&filter2=bar&limit=25",
			restFilters{
				pagination{limit: 25, page: 1},
				map[string][]string{
					"filter1": {"foo"},
					"filter2": {"bar"},
				},
			},
		},
		{
			"http://test.domain/ressource?filter1=foo&filter2=bar&page=2&limit=25",
			restFilters{
				pagination{limit: 25, page: 2},
				map[string][]string{
					"filter1": {"foo"},
					"filter2": {"bar"},
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("GetFilters %d", i), func(st *testing.T) {
			url, _ := url.Parse(tc.url)
			f, _ := GetFilters(url)
			assert.Equal(t, tc.expected.pagination.limit, f.pagination.limit)
			assert.Equal(t, tc.expected.pagination.page, f.pagination.page)
			assert.Equal(t, len(tc.expected.filters), len(f.filters))
			log.Printf("%+v\n", f)
		})
	}
}
