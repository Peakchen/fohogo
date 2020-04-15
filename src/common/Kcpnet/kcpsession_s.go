package Kcpnet

import (
	"github.com/xtaci/kcp-go"
)

type KcpServerSession struct {
	conn *kcp.UDPSession

	readCh chan []byte
	writeCh chan []byte

}

func NewKcpSvrSession(c *kcp.UDPSession) *KcpServerSession{
	return &KcpServerSession{
		conn: c,
		readCh: make(chan []byte, 1000),
		writeCh: make(chan []byte, 1000),
	}
}

func (this *KcpServerSession) Handler(){
	go this.readloop()
	go this.writeloop()
}

func (this *KcpServerSession) close(){
	this.conn.Close()
}

func (this *KcpServerSession) readloop(){

	defer func(){
		this.close()
	}()

	header := make([]byte, 2)
	
	for {
		this.conn.SetReadDeadline(time.Now().Add(config.readDeadline))
		
		n, err := io.ReadFull(this.conn, header)
		if err != nil {
			Log.Error("read header failed, ip:%v reason:%v size:%v", this.conn.RemoteAddr().String(), err, n)
			return
		}

		size := binary.BigEndian.Uint16(header)
		payload := make([]byte, size)
		n, err = io.ReadFull(conn, payload)
		if err != nil {
			Log.Error("read payload failed, ip:%v reason:%v size:%v", this.conn.RemoteAddr().String(), err, n)
			return
		}

		//是否加个消息队列处理 ?
		go this.read(payload)
	}
}

func (this *KcpServerSession) read([]byte){

}

func (this *KcpServerSession) writeloop(){

	defer func (){
		this.close()
	}()

	for {
		select {
		case data := <-this.writeCh:
			n, err := this.conn.Write(data)
			if err != nil {
				Log.Error("send reply data fail, size: %v, err: %v.", n, err)
				return
			}
		}
	}
}