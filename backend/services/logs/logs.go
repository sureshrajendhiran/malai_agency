package logs

import "log"

func Logs(val ...interface{}) {
	// for _, x := range val {
	// 	log.Println(x)
	// }
	log.Println(val)
	return
}
