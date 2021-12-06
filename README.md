# yomo flow for noise-example

This is part of the [example-noise](https://github.com/yomorun/example-noise), which describes how to write a [**noise-flow**](https://github.com/yomorun/yomo-flow-noise-example) to process data from source or other flow apps.

![arch1.png](https://github.com/yomorun/example-noise/raw/main/docs/arch1.png?raw=true)

## 🚀 Getting Started

### 1. Install CLI

> **Note:** YoMo requires Go 1.15 and above, run `go version` to get the version of Go in your environment, please follow [this link](https://golang.org/doc/install) to install or upgrade if it doesn't fit the requirement.

```bash
go install github.com/yomorun/cli/yomo@v0.1.3
```

### 2. Create your serverless app

```
$ mkdir -p $GOPATH/src/github.com/{YOUR_GITHUB_USERNAME} && cd $_

$ yomo init yomo-flow-noise-example
2021/04/28 15:20:57 Initializing the Serverless app...
2021/04/28 15:21:01 🛠 go.mod replaced
2021/04/28 15:21:01 ✅ Congratulations! You have initialized the serverless app successfully.
2021/04/28 15:21:01 🎉 You can enjoy the YoMo Serverless via the command: yomo dev

$ cd yomo-flow-noise-example
```

Update the `app.go`

```go
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
		Map(computePeek)

	return stream
}

func DataID() []byte {
	return []byte{0x33}
}

```

### 3. Run your serverless app

```go
yomo run app.go -m go.mod -u localhost:9999 -n NoiseServerless
```

### Container

#### Docker Image

The case provides [Dockefile](https://github.com/yomorun/yomo-flow-noise-example/blob/main/Dockerfile) files for packaging into images.

Also, you can get the official packaged image ([noise-flow](https://hub.docker.com/r/yomorun/noise-flow)) from the mirror repository.

```
docker pull yomorun/noise-flow
```

#### Docker run

You can run the service with the following command:

```
docker run --rm --name noise-flow -p 4242:4242 yomorun/noise-flow:0.0.6
```

