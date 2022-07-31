package greetings

import( 
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person

func Hello(name string) (string, error) {
    // If no name was given return an error message
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a greeting 
    // that embeds the name in the message
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets the initial value of variables used in a function
func init() {
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.

func randomFormat() string {
    formats := []string {
        "Namaste %v, padharo!",
        "Kem cho %v!",
        "Kem %v moj ma ne!",
    }
    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.

    return formats[rand.Intn(len(formats))]
}
