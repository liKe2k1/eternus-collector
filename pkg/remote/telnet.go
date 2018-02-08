package remote

import (
	"github.com/ziutek/telnet"
	"log"
	"reflect"
	"time"
	"unsafe"
)

type Telnet struct {
	host string
	port int
	user string
	pass string
	t    *telnet.Conn
}

func NewTelnet(host, user, pass string) *Telnet {
	Telnet := new(Telnet)
	Telnet.host = host
	Telnet.user = user
	Telnet.pass = pass
	return Telnet
}

const timeout = 10 * time.Second

func (c *Telnet) CheckErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func (c *Telnet) expect(t *telnet.Conn, d ...string) {
	c.CheckErr(t.SetReadDeadline(time.Now().Add(timeout)))
	c.CheckErr(t.SkipUntil(d...))
}

func (c *Telnet) sendln(t *telnet.Conn, s string) {
	c.CheckErr(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	c.CheckErr(err)
}

func (c *Telnet) Open() (bool, error) {
	t, err := telnet.Dial("tcp", c.host+":23")
	c.t = t
	c.CheckErr(err)
	c.t.SetUnixWriteMode(true)

	c.expect(c.t, "Login:")
	c.sendln(c.t, c.user)
	c.expect(c.t, "Password:")
	c.sendln(c.t, c.pass)
	c.expect(c.t, "CLI>")

	return true, nil
}

func (c *Telnet) Close() {
	c.sendln(c.t, "exit")

	c.t.Close()
}

func (c *Telnet) Send(command string) ([]byte, error) {
	var data []byte
	var err error
	c.sendln(c.t, command)
	data, err = c.t.ReadBytes('>')
	c.CheckErr(err)
	return data, nil
}

func (c *Telnet) BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
