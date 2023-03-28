package aclelement

type Action struct{

	Key string
	Rank int
	
	
}

type ACLElement struct{

	 Element map[string]int 
}

func (actions Actions)GetCode(act Actions){
var acts = make(map[string]int) 

acts["CONNECT"] = 1
acts["GET"] = 2
acts["HEAD"] = 8
acts["OPTION"] = 16
acts["PATCH"] = 32
acts["POST"] = 64
acts["PUT"] = 128
acts["TRACE"] = 256
acts["DELETE"] = 512
}
// TRACE, PUT, POST, PATCH, OPTION, HEAD, DELETE, GET, CONNECT 
// 2^8 , 2^7, 2^6  , 2^5 , 2^4    , 2^3 , 2^2   , 2^1, 2^0