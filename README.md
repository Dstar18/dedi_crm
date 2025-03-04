# Customer Relationship Management

Aplikasi sederhana untuk melakukan rekap data calon customer (lead), customer, produk (layanan internet), dan penjualan.
Aplikasi ini dirancang menggunakan bahasa pemrograman Golang dengan framework echo dan database postgreSQL.

## Fitur  

- **Authentication Login**: Mengautentikasi pengguna menggunakan email dan password.  
- **Enkripsi Password dengan bcrypt**: Semua password pengguna di-enkripsi menggunakan metode bcrypt untuk keamanan.  
- **Error Handler**: Penanganan kesalahan yang efektif untuk setiap permintaan.  
- **Level Log**: Logging untuk informasi, peringatan, dan kesalahan dalam aplikasi.
- **Level User**: admin, manager, sales.  

## Endpoint API  

### Public Routes  

- **[POST] /register**  
  - Mendaftarkan pengguna baru.  
  - **Body**:  
    ```json  
    {
        "name": "Dedi",
        "email": "dedi@mail.com",
        "password": "Dedi12345#",
        "role": "manager"
    }  
    ```  

- **[POST] /login**  
  - Mengautentikasi pengguna.
  - **Body**:  
    ```json  
    {  
      "email": "dedi@mail.com",  
      "password": "Dedi12345#"  
    }  
    ```  

- **[GET] /logout**  
  - Mengeluarkan pengguna dari sesi aktif.  

### Protected Routes  

- **[GET] /api/customers**  
  - Menampilkan daftar semua pelanggan.  

- **[POST] /api/customers/add**  
  - Menambahkan pelanggan baru.  
  - **Body**:  
    ```json  
    {
        "name": "Budi",
        "email": "budi@mail.com",
        "phone": "012393023888",
        "address": "Jl. Raya"
    }
    ```  

- **[GET] /api/customers/lead**  
  - Menampilkan daftar customer (lead) / pengguna baru.  

- **[GET] /api/product**  
  - Menampilkan semua data produk.  

- **[POST] /api/product/add**  
  - Menambahkan produk baru.  
  - **Body**:  
    ```json  
    {
        "name": "Smarthome",
        "description": "Layanan Smart Home",
        "price": "2000000"
    }
    ```  

- **[POST] /api/product/update/{id}**  
  - Memperbarui produk berdasarkan ID.  
  - **Body**:  
    ```json  
    {
        "name": "Smarthome123",
        "description": "Layanan Smart Home",
        "price": "4000000"
    }
    ```  

- **[GET] /api/product/delete/{id}**  
  - Menghapus produk berdasarkan ID.  

- **[GET] /api/project/{id}**  
  - Mendapatkan detail proyek berdasarkan ID.  

- **[GET] /api/project/customer/{id}**  
  - Mendapatkan detail proyek berdasarkan ID pelanggan.  

- **[POST] /api/project/add** (Access: Sales)  
  - Menambahkan proyek baru. Hanya dapat diakses oleh pengguna dengan peran "sales".  
  - **Body**:  
    ```json  
    {
        "lead_id": 1,
        "product_id": "1,3"
    }
    ```  

- **[POST] /api/project/verifier/{id}** (Access: Manager)  
  - Memverifikasi proyek berdasarkan ID. Hanya dapat diakses oleh pengguna dengan peran "manager".  
  - **Body**:  
    ```json  
    {
        "status": "approved"
    }
    ```  

## Prasyarat  

Pastikan Anda memiliki:  

- [Docker](https://www.docker.com/get-started) terpasang di komputer Anda.  
- [Docker Compose](https://docs.docker.com/compose/install/) terinstal.  

## Cara Menjalankan Aplikasi  

1. **Clone repositori**:  

   ```bash  
   git clone https://github.com/Dstar18/dedi_crm.git
   cd dedi_crm  

2. **Jalankan docker compose**:  

   ```bash  
   docker-compose up --build    

3. **Access Aplikasi di Postman**:  

   ```bash  
   localhost:3000