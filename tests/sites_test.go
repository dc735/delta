package delta_test

import (
	"testing"

	"github.com/GeoNet/delta/meta"
)

func TestSites(t *testing.T) {

	var sites meta.SiteList
	t.Log("Load installed sites file")
	{
		if err := meta.LoadList("../network/sites.csv", &sites); err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(sites); i++ {
		for j := i + 1; j < len(sites); j++ {
			if sites[i].Station == sites[j].Station && sites[i].Location == sites[j].Location {
				t.Errorf("site duplication: " + sites[i].Station + "/" + sites[i].Location)
			}
		}
	}

}
