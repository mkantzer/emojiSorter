////////////////////////////////////////

// // Test types can help frame the test
// // If your request/response types are hidden inside the handler, you can just declare new types in your test code.

// // This is an opportunity to do a bit of storytelling to future generations who will need to understand your code.

// // For example, let’s say we have a Person type in our code, and we reuse it on many endpoints. If we had a /greet endpoint, we might only care about their name, so we can express this in test code:

// func TestGreet(t *testing.T) {
//     is := is.New(t)
//     p := struct {
//         Name string `json:"name"`
//     }{
//         Name: "Mat Ryer",
//     }
//     var buf bytes.Buffer
//     err := json.NewEncoder(&buf).Encode(p)
//     is.NoErr(err) // json.NewEncoder
//     req, err := http.NewRequest(http.MethodPost, "/greet", &buf)
//     is.NoErr(err)
// 	//... more test code here
// // It’s clear from this test, that the only field we care about is the Name of the person.

////////////////////////////////////////

// // The server is testable
// // Our server type is very testable.

// func TestHandleAbout(t *testing.T) {
//     is := is.New(t)
//     srv := server{
//         db:    mockDatabase,
//         email: mockEmailSender,
//     }
//     srv.routes()
//     req := httptest.NewRequest("GET", "/about", nil)
//     w := httptest.NewRecorder()
//     srv.ServeHTTP(w, req)
//     is.Equal(w.StatusCode, http.StatusOK)
// }

// // Create a server instance inside each test — if expensive things lazy load, this won’t take much time at all, even for big components
// // By calling ServeHTTP on the server, we are testing the entire stack including routing and middleware, etc. You can of course call the handler methods directly if you want to avoid this
// // Use httptest.NewRequest and httptest.NewRecorder to record what the handlers are doing
// // This code sample uses is testing mini-framework (a mini alternative to Testify) github.com/matryer/is
