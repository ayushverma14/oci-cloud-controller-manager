package protobuf

import (
	"time"

	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func ToUint32Value(val uint32) wrappers.UInt32Value {
	return wrappers.UInt32Value{Value: val}
}

// Specifically does not return a default uint32 value of 0
func ToUint32(val *wrappers.UInt32Value) *uint32 {
	if val != nil {
		return &val.Value
	}
	return nil
}

func ToFloatValue(val float32) wrappers.FloatValue {
	return wrappers.FloatValue{Value: val}
}

func ToFloat32(val *wrappers.FloatValue) *float32 {
	if val != nil {
		return &val.Value
	}
	return nil
}

func FromTime(src time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: src.Unix(),
		Nanos:   int32(src.UnixNano()),
	}
}

func ToTime(t *timestamp.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return time.Unix(t.Seconds, int64(t.Nanos))
}

func ToDuration(m *duration.Duration) time.Duration {
	if m == nil {
		return 0
	}
	return time.Duration(m.Seconds)*time.Second + time.Duration(m.Nanos)
}

func GetZeroTimestamp() *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: 0,
		Nanos:   0,
	}
}

func IsTimestampZero(t *timestamp.Timestamp) bool {
	if t == nil {
		return true
	}

	if t.Seconds == 0 && t.Nanos == 0 {
		return true
	}

	return false
}
