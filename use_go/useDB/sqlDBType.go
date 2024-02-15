// sql.DB 객체 db 생성
db, err := sql.Open("mysql", "root:xman0922@tcp(127.0.0.1:3306)/testdb")

// db 차후에 닫기
defer db.Close()

// SELECT 쿼리
rows, err := db.Query("SELECT id, name FROM test")

// INSERT 실행
db.Exec("INSERT INTO test(id, name) VALUES (1, 'Alex')")

