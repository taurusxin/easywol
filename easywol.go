package easywol

import (
	"errors"
	"net"
)

type MagicPacket [102]byte

func CreateMagicPacket(macAddr string) (packet MagicPacket, err error) {
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return packet, err
	}

	if len(mac) != 6 {
		return packet, errors.New("invalid MAC address")
	}

	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6
	for i := 0; i < 16; i++ {
		copy(packet[offset:], mac)
		offset += 6
	}

	return packet, nil
}

func SendUDPPacket(packet MagicPacket, addr string) (err error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet[:])
	return err
}

func (packet MagicPacket) Send(addr string) error {
	return SendUDPPacket(packet, addr + ":9")
}

func (packet MagicPacket) SendPort(addr string, port string) error {
	return SendUDPPacket(packet, addr + ":" + port)
}
