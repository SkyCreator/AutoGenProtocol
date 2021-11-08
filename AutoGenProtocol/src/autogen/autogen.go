package autogen

func AutoGenProtocol() {
	m := new(ProtocolGenManager)
	m.init()
	m.run()
}
