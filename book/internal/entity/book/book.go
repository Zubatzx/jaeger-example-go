package book

// Metadata ...
type Metadata struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

// Err ...
type Err struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
	Code   int    `json:"code"`
}

// Response ...
type Response struct {
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
	Error    Err         `json:"error"`
}

// BookDetail ...
type BookDetail struct {
	Showname string `json:"showname"`
	Showtime string `json:"showtime"`
}
