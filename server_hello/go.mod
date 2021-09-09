module server.com/server

go 1.16

replace example.com/hello_func => ./hello

require example.com/hello_func v0.0.0-00010101000000-000000000000 // indirect
