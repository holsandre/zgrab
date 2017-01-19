package opcua

import "encoding/hex"
// parse open secure channel request
func ParseOscr(s []byte) (string, string,string) {
	// get server_nonce & server_prtocol_version & server_timestamp from the end of the string (len-x) 
	// this is done because there are some fields with variable length in the middle of the packet
	if (len(s) > 29){
	server_prtocol_version := hex.EncodeToString(s[len(s)-29 : len(s)-25])
	server_timestamp := converttime(s[len(s)-17:len(s)-9])
	server_nonce := hex.EncodeToString(s[len(s)-1 : len(s)])
	return server_nonce, server_prtocol_version, server_timestamp
	}
	return "","",""
}
