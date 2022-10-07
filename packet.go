package dtls

import "github.com/goldenrye/dtls/pkg/protocol/recordlayer"

type packet struct {
	record                   *recordlayer.RecordLayer
	shouldEncrypt            bool
	resetLocalSequenceNumber bool
}
