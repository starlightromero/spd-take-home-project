# Weather API

A simple weather API implemented with Go

## API Routes

### Hello World

**Request** 

```
GET /
```

**Response**

```
"Hello, World!"
```

### Weather

**Request** 

```
GET /${zipCode}
```

**Response**
```
"coord": {
    "lon": float64,
    "lat": float64
},
"weather": [
    {
        "id": int,
        "main": string,
        "description": string,
        "icon": string
    },
    {
        ...
    }
],
"base": string,
"main": {
    "temp": float64,
	"feels_like": float64,
	"temp_min": float64,
	"temp_max": floast64,
	"pressure": int,
	"humidity": int,
},
"visibility": int,
"wind": {
    "speed": float64,
	"deg": int,
	"gust": float64
},
"clouds": {
    "all": int
},
"dt": int,
"sys": {
    "type": int,
	"id": int,
	"country": string,
	"sunrise": int,
	"sunset": int
}
"timezone": int,
"id": int,
"name": string,
"cod": int || string
```

### Mood

**Request**

```
GET /mood/${mood}
```

**Response**

```
"You are in a ${mood} mood. Thanks for sharing!"
```