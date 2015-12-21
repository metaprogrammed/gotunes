package gotunes

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Response

type ItunesResponse struct {
	ResultCount int         `json:"resultCount"`
	Results     ItunesItems `json:"results"`
}

type ItunesItem struct {
	TrackId           int     `json:"trackId,omitempty"`
	TrackName         string  `json:"trackName,omitempty"`
	TrackExplicitness string  `json:"trackExplicitness,omitempty"`
	TrackCensoredName string  `json:"trackCensoredName,omitempty"`
	TrackViewUrl      string  `json:"trackViewUrl,omitempty"`
	TrackPrice        float32 `json:"trackPrice,omitempty"`

	CollectionId           int     `json:"collectionId,omitempty"`
	CollectionName         string  `json:"collectionName,omitempty"`
	CollectionExplicitness string  `json:"collectionExplicitness,omitempty"`
	CollectionCensoredName string  `json:"collectionCensoredName,omitempty"`
	CollectionViewUrl      string  `json:"collectionViewUrl,omitempty"`
	CollectionPrice        float32 `json:"collectionPrice,omitempty"`

	ArtistId      int    `json:"artistId,omitempty"`
	ArtistName    string `json:"artistName,omitempty"`
	ArtistViewUrl string `json:"artistViewUrl,omitempty"`

	WrapperType           string `json:"wrapperType,omitempty"`
	Kind                  string `json:"kind,omitempty"`
	ContentAdvisoryRating string `json:"contentAdvisoryRating,omitempty"`

	TrackTimeMillis int `json:"trackTimeMillis,omitempty"`
	TrackCount      int `json:"trackCount,omitempty"`
	TrackNumber     int `json:"trackNumber,omitempty"`

	DiscCount  int    `json:"discCount,omitempty"`
	DiscNumber int    `json:"discNumber,omitempty"`
	Currency   string `json:"currency,omitempty"`

	PrimaryGenreName   string   `json:"primaryGenreName,omitempty"`
	Genres             []string `json:"genres,omitempty"`
	SupportedDevices   []string `json:"supportedDevices,omitempty"`
	Advisories         []string `json:"advisories,omitempty"`
	FileSizeBytes      string   `json:"fileSizeBytes,omitempty"`
	Description        string   `json:"description,omitempty"`
	LongDescription    string   `json:"longDescription,omitempty"`
	Country            string   `json:"country,omitempty"`
	ReleaseDate        string   `json:"releaseDate,omitempty"`
	ArtworkUrl100      string   `json:"artworkUrl100,omitempty"`
	ArtworkUrl130      string   `json:"artworkUrl130,omitempty"`
	ArtworkUrl160      string   `json:"artworkUrl160,omitempty"`
	ArtworkUrl512      string   `json:"artworkUrl512,omitempty"`
	FeedUrl            string   `json:"feedUrl,omitempty"`
	PreviewUrl         string   `json:"previewUrl,omitempty"`
	RadioStationUrl    string   `json:"radioStationUrl,omitempty"`
	IsStreamable       bool     `json:"isStreamable,omitempty"`
	ScreenshotUrls     []string `json:"screenshotUrls,omitempty"`
	IPadScreenshotUrls []string `json:"ipadScreenshotUrls,omitempty"`
}
type ItunesItems []ItunesItem
