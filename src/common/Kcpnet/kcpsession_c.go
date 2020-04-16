package Kcpnet

type KcpClientSession struct {
	conn 		*kcp.UDPSession

	readCh 		chan []byte
	writeCh 	chan []byte

	remoteAddr 	string
	//message pack
	pack 		IMessagePack
}

func NewKcpClientSession(c *kcp.UDPSession, Msgpack IMessagePack) *KcpServerSession{
	return &KcpServerSession{
		conn: c,
		readCh: make(chan []byte, 1000),
		writeCh: make(chan []byte, 1000),
		remoteAddr: this.conn.RemoteAddr().String(),
		pack: Msgpack,
	}
}

func (this *KcpClientSession) GetRemoteAddr() string {
	return this.remoteAddr
}

func (this *KcpClientSession) Handler() {

}

func (this *KcpClientSession) readloop(){

}