package openrtb

import (
	"encoding/json"

	"github.com/mxmCherry/openrtb/native"
	"github.com/mxmCherry/openrtb/native/request"
)

// 3.2.9 Object: Native
//
// This object represents a native type impression.
// Native ad units are intended to blend seamlessly into the surrounding content (e.g., a sponsored Twitter or Facebook post).
// As such, the response must be well-structured to afford the publisher fine-grained control over rendering.
//
// The Native Subcommittee has developed a companion specification to OpenRTB called the Dynamic Native Ads API.
// It defines the request parameters and response markup structure of native ad units.
// This object provides the means of transporting request parameters as an opaque string so that the specific parameters can evolve separately under the auspices of the Dynamic Native Ads API.
// Similarly, the ad markup served will be structured according to that specification.
//
// The presence of a Native as a subordinate of the Imp object indicates that this impression is offered as a native type impression.
// At the publisher’s discretion, that same impression may also be offered as banner, video, and/or audio by also including as Imp subordinates objects of those types.
// However, any given bid for the impression must conform to one of the offered types.
type Native struct {
	// Attribute:
	//   request
	// Type:
	//   string; required
	// Description:
	//   Request payload complying with the Native Ad Specification.
	Request string `json:"request"`

	// Attribute:
	//   ver
	// Type:
	//   string; recommended
	// Description:
	//   Version of the Dynamic Native Ads API to which request
	//   complies; highly recommended for efficient parsing.
	Ver string `json:"ver,omitempty"`

	// Attribute:
	//   api
	// Type:
	//   integer array
	// Description:
	//   List of supported API frameworks for this impression. Refer to
	//   List 5.6. If an API is not explicitly listed, it is assumed not to be
	//   supported.
	API []APIFramework `json:"api,omitempty"`

	// Attribute:
	//   sequence
	// Type:
	//   integer array
	// Description:
	//   Blocked creative attributes. Refer to List 5.3.
	BAttr []CreativeAttribute `json:"battr,omitempty"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`

	// HACK for android SDK

	// Field:
	//   layout
	// Scope:
	//   recommended in 1.0, deprecated/removed in 1.2
	// Type:
	//   integer
	// Description:
	//   The Layout ID of the native ad unit.
	//   See the Table of Layout IDs below.
	Layout native.Layout `json:"layout,omitempty"`

	// Field:
	//   adunit
	// Scope:
	//   recommended in 1.0, deprecated/removed in 1.2
	// Type:
	//   integer
	// Description:
	//   The Ad unit ID of the native ad unit.
	//   See Table of Ad Unit IDs below for a list of supported core ad units.
	AdUnit native.AdUnit `json:"adunit,omitempty"`

	// Field:
	//   context
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Description:
	//   The context in which the ad appears.
	//   See Table of Context IDs below for a list of supported context types.
	Context native.ContextType `json:"context,omitempty"`

	// Field:
	//   contextsubtype
	// Scope:
	//   optional
	// Type:
	//   integer
	// Description:
	//   A more detailed context in which the ad appears.
	//   See Table of Context SubType IDs below for a list of supported context subtypes.
	ContextSubType native.ContextSubType `json:"contextsubtype,omitempty"`

	// Field:
	//   plcmttype
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Description:
	//   The design/format/layout of the ad unit being offered.
	//   See Table of Placement Type IDs below for a list of supported placement types.
	PlcmtType native.PlacementType `json:"plcmttype,omitempty"`

	// Field:
	//   plcmtcnt
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   1
	// Description:
	//   The number of identical placements in this Layout.
	//   Refer Section 8.1 Multiplacement Bid Requests for further detail.
	PlcmtCnt int64 `json:"plcmtcnt,omitempty"`

	// Field:
	//   seq
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   0
	// Description:
	//   0 for the first ad, 1 for the second ad, and so on.
	//   Note this would generally NOT be used in combination with plcmtcnt - either you are auctioning multiple identical placements (in which case plcmtcnt>1, seq=0) or you are holding separate auctions for distinct items in the feed (in which case plcmtcnt=1, seq=>=1)
	Seq int64 `json:"seq,omitempty"`

	// Field:
	//   assets
	// Scope:
	//   required
	// Type:
	//   array of objects
	// Description:
	//   An array of Asset Objects.
	//   Any bid response must comply with the array of elements expressed in the bid request.
	Assets []request.Asset `json:"assets"`

	// Field:
	//   aurlsupport
	// Scope:
	//   optional
	// Type:
	//   int
	// Default:
	//   0
	// Description:
	//   Whether the supply source / impression supports returning an assetsurl instead of an asset object.
	//   0 or the absence of the field indicates no such support.
	AURLSupport int8 `json:"aurlsupport,omitempty"`

	// Field:
	//   durlsupport
	// Scope:
	//   optional
	// Type:
	//   int
	// Default:
	//   0
	// Description:
	//   Whether the supply source / impression supports returning a dco url instead of an asset object.
	//   0 or the absence of the field indicates no such support.
	//   Beta feature.
	DURLSupport int8 `json:"durlsupport,omitempty"`

	// Field:
	//   eventtrackers
	// Scope:
	//   optional
	// Type:
	//   array of objects
	// Description:
	//   Specifies what type of event tracking is supported - see Event Trackers Request Object
	EventTrackers []request.EventTracker `json:"eventtrackers,omitempty"`

	// Field:
	//   privacy
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Default:
	//   0
	// Description:
	//   Set to 1 when the native ad supports buyer-specific privacy notice.
	//   Set to 0 (or field absent) when the native ad doesn’t support custom privacy links or if support is unknown.
	Privacy int8 `json:"privacy,omitempty"`
}
