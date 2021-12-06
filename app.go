package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yomorun/yomo/rx"
)

// ThresholdSingleValue is the threshold of a single value.
const ThresholdSingleValue = 16

// NoiseData represents the structure of data
type NoiseData struct {
	Noise float32 `json:"noise"` // Noise value
	Time  int64   `json:"time"` // Timestamp (ms)
	From  string  `json:"from"` // Source IP
}

// Print every value and alert for value greater than ThresholdSingleValue
var computePeek = func(_ context.Context, i interface{}) (interface{}, error) {
	value := i.(*NoiseData)
	value.Noise = value.Noise / 10
	rightNow := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(fmt.Sprintf("[%s] %d > value: %f ⚡️=%dms", value.From, value.Time, value.Noise, rightNow-value.Time))

	// Compute peek value, if greater than ThresholdSingleValue, alert
	if value.Noise >= ThresholdSingleValue {
		fmt.Println(fmt.Sprintf("❗ value: %f reaches the threshold %d! 𝚫=%f", value.Noise, ThresholdSingleValue, value.Noise-ThresholdSingleValue))
	}

	return value, nil
}

// Handler will handle data in Rx way
func Handler(rxstream rx.Stream) rx.Stream {
	stream := rxstream.
		Unmarshal(json.Unmarshal, func() interface{} { return &NoiseData{} }).
		Map(computePeek).
		Marshal(json.Marshal).
		PipeBackToZipper(0x34)

	return stream
}

func DataID() []byte {
	return []byte{0x33}
}
