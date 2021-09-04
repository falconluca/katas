package main

import "strconv"

type client interface {
	get(req map[string]string) string
	post(req map[string]string) string
}

// Three Net

type ThreeNetResponse struct {
	code int64
	msg  string
}

func (t *ThreeNetResponse) get(req map[string]string) string {
	url := req["url"]
	return "ThreeNet resp(get): " + url + " " + t.msg + " " + strconv.FormatInt(t.code, 10)
}

func (t *ThreeNetResponse) post(req map[string]string) string {
	auth := req["auth"]
	return "ThreeNet resp(post): " + auth + " " + t.msg + " " + strconv.FormatInt(t.code, 10)
}

// Xfl

type XflResponse struct {
	returnMsg string
}

func (x *XflResponse) get(req map[string]string) string {
	url := req["url"]
	return "Xfl resp(get): " + url + " " + x.returnMsg
}

func (x *XflResponse) post(req map[string]string) string {
	auth := req["auth"]
	return "Xfl resp(post): " + auth + " " + x.returnMsg
}
