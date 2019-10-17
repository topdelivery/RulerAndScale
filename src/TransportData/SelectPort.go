package TransportData

import (
	"github.com/jacobsa/go-serial/serial"
	"strconv"
	"sync"
)

var Ports = InitPortStorage()

type PortStorage struct {
	Ports map[string]*Port
	mx    sync.Mutex
}

func InitPortStorage() *PortStorage {
	return &PortStorage{
		Ports: make(map[string]*Port),
	}
}

func (p *PortStorage) GetPort(device string) *Port {
	p.mx.Lock()
	defer p.mx.Unlock()

	return p.Ports[device]
}

func (p *PortStorage) ResetPort(device string) {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.Ports[device] = nil
}

func (p *PortStorage) SetPort(port *Port, device string) {

	if port == nil {
		return
	}

	p.mx.Lock()
	defer p.mx.Unlock()

	p.Ports[device] = port
}

func SelectPort() {

	println("Поиск портов")
	portClass := []string{"/dev/ttyS", "/dev/ttyACM", "/dev/ttyUSB"}

	for {
		for _, nameClass := range portClass {
			for i := 0; i < 20; i++ {

				portName := nameClass + strconv.Itoa(i)

				if Ports.GetPort("scale") == nil {
					//Ports.SetPort(FindScale(portName), "scale")
				}

				if Ports.GetPort("ruler") == nil {
					Ports.SetPort(FindRuler(portName), "ruler")
				}

				if Ports.GetPort("scale") != nil && Ports.GetPort("ruler") != nil {
					println("Все устройства подключены.")
				}
			}
		}
	}
}

//func FindScale(portName string) (port *Port) {
//
//	weightConfig := &serial.OpenOptions{
//		Name: portName,
//		Baud:        4800,
//		Parity:      'E',
//		ReadTimeout: time.Millisecond * 200,
//	}
//
//	port = &Port{Name: portName, Config: weightConfig}
//	connect := port.Connect()
//	if connect == nil {
//		return nil
//	}
//
//	connect.Flush()
//
//	_, err := connect.Write([]byte{0x48})
//	if err != nil {
//		connect.Close()
//		return nil
//	}
//
//	buf := make([]byte, 2)
//	n, err := connect.Read(buf)
//
//	if err != nil {
//		connect.Close()
//		return nil
//	} else {
//		if n == 2 && (buf[0] == 128 || buf[0] == 192) {
//			println("Весы подключены к порту " + portName)
//			return port
//		} else {
//			return nil
//		}
//	}
//}

func FindRuler(portName string) *Port {

	rulerConfig := serial.OpenOptions{
		PortName:              portName,
		BaudRate:              19200,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       0,
		InterCharacterTimeout: 200,
	}

	connect, err := serial.Open(rulerConfig)
	if err != nil {
		//println("serial.Open: %v", err.Error())
		return nil
	}

	_, err = connect.Write([]byte{0x95})
	if err != nil {
		connect.Close()
		return nil
	}

	buf := make([]byte, 4)
	n, err := connect.Read(buf)

	if err != nil {
		connect.Close()
		return nil
	} else {
		if n == 4 && buf[0] == 127 {
			println("Линейка подключена к порту " + portName)

			connect.Close()
			rulerConfig.MinimumReadSize = 12
			connect, err = serial.Open(rulerConfig)
			if err != nil {
				println("serial.Open: %v", err.Error())
				return nil
			}

			return &Port{Name: portName, Config: &rulerConfig, Connection: connect}
		} else {
			return nil
		}
	}
}
