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
func GetFilters(u *url.URL) (restFilters, error) {
	// Set some default values
	var rf restFilters
	rf.pagination.limit = 100
	rf.pagination.page = 1

	queryVals := u.Query()
	if val, ok := queryVals["limit"]; ok {
		limit, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			return rf, err
		}
		rf.pagination.limit = limit
		delete(queryVals, "limit")
	}

	if val, ok := queryVals["page"]; ok {
		page, err := strconv.ParseInt(val[0], 10, 64)
		if err != nil {
			return rf, err
		}
		rf.pagination.page = page
		delete(queryVals, "page")
	}

	rf.filters = make(map[string][]string)
	for key, val := range queryVals {
		rf.filters[key] = strings.Split(val[0], ",")
	}

	return rf, nil
}
