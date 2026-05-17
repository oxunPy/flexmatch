package main

type ReceiveMsg struct {
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}

type SendMsg struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
