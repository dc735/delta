package delta_test

import (
	"sort"
	"testing"

	"github.com/GeoNet/delta/meta"
)

func TestMetSensors(t *testing.T) {

	var metsensors meta.InstalledMetSensorList
	t.Log("Load deployed metsensors file")
	{
		if err := meta.LoadList("../install/metsensors.csv", &metsensors); err != nil {
			t.Fatal(err)
		}
	}

	t.Log("Check for metsensors installation equipment overlaps")
	{
		installs := make(map[string]meta.InstalledMetSensorList)
		for _, s := range metsensors {
			_, ok := installs[s.Model]
			if ok {
				installs[s.Model] = append(installs[s.Model], s)

			} else {
				installs[s.Model] = meta.InstalledMetSensorList{s}
			}
		}

		var keys []string
		for k, _ := range installs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := installs[k]

			for i, n := 0, len(v); i < n; i++ {
				for j := i + 1; j < n; j++ {
					switch {
					case v[i].Serial != v[j].Serial:
					case v[i].End.Before(v[j].Start):
					case v[i].Start.After(v[j].End):
					case v[i].End.Equal(v[j].Start):
					case v[i].Start.Equal(v[j].End):
					default:
						t.Errorf("metsensors %s at %-5s has mark %s overlap between %s and %s",
							v[i].Model, v[i].Serial, v[i].Mark, v[i].Start.Format(meta.DateTimeFormat), v[i].End.Format(meta.DateTimeFormat))
					}
				}
			}
		}
	}

	t.Log("Check for missing metsensor marks")
	{
		var marks meta.MarkList

		if err := meta.LoadList("../network/marks.csv", &marks); err != nil {
			t.Fatal(err)
		}

		keys := make(map[string]interface{})

		for _, m := range marks {
			keys[m.Code] = true
		}

		for _, c := range metsensors {
			if _, ok := keys[c.Mark]; ok {
				continue
			}
			t.Errorf("unable to find metsensor mark %-5s", c.Mark)
		}
	}

	var assets meta.AssetList
	t.Log("Load metsensor assets file")
	{
		if err := meta.LoadList("../assets/metsensors.csv", &assets); err != nil {
			t.Fatal(err)
		}
	}

	t.Log("Check for metsensor assets")
	{
		for _, r := range metsensors {
			var found bool
			for _, a := range assets {
				if a.Model != r.Model {
					continue
				}
				if a.Serial != r.Serial {
					continue
				}
				found = true
			}
			if !found {
				t.Errorf("unable to find metsensor asset: %s [%s]", r.Model, r.Serial)
			}
		}
	}

}
