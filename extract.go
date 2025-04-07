package placeholder

import (
	"slices"
	"strings"
)

// This list was derived using cmd/properties/main.go

var extract_properties = []string{
	"edtf:deprecated",
	"geom:area",
	"geom:bbox",
	"geom:latitude",
	"geom:longitude",
	"gn:pop",
	"gn:population",
	"iso:country",
	"lbl:bbox",
	"lbl:latitude",
	"lbl:longitude",
	"meso:pop",
	"mz:is_current",
	"ne:iso_a2",
	"ne:iso_a3",
	"ne:pop_est",
	"qs:gn_pop",
	"qs:photo_sum",
	"qs:pop",
	"statoids:population",
	"wk:population",
	"wof:country_alpha3",
	"wof:hierarchy",
	"wof:id",
	"wof:label",
	"wof:megacity",
	"wof:name",
	"wof:parent_id",
	"wof:placetype",
	"wof:population",
	"wof:shortcode",
	"wof:superseded_by",
	"zs:pop10",
}

// IsExtractProperty returns a boolean value indicating whether the key 'k' should be included in the
// Placeholder "extract" document.
func IsExtractProperty(k string) bool {

	if strings.HasPrefix(k, "name:") {
		return true
	}

	if slices.Contains(extract_properties, k) {
		return true
	}

	return false
}
