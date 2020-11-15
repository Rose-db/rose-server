package roseServer

func NewServer(t string) Server {
	if t == HttpServerType {
		return newHttpServer()
	}

	return nil
}
