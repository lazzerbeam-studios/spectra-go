package bunny

var Video *BunnyVideoClient
var Storage *BunnyStorageClient

type BunnyVideoClient struct {
	apiKey    string
	libraryID string
}

type BunnyStorageClient struct {
	apiKey      string
	storageZone string
	region      string
	cdnHost     string
}

func SetBunnyVideoClient(apiKey string, libraryID string) {
	Video = &BunnyVideoClient{
		apiKey:    apiKey,
		libraryID: libraryID,
	}
}

func SetBunnyStorageClient(apiKey string, storageZone string, region string, cdnHost string) {
	Storage = &BunnyStorageClient{
		apiKey:      apiKey,
		storageZone: storageZone,
		region:      region,
		cdnHost:     cdnHost,
	}
}
