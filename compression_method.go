package dtls

import "github.com/goldenrye/dtls/pkg/protocol"

func defaultCompressionMethods() []*protocol.CompressionMethod {
	return []*protocol.CompressionMethod{
		{},
	}
}
