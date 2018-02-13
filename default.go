package main

var DefaultConfig string = `
# [[conf]]
#     [conf.static]
#         text = "hello, world"

[[conf]]
    [conf.battery]

[[conf]]
    [conf.time]
        TimeFormat = "2006-01-02 15:04:05"
        TimeZone = "Local"
        Format = "{{.Time}}"
`
