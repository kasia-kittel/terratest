package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomRegion(t *testing.T) {
	t.Parallel()

	randomRegion := GetRandomRegion(t, nil, nil)
	assertLooksLikeRegionName(t, randomRegion)
}

func TestGetRandomZone(t *testing.T) {
	t.Parallel()

	randomZone := GetRandomZone(t, nil, nil)
	assertLooksLikeZoneName(t, randomZone)
}

func TestGetRandomRegionExcludesForbiddenRegions(t *testing.T) {
	t.Parallel()

	approvedRegions := []string{"asia-east1", "asia-northeast1", "asia-south1", "asia-southeast1", "australia-southeast1", "europe-north1", "europe-west1", "europe-west2", "europe-west3", "northamerica-northeast1", "southamerica-east1", "us-central1", "us-east1", "us-east4", "us-west2"}
	forbiddenRegions := []string{"europe-west4", "us-west1"}

	for i := 0; i < 1000; i++ {
		randomRegion := GetRandomRegion(t, approvedRegions, forbiddenRegions)
		assert.NotContains(t, forbiddenRegions, randomRegion)
	}
}

func TestGetRandomZoneExcludesForbiddenZones(t *testing.T) {
	t.Parallel()

	approvedZones := []string{"us-east1-b", "us-east1-c", "us-east1-d", "us-east4-a", "us-east4-b", "us-east4-c", "us-west2-a", "us-west2-b", "us-west2-c", "us-central1-f", "europe-west2-b"}
	forbiddenZones := []string{"us-east1-a", "europe-west1-a", "europe-west2-a", "europe-west2-c"}

	for i := 0; i < 1000; i++ {
		randomZone := GetRandomZone(t, approvedZones, forbiddenZones)
		assert.NotContains(t, forbiddenZones, randomZone)
	}
}

func TestGetAllGcpRegions(t *testing.T) {
	t.Parallel()

	regions := GetAllGcpRegions(t)

	// The typical account had access to 17 regions as of August, 2018: https://cloud.google.com/compute/docs/regions-zones/
	assert.True(t, len(regions) >= 17, "Number of regions: %d", len(regions))
	for _, region := range regions {
		assertLooksLikeRegionName(t, region)
	}
}

func TestGetAllGcpZones(t *testing.T) {
	t.Parallel()

	zones := GetAllGcpZones(t)

	// The typical account had access to 52 zones as of August, 2018: https://cloud.google.com/compute/docs/regions-zones/
	assert.True(t, len(zones) >= 52, "Number of zones: %d", len(zones))
	for _, zone := range zones {
		assertLooksLikeZoneName(t, zone)
	}
}

func assertLooksLikeRegionName(t *testing.T, regionName string) {
	assert.Regexp(t, "[a-z]+-[a-z]+?[[:digit:]]+", regionName)
}

func assertLooksLikeZoneName(t *testing.T, zoneName string) {
	assert.Regexp(t, "[a-z]+-[a-z]+?[[:digit:]]+-[a-z]{1}", zoneName)
}
