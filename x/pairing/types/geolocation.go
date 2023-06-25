package types

import (
	"fmt"
	"strconv"
	"strings"

	commontypes "github.com/lavanet/lava/common/types"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	planstypes "github.com/lavanet/lava/x/plans/types"
	"github.com/spf13/cast"
)

var (
	GLS = planstypes.Geolocation_GLS
	AF  = planstypes.Geolocation_AF
	AS  = planstypes.Geolocation_AS
	AU  = planstypes.Geolocation_AU
	EU  = planstypes.Geolocation_EU
	USC = planstypes.Geolocation_USC
	USE = planstypes.Geolocation_USE
	USW = planstypes.Geolocation_USW
	GL  = planstypes.Geolocation_GL
)

const DEFAULT_GLOBAL_GEOLOCATION = 0xFFFF

func ExtractGeolocations(g uint64) ([]planstypes.Geolocation, string) {
	locations := make([]planstypes.Geolocation, 0)
	var printStr string

	if g&uint64(GL) == DEFAULT_GLOBAL_GEOLOCATION {
		locations = append(locations, GL)
		printStr += "GL"
		return locations, printStr
	}

	if g&uint64(AF) == uint64(AF) {
		locations = append(locations, AF)
		printStr += "AF,"
	}

	if g&uint64(AS) == uint64(AS) {
		locations = append(locations, AS)
		printStr += "AS,"
	}

	if g&uint64(AU) == uint64(AU) {
		locations = append(locations, AU)
		printStr += "AU,"
	}

	if g&uint64(EU) == uint64(EU) {
		locations = append(locations, EU)
		printStr += "EU,"
	}

	if g&uint64(USC) == uint64(USC) {
		locations = append(locations, USC)
		printStr += "USC,"
	}

	if g&uint64(USE) == uint64(USE) {
		locations = append(locations, USE)
		printStr += "USE,"
	}

	if g&uint64(USW) == uint64(USW) {
		locations = append(locations, USW)
		printStr += "USW,"
	}

	return locations, strings.TrimSuffix(printStr, ",")
}

func IsValidGeoEnum(s string) (uint64, bool) {
	val, ok := planstypes.Geolocation_value[s]
	if ok && val != planstypes.Geolocation_value["GLS"] && val > 0 {
		return uint64(val), true
	}

	return uint64(val), false
}

func GetCurrentGlobalGeolocation() uint64 {
	var globalGeo int32
	for k := range planstypes.Geolocation_name {
		if planstypes.Geolocation_name[k] == "GL" {
			continue
		}
		globalGeo += k
	}

	return uint64(globalGeo)
}

func CalculateGeoUintFromStrings(geos []string) uint64 {
	var uintGeo uint64
	for _, geo := range geos {
		geoVal := planstypes.Geolocation_value[geo]
		uintGeo += uint64(geoVal)
	}

	return uintGeo
}

func HandleEndpointsAndGeolocationArgs(endpArg []string, geoArg string) (endp []epochstoragetypes.Endpoint, geo uint64, err error) {
	for _, endpointStr := range endpArg {
		splitted := strings.Split(endpointStr, ",")
		if len(splitted) != 3 {
			return nil, 0, fmt.Errorf("invalid argument format in endpoints, must be: HOST:PORT,useType,geolocation HOST:PORT,useType,geolocation, received: %s", endpointStr)
		}

		// geolocation of an endpoint can be uint or a string representing a single geo region
		var geoloc uint64
		geoloc, valid := IsValidGeoEnum(splitted[2])
		if !valid {
			geoloc, err = strconv.ParseUint(splitted[2], 10, 64)
			if err != nil {
				return nil, 0, fmt.Errorf("invalid argument format in endpoints, geolocation must be a number or valid geolocation string")
			}
		}

		// verify that the endpoint's geolocation represents a single geo region
		geoRegions, _ := ExtractGeolocations(geoloc)
		if len(geoRegions) != 1 {
			return nil, 0, fmt.Errorf("invalid geolocation for endpoint, must represent one region")
		}

		// if the user specified global ("GL"), append the endpoint in all possible geolocations
		if geoloc == uint64(planstypes.Geolocation_value["GL"]) {
			for geoName, geoVal := range planstypes.Geolocation_value {
				if geoName == "GL" || geoName == "GLS" {
					continue
				}
				endpoint := epochstoragetypes.Endpoint{IPPORT: splitted[0], UseType: splitted[1], Geolocation: uint64(geoVal)}
				endp = append(endp, endpoint)
			}
		} else {
			endpoint := epochstoragetypes.Endpoint{IPPORT: splitted[0], UseType: splitted[1], Geolocation: geoloc}
			endp = append(endp, endpoint)
		}
	}

	// handle the case of string geolocation (example: "EU,AF,AS")
	splitted := strings.Split(geoArg, ",")
	strEnums := commontypes.RemoveDuplicatesFromSlice(splitted)
	for _, s := range strEnums {
		g, valid := IsValidGeoEnum(s)
		if valid {
			// if one of the endpoints is global, assign global value and break
			if g == uint64(planstypes.Geolocation_GL) {
				geo = GetCurrentGlobalGeolocation()
				break
			}
			// for non-global geolocations constructs the final geolocation uint value (addition works because we
			// removed duplicates and the geo regions represent a bitmap)
			geo += g
		}
	}

	// geolocation is not a list of enums, try to parse it as an uint
	if geo == 0 {
		geo, err = cast.ToUint64E(geoArg)
		if err != nil {
			return nil, 0, err
		}
	}
	return endp, geo, nil
}
