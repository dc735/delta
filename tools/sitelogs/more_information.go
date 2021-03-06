package main

import (
	"encoding/xml"
)

type MoreInformation struct {
	XMLName xml.Name `xml:"moreInformation"`

	PrimaryDataCenter             string `xml:"mi:primaryDataCenter"`
	SecondaryDataCenter           string `xml:"mi:secondaryDataCenter"`
	UrlForMoreInformation         string `xml:"mi:urlForMoreInformation"`
	HardCopyOnFile                string `xml:"mi:hardCopyOnFile"`
	SiteMap                       string `xml:"mi:siteMap"`
	SiteDiagram                   string `xml:"mi:siteDiagram"`
	HorizonMask                   string `xml:"mi:horizonMask"`
	MonumentDescription           string `xml:"mi:monumentDescription"`
	SitePictures                  string `xml:"mi:sitePictures"`
	Notes                         string `xml:"mi:notes"`
	AntennaGraphicsWithDimensions string `xml:"mi:antennaGraphicsWithDimensions"`
	InsertTextGraphicFromAntenna  string `xml:"mi:insertTextGraphicFromAntenna"`
}
