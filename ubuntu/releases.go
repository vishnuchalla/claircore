package ubuntu

import (
	"github.com/quay/claircore"
)

type Release string

const (
	Artful  Release = "artful"
	Bionic  Release = "bionic"
	Cosmic  Release = "cosmic"
	Disco   Release = "disco"
	Precise Release = "precise"
	Trusty  Release = "trusty"
	Xenial  Release = "xenial"
)

var AllReleases = map[Release]struct{}{
	Artful:  struct{}{},
	Bionic:  struct{}{},
	Cosmic:  struct{}{},
	Disco:   struct{}{},
	Precise: struct{}{},
	Trusty:  struct{}{},
	Xenial:  struct{}{},
}
var ReleaseToVersionID = map[Release]string{
	Artful:  "17.10",
	Bionic:  "18.04",
	Cosmic:  "18.10",
	Disco:   "19.04",
	Precise: "12.04",
	Trusty:  "14.04",
	Xenial:  "16.04",
}

var artfulDist = &claircore.Distribution{
	Name:            "Ubuntu",
	Version:         "17.10 (Artful Aardvark)",
	DID:             "ubuntu",
	PrettyName:      "Ubuntu 17.10",
	VersionID:       "17.10",
	VersionCodeName: "artful",
}

var bionicDist = &claircore.Distribution{
	Name:            "Ubuntu",
	Version:         "18.04.3 LTS (Bionic Beaver)",
	DID:             "ubuntu",
	PrettyName:      "Ubuntu 18.04.3 LTS",
	VersionID:       "18.04",
	VersionCodeName: "bionic",
}

var cosmicDist = &claircore.Distribution{
	Name:            "Ubuntu",
	Version:         "18.10 (Cosmic Cuttlefish)",
	DID:             "ubuntu",
	VersionID:       "18.10",
	VersionCodeName: "cosmic",
	PrettyName:      "Ubuntu 18.10",
}

var discoDist = &claircore.Distribution{
	Name:            "Ubuntu",
	Version:         "19.04 (Disco Dingo)",
	DID:             "ubuntu",
	VersionID:       "19.04",
	VersionCodeName: "disco",
	PrettyName:      "Ubuntu 19.04",
}

var preciseDist = &claircore.Distribution{
	Name:       "Ubuntu",
	Version:    "12.04.5 LTS, Precise Pangolin",
	DID:        "ubuntu",
	VersionID:  "12.04",
	PrettyName: "Ubuntu precise (12.04.5 LTS)",
}

var trustyDist = &claircore.Distribution{
	Name:       "Ubuntu",
	Version:    "14.04.6 LTS, Trusty Tahr",
	DID:        "ubuntu",
	PrettyName: "Ubuntu 14.04.6 LTS",
	VersionID:  "14.04",
}

var xenialDist = &claircore.Distribution{
	Name:            "Ubuntu",
	Version:         "14.04.6 LTS, Trusty Tahr",
	DID:             "ubuntu",
	PrettyName:      "Ubuntu 16.04.6 LTS",
	VersionID:       "16.04",
	VersionCodeName: "xenial",
}

func releaseToDist(r Release) *claircore.Distribution {
	switch r {
	case Artful:
		return artfulDist
	case Bionic:
		return bionicDist
	case Cosmic:
		return bionicDist
	case Disco:
		return discoDist
	case Precise:
		return preciseDist
	case Trusty:
		return trustyDist
	case Xenial:
		return xenialDist
	default:
		// return empty dist
		return &claircore.Distribution{}
	}
}
