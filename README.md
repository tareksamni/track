# track — event collection HTTP service

[![Build Status](https://travis-ci.org/simonz05/track.png?branch=master)](https://travis-ci.org/simonz05/track)

`track` is an HTTP server used to collect events for an
in-house analytics service.

Usage:

    track [flag]

The flags are:

    -v
            verbose mode
    -h
            help text
    -http=":8080"
            set bind address for the HTTP server
    -log=0
            set log level
    -dsn=""
            MySQL data source name
    -version=false
            display version number and exit
    -debug.cpuprofile=""
            run cpu profiler

## API

### Track Session

    POST /api/1.0/track/session/

**Parameters**

    Region      
        string  Required

    SessionID   
        string  Required

    RemoteIP    
        string  Required

    SessionType
        string  Required

    Message
        string  Optional

    ProfileID   
        int     Optional

**Response**

    HTTP/1.1 201 Created
    Date: Mon, 12 Aug 2013 09:37:17 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8

### Track User

    POST /api/1.0/track/user/

**Parameters**

    ProfileID   
        int     Required

    Region      
        string  Required

    Referrer    
        string  Optional

    Message
        string  Optional

**Response**

    HTTP/1.1 201 Created
    Date: Mon, 12 Aug 2013 09:37:17 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8

### Track Item

    POST /api/1.0/track/item/

**Parameters**

    ProfileID   
        int     Required

    Region      
        string  Required

    ItemName    
        string  Required

    ItemType    
        string  Required

    IsUGC    
        bool    Required

    PriceGold    
        int     Optional

    PriceSilver    
        int     Optional

**Response**

    HTTP/1.1 201 Created
    Date: Mon, 12 Aug 2013 09:37:17 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8


### Track Purchase

    POST /api/1.0/track/purchase/

**Parameters**

    ProfileID   
        int     Required

    Region      
        string  Required

    Currency    
        string  Required

    GrossAmount    
        int     Required

    NetAmount    
        int     Required

    PaymentProvider    
        string  Required

    Product    
        string  Required

**Response**

    HTTP/1.1 201 Created
    Date: Mon, 12 Aug 2013 09:37:17 GMT
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8
