package greetings

import(
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// multi-input hello mapping name to greeting
// hello kept for backwards compatibility
func Hellos(names []string) (map[string]string, error) {
    // A map which associates names with messages
    messages := make(map[string]string) // initialize

    // seems to be a for loop, going through the slice (not array!)
    // so for each _ and name in []string (where[] is nil and names is the string)
    // assigning either one to what's in names (range seems to be an interator)
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }

    return messages, nil
}

func randomFormat() string {
    formats:= []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}