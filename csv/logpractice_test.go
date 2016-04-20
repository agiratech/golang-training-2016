package main

import(  
    "net/http"
    "net/http/httptest"
    "testing"
)




func Test_Log(t *testing.T) {  
    url := `/empid?empid=309&name=REDDY&age=22&address=tpt&salary=12000`
    homeHandle := Log(url)
    req, _ := http.NewRequest("GET", "", nil)
    w := httptest.NewRecorder()
    homeHandle.ServeHTTP(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("Home page didn't return %v", http.StatusOK)
    }
}
