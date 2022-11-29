# go-unit-test
Unit Testing

- go test -v กรณีอยู่ในโฟเดอที่จะเทส
- go test {module}/package -v [-run={ชื่อ func ที่จะเทส}]/[-run="{ชื่อ func ที่จะเทส}/subtest"]
- go test ./... รันทั้งหมด
- go test {....} -cover => option check ว่าตรวจสอบไปกี่ % 
- go test github.com/pakawatkung/go-unit-test/service -bench=. -benchmem => เป็นการเช็ค performance
- go test -v -tags={=tag_name} => run integration test + unit Test
- godoc -http=:{port}


lib
- golang.org/x/tools/cmd/godoc