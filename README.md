Delivery fungsi menghandle http termasuk controller dto router

controller menerima request dari router
dto ada rquest dan response contract dari router
router memanggil controller
http.go config sebuah fiber

User -> delivery http(router -> controller -> dto) -> domain(usecase -> repositoru -> entity) -> infra(clean architecture)

domain ada entity repository usecase

entity sebuah struct
repository handling ke database
usecase sebuah bisnis logic dari aplikasi

infrasructure sebuah open koneksi antara lain postgre atau redis

shared utils general response
