package data


// Index impl: log record pos
// position define by file id and offset
type LogRecordPos struct {
	Fid uint32 	
	Offset int64 
}