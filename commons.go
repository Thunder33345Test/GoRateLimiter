package RateLimiter

func commonGetDebug(l ILimiterGroup, t string) interface{} {
	var res []interface{}
	res = append(res, "TYPE: "+t)
	for _, limiter := range l.GetGroups() {
		db, ok := limiter.(Debuggable)
		if !ok {
			continue
		}
		res = append(res, db.GetDebug())
	}
	return res
}

func commonSetStorable(l ILimiterGroup, i interface{}) bool {
	td, ok := i.([]interface{})
	rsc := true
	if !ok {
		return false
	}

	dataI := 0
	for limI, liv := range l.GetGroups() {
		if st, ok := liv.(Storable); ok {
			ok := st.SetStorableData(td[limI])
			if !ok {
				rsc = false
			}
			dataI++
		}
	}
	return rsc
}
func commonGetStorable(l ILimiterGroup) interface{} {
	var res []interface{}
	for _, sl := range l.GetGroups() {
		st, ok := sl.(Storable)
		if !ok {
			continue
		}
		res = append(res, st.GetStorableData())
	}
	return res
}
