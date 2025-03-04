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

- **[POST] /login**  
  - Mengautentikasi pengguna.
  - Default akun: email(admin@mail.com), password(Admin123#)  
  - **Body**:  
    ```json  
    {  
      "email": "dedi@mail.com",  
      "password": "Dedi12345#"  
    }  
    ```  
    
- **[POST] /register**  
  - Mendaftarkan pengguna baru.  

- **[GET] /logout**  
  - Mengeluarkan pengguna dari sesi aktif.  

### Protected Routes  

- **[GET] /customers**  
  - Mendapatkan daftar semua pelanggan.  

- **[POST] /customers/add**  
  - Menambahkan pelanggan baru.  

- **[GET] /customers/lead**  
  - Mendapatkan daftar lead pelanggan.  

- **[GET] /product**  
  - Mengambil daftar produk.  

- **[POST] /product/add**  
  - Menambahkan produk baru.  

- **[POST] /product/update/{id}**  
  - Memperbarui produk berdasarkan ID.  

- **[GET] /product/delete/{id}**  
  - Menghapus produk berdasarkan ID.  

- **[GET] /project/{id}**  
  - Mendapatkan detail proyek berdasarkan ID.  

- **[GET] /project/customer/{id}**  
  - Mendapatkan detail proyek berdasarkan ID pelanggan.  

- **[POST] /project/add** (Access: Sales)  
  - Menambahkan proyek baru. Hanya dapat diakses oleh pengguna dengan peran "sales".  

- **[POST] /project/verifier/{id}** (Access: Manager)  
  - Memverifikasi proyek berdasarkan ID. Hanya dapat diakses oleh pengguna dengan peran "manager".  

## Prasyarat  

Pastikan Anda memiliki:  

- [Docker](https://www.docker.com/get-started) terpasang di komputer Anda.  
- [Docker Compose](https://docs.docker.com/compose/install/) terinstal.  

## Cara Menjalankan Aplikasi  

1. **Klon repositori ini**:  

   ```bash  
   git clone https://github.com/username/your-app.git  
   cd your-app  