// Copyright (c) 2013 Simon Zimmermann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package track provides an event collection HTTP server
// for an in-house analytics service.
// 
// Usage:
// 
//     track [flag]
// 
// The flags are:
// 
//     -v
//             verbose mode
//     -h
//             help text
//     -http=":8080"
//             set bind address for the HTTP server
//     -log=0
//             set log level
//     -dsn=""
//             MySQL data source name
//     -version=false
//             display version number and exit
//     -debug.cpuprofile=""
//             run cpu profiler
// 
// API
// 
// Track Session
// 
//     POST /api/1.0/track/session/
// 
// **Parameters**
// 
//     Region      
//         string  Required
// 
//     SessionID   
//         string  Required
// 
//     RemoteIP    
//         string  Required
// 
//     SessionType
//         string  Required
// 
//     Message
//         string  Optional
// 
//     ProfileID   
//         int     Optional
// 
// **Response**
// 
//     HTTP/1.1 201 Created
//     Date: Mon, 12 Aug 2013 09:37:17 GMT
//     Content-Length: 0
//     Content-Type: text/plain; charset=utf-8
// 
// Track User
// 
//     POST /api/1.0/track/user/
// 
// **Parameters**
// 
//     ProfileID   
//         int     Required
// 
//     Region      
//         string  Required
// 
//     Referrer    
//         string  Optional
// 
//     Message
//         string  Optional
// 
// **Response**
// 
//     HTTP/1.1 201 Created
//     Date: Mon, 12 Aug 2013 09:37:17 GMT
//     Content-Length: 0
//     Content-Type: text/plain; charset=utf-8
// 
// Track Item
// 
//     POST /api/1.0/track/item/
// 
// **Parameters**
// 
//     ProfileID   
//         int     Required
// 
//     Region      
//         string  Required
// 
//     ItemName    
//         string  Required
// 
//     ItemType    
//         string  Required
// 
//     IsUGC    
//         bool    Required
// 
//     PriceGold    
//         int     Optional
// 
//     PriceSilver    
//         int     Optional
// 
// **Response**
// 
//     HTTP/1.1 201 Created
//     Date: Mon, 12 Aug 2013 09:37:17 GMT
//     Content-Length: 0
//     Content-Type: text/plain; charset=utf-8
// 
// 
// Track Purchase
// 
//     POST /api/1.0/track/purchase/
// 
// **Parameters**
// 
//     ProfileID   
//         int     Required
// 
//     Region      
//         string  Required
// 
//     Currency    
//         string  Required
// 
//     GrossAmount    
//         int     Required
// 
//     NetAmount    
//         int     Required
// 
//     PaymentProvider    
//         string  Required
// 
//     Product    
//         string  Required
// 
// **Response**
// 
//     HTTP/1.1 201 Created
//     Date: Mon, 12 Aug 2013 09:37:17 GMT
//     Content-Length: 0
//     Content-Type: text/plain; charset=utf-8
// 
// TODO
// 
// - Correctly close down, cleanup and flush.
// - Check program for race conditions
// 
package track
