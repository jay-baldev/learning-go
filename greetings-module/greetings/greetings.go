package greetings

import( 
    "fmt"
    "errors"
)

// Hello returns a greeting for the named person

func Hello(name string) (string, error) {
    // If no name was given return an error message
    if name == "" {
        return "", errors.New("empty name")
    }
    // If a name was received, return a greeting 
    // that embeds the name in the message
    message := fmt.Sprintf("Namaste %v, padhariye!", name)
    return message, nil
}
