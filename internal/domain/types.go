package domain

import (
	"time"
)

type Identity string

type ClientID string

func (cid ClientID) String() string {
	return string(cid)
}

type AccountID string

func (aid AccountID) String() string {
	return string(aid)
}

//type Dispatchers map[string]map[string]string
type Dispatchers interface{}
type CanonicalFacts interface{}
type Tags interface{}

type MessageMetadata struct {
	LatestMessageID string
	LatestTimestamp time.Time
}

type ConnectorClientState struct {
	Account         AccountID
	ClientID        ClientID
	CanonicalFacts  CanonicalFacts
	Dispatchers     Dispatchers
	Tags            Tags
	MessageMetadata MessageMetadata
}
