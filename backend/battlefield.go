package backend

import "reflect"

type Battlefield struct {
	You   Player
	Enemy Player
	Zone  Zones
}

func (b Battlefield) Equal(c Battlefield) bool {
	return reflect.DeepEqual(b, c)
}
