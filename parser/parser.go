package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Parse(fileName string) (*[]StationDetails, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// id;name;slug;uic;uic8_sncf;latitude;longitude;parent_station_id;country;time_zone;is_city;is_main_station;is_airport;is_suggestable;country_hint;main_station_hint;sncf_id;sncf_tvs_id;sncf_is_enabled;idtgv_id;idtgv_is_enabled;db_id;db_is_enabled;busbud_id;busbud_is_enabled;distribusion_id;distribusion_is_enabled;flixbus_id;flixbus_is_enabled;cff_id;cff_is_enabled;leoexpress_id;leoexpress_is_enabled;obb_id;obb_is_enabled;ouigo_id;ouigo_is_enabled;trenitalia_id;trenitalia_is_enabled;trenitalia_rtvt_id;ntv_rtiv_id;ntv_id;ntv_is_enabled;hkx_id;hkx_is_enabled;renfe_id;renfe_is_enabled;atoc_id;atoc_is_enabled;benerail_id;benerail_is_enabled;westbahn_id;westbahn_is_enabled;sncf_self_service_machine;same_as;info:de;info:en;info:es;info:fr;info:it;info:nb;info:nl;info:cs;info:da;info:hu;info:ja;info:ko;info:pl;info:pt;info:ru;info:sv;info:tr;info:zh;normalised_code
	stations := make([]StationDetails, 0)
	scanner := bufio.NewScanner(file)

	firstLine := true
	for scanner.Scan() {
		// the first line is the header line - we can skip it
		if firstLine {
			firstLine = false
			continue
		}

		chunk := strings.Split(scanner.Text(), ";")
		stations = append(stations, StationDetails{
			Id:                     parseInt(chunk[0]),
			Name:                   chunk[1],
			Slug:                   chunk[2],
			UIC:                    chunk[3],
			UIC8_SNCF:              chunk[4],
			Latitude:               parseFloat(chunk[5]),
			Longitude:              parseFloat(chunk[6]),
			ParentStationID:        parseInt(chunk[7]),
			Country:                chunk[8],
			TimeZone:               chunk[9],
			IsCity:                 parseBool(chunk[10]),
			IsMainStation:          parseBool(chunk[11]),
			IsAirport:              parseBool(chunk[12]),
			IsSuggestable:          parseBool(chunk[13]),
			CountryHint:            parseBool(chunk[14]),
			SNCFSelfServiceMachine: parseBool(chunk[15]),
			SameAs:                 parseNilableInt(chunk[16]),
			NormalisedCode:         chunk[17],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &stations, nil
}

func parseBool(input string) bool {
	b, e := strconv.ParseBool(input)
	if e != nil {
		return false
	}

	return b
}

func parseFloat(input string) float32 {
	f, e := strconv.ParseFloat(input, 32)
	if e != nil {
		return 0.0
	}

	return float32(f)
}

func parseInt(input string) int {
	i, e := strconv.Atoi(input)
	if e != nil {
		return -1
	}

	return i
}

func parseNilableInt(input string) *int {
	if input == "" {
		return nil
	}

	i, e := strconv.Atoi(input)
	if e != nil {
		return nil
	}

	return &i
}
