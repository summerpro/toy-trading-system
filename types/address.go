package types

type Address string

const SystemAddress Address = "System"

func (addr Address) Validate() bool {
	if len(addr) > 20 {
		return false
	}
	if addr == "" {
		return false
	}
	return true
}

func (addr Address) Bytes() []byte {
	return []byte(addr)
}
