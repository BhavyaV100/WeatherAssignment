# WeatherAssignment
•	Got an API key from OpenWeatherMap by signing up on their website.
•	Used Go's HTTP package to make a request to the OpenWeatherMap API, providing the necessary parameters.
•	Parse the API response and format it into JSON.
•	In this file The http.Get function is used to make the API request, and ioutil.ReadAll reads the response body.
•	json.Unmarshal is used to parse the JSON response into the WeatherData struct.
• json.MarshalIndent is used to format the struct into a pretty-printed JSON string.

#Endpoints
-> Endpoint: /weather
-> Method: GET
-> Description: Returns the current weather data for Toronto.

#MainTest.go file 
-> A mock server is created using httptest.NewServer to simulate the OpenWeatherMap API. The server returns a predefined JSON response. 
-> The apiURL variable in the main code is overridden with the mock server URL. 
-> A request to the /weather endpoint is created using httptest.NewRequest, and a ResponseRecorder is used to record the response. 
-> The WeatherHandler function is called with the mock request and recorder. 
-> Assertions are made to check if the status code is HTTP 200 OK and to verify the response body. 
