package main

import "rose/rose"

type HttpServer struct {
	Host string
	Port string
	Rose *rose.Rose
}

func (s *HttpServer) Start() {
	
}