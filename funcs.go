package timezones

import (
	"time"
)

func LocationByZonename(zonename string) (l *time.Location) {
	z := ZoneinfoByName(zonename)
	if z == nil {
		l = time.UTC
		return
	}
	l = time.FixedZone(z.Name, z.Offset*60*60)
	return
}

func ZoneinfoByName(zonename string) (r *Zoneinfo) {
	for _, z := range AllZones {
		if zonename == z.Name {
			r = z
			return
		}
	}
	return
}

func ZoneinfoByOffset(offset int) (r *Zoneinfo) {
	for _, z := range AllZones {
		if offset == z.Offset {
			r = z
			return
		}
	}
	return
}
