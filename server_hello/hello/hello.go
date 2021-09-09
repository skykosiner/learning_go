package sayHello;

import (
    "net/http"
    "strings"
    //"fmt"
);

func sayHello(w http.ResponseWriter, r *http.Request) {
    message := r.URL.Path;
    message = strings.TrimPrefix(message, "/");
    message = "Hello " + message;

    w.Write([]byte(message));
};
