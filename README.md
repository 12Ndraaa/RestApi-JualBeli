# restapi-jualbeli

REST API sistem jual beli menggunakan Golang dan PostgreSQL. Dibuat tanpa framework — hanya menggunakan package standar `net/http` dan driver `lib/pq`.

---

## Cara Menjalankan

**1. Clone repository**

```bash
git clone https://github.com/12Ndraaa/restapi-jualbeli.git
cd restapi-jualbeli
```

**2. Setup database**

Buat database baru di PostgreSQL, kemudian jalankan script yang tersedia:

```
IMPORT_THIS/DATABASE_JUAL_BELI.sql
```

**3. Konfigurasi environment**

Salin file contoh dan isi dengan kredensial database:

```bash
cp .env.example .env
```

Isi file `.env`:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password_anda
DB_NAME=nama_database_anda
SERVER_PORT=9000
```

**4. Install dependencies**

```bash
go mod tidy
```

**5. Jalankan server**

```bash
go run .
```

Server berjalan di `http://localhost:9000`.

---

## Struktur Project

```
restapi-jualbeli/
├── config/          # Koneksi database
├── model/           # Definisi struct
├── repository/      # Query SQL
├── service/         # Business logic
├── handler/         # HTTP handler
├── IMPORT_THIS/     # File SQL dan Postman collection
├── main.go
└── .env.example
```

Setiap fitur mengikuti pola yang sama: `model` -> `repository` -> `service` -> `handler`.

---

## Endpoint

**Master Data**

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | /produk | Ambil semua produk |
| GET | /produk/:id | Ambil produk berdasarkan ID |
| POST | /produk | Tambah produk baru |
| PUT | /produk/:id | Update produk |
| DELETE | /produk/:id | Hapus produk |
| GET | /mitra | Ambil semua mitra bisnis |
| GET | /mitra/:id | Ambil mitra bisnis berdasarkan ID |
| POST | /mitra | Tambah mitra bisnis baru |
| PUT | /mitra/:id | Update mitra bisnis |
| DELETE | /mitra/:id | Hapus mitra bisnis |
| GET | /gudang | Ambil semua gudang |
| GET | /gudang/:id | Ambil gudang berdasarkan ID |
| POST | /gudang | Tambah gudang baru |
| PUT | /gudang/:id | Update gudang |
| DELETE | /gudang/:id | Hapus gudang |

**Transaksi**

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | /pembelian | Ambil semua transaksi pembelian |
| GET | /pembelian/:id | Ambil transaksi pembelian berdasarkan ID |
| POST | /pembelian | Buat transaksi pembelian, stok otomatis bertambah |
| GET | /penjualan | Ambil semua transaksi penjualan |
| GET | /penjualan/:id | Ambil transaksi penjualan berdasarkan ID |
| POST | /penjualan | Buat transaksi penjualan, stok otomatis berkurang |

**Laporan**

| Method | Endpoint | Keterangan |
|--------|----------|------------|
| GET | /laporan/faktur | Rekap penjualan per faktur |
| GET | /laporan/item | Rekap penjualan per item |
| GET | /laporan/stok | Lihat stok per gudang |

---

## Contoh Request

**POST /pembelian**

```json
{
    "header": {
        "trxno": "PB-001",
        "bp_id": 1,
        "tgl": "2026-03-09T00:00:00Z",
        "diskon": 0
    },
    "details": [
        {
            "gudang_id": 1,
            "item_id": 1,
            "qty": 10,
            "harga": 5000,
            "diskon": 0
        }
    ]
}
```

**POST /penjualan**

```json
{
    "header": {
        "trxno": "PJ-001",
        "bp_id": 1,
        "tgl": "2026-03-09T00:00:00Z",
        "diskon": 0
    },
    "details": [
        {
            "gudang_id": 1,
            "item_id": 1,
            "qty": 5,
            "harga": 7000,
            "diskon": 0
        }
    ]
}
```

---

## Test API

Import file berikut ke Postman untuk mencoba semua endpoint:

```
IMPORT_THIS/postman_collection.json
```
