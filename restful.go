package httputils

import (
	"net/url"
	"strconv"
	"strings"
)

type restFilters struct {
	pagination pagination
	filters    map[string][]string
}

type pagination struct {
	limit int64
	page  int64
}

// GetFilters parses the URL for a querystring and responds with the restful filters
func GetFilters(u *url.URL) (map[string]int, map[string][]string, error) {
	// Set some default values
	var p map[string]int
	p["limit"] = 100
	p["page"] = 1

	// Set empty filters
	var f map[string][]string

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

	for key, val := range queryVals {
		f[key] = strings.Split(val[0], ",")
	}

	return p, f, nil
}
