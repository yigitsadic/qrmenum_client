package shared

import "errors"

var (
	UnableToFetchFromCMS         = errors.New("unable to make request to CMS")
	UnableMarshalResponseFromCMS = errors.New("unable to decode JSON from CMS")
	KeyNotFoundOnInMemoryStore   = errors.New("key not found on in memory store")
	KeyNotFoundOnCMS             = errors.New("key not found on in memory store")
)
