// this file was generated by gomacro command: import "time"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	pkg "time"
)

func Package_time() (map[string]Value, map[string]Type) {
	return map[string]Value{
			"ANSIC":           ValueOf(pkg.ANSIC),
			"After":           ValueOf(pkg.After),
			"AfterFunc":       ValueOf(pkg.AfterFunc),
			"April":           ValueOf(pkg.April),
			"August":          ValueOf(pkg.August),
			"Date":            ValueOf(pkg.Date),
			"December":        ValueOf(pkg.December),
			"February":        ValueOf(pkg.February),
			"FixedZone":       ValueOf(pkg.FixedZone),
			"Friday":          ValueOf(pkg.Friday),
			"Hour":            ValueOf(pkg.Hour),
			"January":         ValueOf(pkg.January),
			"July":            ValueOf(pkg.July),
			"June":            ValueOf(pkg.June),
			"Kitchen":         ValueOf(pkg.Kitchen),
			"LoadLocation":    ValueOf(pkg.LoadLocation),
			"Local":           ValueOf(&pkg.Local).Elem(),
			"March":           ValueOf(pkg.March),
			"May":             ValueOf(pkg.May),
			"Microsecond":     ValueOf(pkg.Microsecond),
			"Millisecond":     ValueOf(pkg.Millisecond),
			"Minute":          ValueOf(pkg.Minute),
			"Monday":          ValueOf(pkg.Monday),
			"Nanosecond":      ValueOf(pkg.Nanosecond),
			"NewTicker":       ValueOf(pkg.NewTicker),
			"NewTimer":        ValueOf(pkg.NewTimer),
			"November":        ValueOf(pkg.November),
			"Now":             ValueOf(pkg.Now),
			"October":         ValueOf(pkg.October),
			"Parse":           ValueOf(pkg.Parse),
			"ParseDuration":   ValueOf(pkg.ParseDuration),
			"ParseInLocation": ValueOf(pkg.ParseInLocation),
			"RFC1123":         ValueOf(pkg.RFC1123),
			"RFC1123Z":        ValueOf(pkg.RFC1123Z),
			"RFC3339":         ValueOf(pkg.RFC3339),
			"RFC3339Nano":     ValueOf(pkg.RFC3339Nano),
			"RFC822":          ValueOf(pkg.RFC822),
			"RFC822Z":         ValueOf(pkg.RFC822Z),
			"RFC850":          ValueOf(pkg.RFC850),
			"RubyDate":        ValueOf(pkg.RubyDate),
			"Saturday":        ValueOf(pkg.Saturday),
			"Second":          ValueOf(pkg.Second),
			"September":       ValueOf(pkg.September),
			"Since":           ValueOf(pkg.Since),
			"Sleep":           ValueOf(pkg.Sleep),
			"Stamp":           ValueOf(pkg.Stamp),
			"StampMicro":      ValueOf(pkg.StampMicro),
			"StampMilli":      ValueOf(pkg.StampMilli),
			"StampNano":       ValueOf(pkg.StampNano),
			"Sunday":          ValueOf(pkg.Sunday),
			"Thursday":        ValueOf(pkg.Thursday),
			"Tick":            ValueOf(pkg.Tick),
			"Tuesday":         ValueOf(pkg.Tuesday),
			"UTC":             ValueOf(&pkg.UTC).Elem(),
			"Unix":            ValueOf(pkg.Unix),
			"UnixDate":        ValueOf(pkg.UnixDate),
			"Until":           ValueOf(pkg.Until),
			"Wednesday":       ValueOf(pkg.Wednesday),
		}, map[string]Type{
			"Duration":   TypeOf((*pkg.Duration)(nil)).Elem(),
			"Location":   TypeOf((*pkg.Location)(nil)).Elem(),
			"Month":      TypeOf((*pkg.Month)(nil)).Elem(),
			"ParseError": TypeOf((*pkg.ParseError)(nil)).Elem(),
			"Ticker":     TypeOf((*pkg.Ticker)(nil)).Elem(),
			"Time":       TypeOf((*pkg.Time)(nil)).Elem(),
			"Timer":      TypeOf((*pkg.Timer)(nil)).Elem(),
			"Weekday":    TypeOf((*pkg.Weekday)(nil)).Elem(),
		}
}

func init() {
	binds, types := Package_time()
	Binds["time"] = binds
	Types["time"] = types
}
