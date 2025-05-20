package nets

import "net"

type Nets struct {
	IP   string
	Port int
	The  net.Conn
}

func NewNets(ip string, port int) *Nets {
	return &Nets{IP: ip, Port: port}
}

func (tar *Nets) Connect() {
	tar.The, _ = net.Dial("tcp", tar.IP+":"+string(tar.Port))
}
