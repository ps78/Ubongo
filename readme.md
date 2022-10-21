run coverage tests as follows:

go test ./... -coverprofile=coverage
go tool cover -html=coverage


create struct-dependency graph:

embedded-struct-visualizer -out dependencies.dot

then visualize it in VS-Code by selecting the file and pressing Ctrl+Shift+V