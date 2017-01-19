package opcua

import (
	"strconv"
	"time"
	"encoding/binary"
	)
	
// calculate message size
func message_size(i int) string {
	if i < 255 {
		return string(inttohex(i)) + "000000"
	} 
	return inttohex(i)[2:4] + inttohex(i)[0:2] + "0000"
}

func inttohex(i int) string {
	i64 := int64(i)
	return strconv.FormatInt(i64, 16) // base 16 for hexadecima
}

func converttime(data []byte) string {
	//var delta = time.Date(1970-369, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
    t := binary.LittleEndian.Uint64(data)
	ts := time.Unix(0, int64(t)*100+time.Date(1970-369, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano())
	return ts.Format("Mon Jan 2 15:04:05 MST 2006")
}