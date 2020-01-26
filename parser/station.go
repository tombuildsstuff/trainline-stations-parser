package parser

type StationDetails struct {
	// Internal unique identifier.
	Id int

	// Name of the station as it is locally known; see info_* for translations.
	Name string

	// Guaranteed to be unique across all the suggestable stations; see `is_suggestable`.
	Slug string

	// The UIC code of the station.
	UIC string

	// SNCF sometimes uses an UIC code with 8 digits instead of 7. The last digit is a checksum.
	UIC8_SNCF string

	// Coordinates as decimal value.
	Latitude float32

	// Coordinates as decimal value.
	Longitude float32

	// A station can belong to a meta station whose id is this value,
	// i.e. Paris Gare d’Austerlitz belongs to the metastation Paris.
	ParentStationID int

	// 2 letters, ISO 3166-1 alpha-2
	Country string

	// Continent/Country ISO codes.
	TimeZone string

	// Is this station a city? This field is unreliable
	IsCity bool

	// Is this station the Main Station? This field is unreliable
	IsMainStation bool

	// Specifies if the station is related to an airport
	IsAirport bool

	// Specify if the user can input this station.
	IsSuggestable bool

	// Specify if the country should be displayed to disambiguate the station's name.
	CountryHint bool

	// Presence of a SNCF self-service machine at the station.
	SNCFSelfServiceMachine bool

	// Some systems allow stations to be split in two, with two id values.
	//If provided, the station identified by the given value should be considered as the actual station.
	SameAs *int

	// an identifier, which can be used to identify if 2 locations across multiple synchronised sources represent the same location.
	NormalisedCode string
}
