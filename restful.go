package httputils

import (
	"net/url"
	"strconv"
	"strings"
)

// Pagination defines the pagination vars
type Pagination map[string]interface{}

// FilterList defines the filterlist vars
type FilterList map[string][]string

// GetFilters parses the URL for a querystring and responds with the restful filters
func GetFilters(u *url.URL) (map[string]interface{}, map[string][]string, error) {
	// Set some default values
	p := make(map[string]interface{})
	p["limit"] = 100
	p["page"] = 1
	p["sort"] = ""

	// Set empty filters
	f := make(map[string][]string)

	queryVals := u.Query()
	if val, ok := queryVals["limit"]; ok {
		limit, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			return p, f, err
		}
		p["limit"] = int(limit)
		delete(queryVals, "limit")
	}

	if val, ok := queryVals["page"]; ok {
		page, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			return p, f, err
		}
		p["page"] = int(page)
		delete(queryVals, "page")
	}

	if val, ok := queryVals["sort"]; ok {
		p["sort"] = val[0]
		delete(queryVals, "sort")
	}

	for key, val := range queryVals {
		f[key] = strings.Split(val[0], ",")
	}

	return p, f, nil
}
