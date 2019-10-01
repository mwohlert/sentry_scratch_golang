package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

const flushDuration time.Duration = 2 * time.Second
const sentryDSN string = "https://62fa8ecced0f4e97a96049c056b5d611@sentry.io/1766414"

func main() {
	sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})

	eventID := sentry.CaptureException(errors.New("ExampleError"))
	fmt.Printf("Sentry Event ID %v\n", eventID)
	sentry.Flush(flushDuration)

	fmt.Println("Error should be reported")

	time.Sleep(flushDuration)
	os.Exit(0)
}
