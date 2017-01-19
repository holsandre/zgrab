package opcua

import (
	"encoding/hex"
)

// hello packet in hex
var hexhel string = "00000000" + // Proto Version = 0
	"00000100" + // ReceivedBufferSize = 65636
	"00000100" + // SendBufferSize = 65636
	"00000001" + // MaxMsgSize = 167
	"88130000" + // MaxChunkCount = 5000
	"00000000" // Endpoint URL = OPC Null String

// function to build hello (+8 because of the first 8 bytes)
func BuildHel() []byte {
	hel, _ := hex.DecodeString("48454c46" + message_size((len(hexhel)/2)+8) + hexhel)
	return hel
}

// same for open secure channel request
var hexoscr string = "00000000" + // Secure Channel ID
	"2f000000687474703a2f2f6f7063666f756e646174696f6e2e6f72672f55412f5365637572697479506f6c696379234e6f6e65" + // Length.ULInt32|Security Policy Uri
	"ffffffff" + // Sender Cert
	"ffffffff" + // Receiver Cert
	"33000000" + // Seq Nr
	"03000000" + // Request ID
	"0100be01" + // Enc mask = Open Secure Channel Request
	"0000" + // Authentikation Token
	"0000000000000000" + // Timestamp (only for diagnostic purpose)
	"00000000" + // Request Handle
	"00000000" + // Return Diagnostics
	"ffffffff" + // Audit Entry ID
	"00000000" + // Timeout Hint
	"000000" + // Additional Header
	"00000000" + // Client Protocol Version
	"00000000" + // Security Token Request Type
	"01000000" + // Message Security Mode = 1 = None
	"01000000" + // Some Padding 
	"00" + // Client Nonce 00
	"e0930400" // Requested Lifetime

// same as hello (+8 because of the first 8 bytes)
func BuildOscr() []byte {
	oscr, _ := hex.DecodeString("4f504e46" + message_size((len(hexoscr)/2)+8) + hexoscr)
	return oscr
}
