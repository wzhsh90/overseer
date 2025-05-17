package fetcher

import "io"

// Interface defines the required fetcher functions
type Interface interface {
	//Init should perform validation on fields. For
	//example, ensure the appropriate URLs or keys
	//are defined or ensure there is connectivity
	//to the appropriate web service.
	Init() error
	// Checksum check the download file is correct
	Checksum(binHash string) bool
	//Fetch should check if there is an updated
	//binary to fetch, and then stream it back the
	//form of an io.Reader. If io.Reader is nil,
	//then it is assumed there are no updates. Fetch
	//will be run repeatedly and forever. It is up the
	//implementation to throttle the fetch frequency.
	Fetch() (io.Reader, error)

	FetchCmd() (io.Reader, error)
}
