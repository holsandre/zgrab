package opcua

import (
	"errors"
	"net"
	"strings"
)

// main function , currently only sends HEL and Open Secure Channel Request with Security Mode "None"
func GetOPCUAData(logStruct *OPCUALog, connection net.Conn) error {

	HELqueryBytes := BuildHel()	
	OSCRqueryBytes := BuildOscr()

	bytesWritten, err := connection.Write(HELqueryBytes)
	if bytesWritten != len(HELqueryBytes) {
		return errors.New("Unable to write OPC UA HEL query...")
	}
	if err != nil {
		return err
	}

	readBuffer := make([]byte, 256)
	length, err := connection.Read(readBuffer)
	if length > len(readBuffer) {
		return errors.New("Need a bigger buffer to read OPCUA response")
	}
	if err != nil {
		return err
	}
	readBuffer=readBuffer[:length]
	// receive ACK, first sign that it is an OPCUA server
	if string(readBuffer[0:3]) == "ACK" {

		bytesWritten, err := connection.Write(OSCRqueryBytes)
		if bytesWritten != len(OSCRqueryBytes) {
			return errors.New("Unable to writeOPC UA Open Secure Channel Request...")
		}
		if err != nil {
			return err
		}

		readBuffer := make([]byte, 512)
		length, err := connection.Read(readBuffer)
		if length > len(readBuffer) {
			return errors.New("Need a bigger buffer to read OPCUA Open Secure Channel response")
		}
		if err != nil {
			return err
		}
		readBuffer=readBuffer[:length]
			// server answer equals OPNF, OPCUA = TRUE
			if string(readBuffer[0:4]) == "OPNF" {
				logStruct.IsOPCUA = true

				// Test if the Security Policy URL is in the packet, if so we can be sure its OPCUA
				if (strings.Contains(string(readBuffer), "http://opcfoundation.org/UA/SecurityPolicy#None")){
					logStruct.SecurityPolicyUri = "http://opcfoundation.org/UA/SecurityPolicy#None"
					logStruct.ServerNonce, logStruct.ServerProtocolVersion, logStruct.ServerTimestamp = ParseOscr(readBuffer)
				} else {
					// Policy URL not in the packet, Message Security Mode None not supported
					logStruct.SecurityPolicyUri = "Message Security Mode None not supported"
				}
			}
		
		
	}

	return nil
}
