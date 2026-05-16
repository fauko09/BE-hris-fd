# Dokumentasi API HRIS

Dokumen ini berisi daftar API HRIS berdasarkan route, handler, model, dan Swagger di project ini.

## Informasi Umum

- **Base URL lokal:** `http://localhost:8080/api/v1`
- **Format request body default:** `application/json`
- **Format response:** `application/json`
- **Static file upload:** `http://localhost:8080/uploads/...`

## Format Response

Response sukses:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

Response error:

```json
{
  "success": false,
  "error": "Pesan error"
}
```

## Authentication

Endpoint selain `/auth/register` dan `/auth/login` membutuhkan JWT token.

Header wajib untuk endpoint protected:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

Jika token tidak ada:

```json
{
  "success": false,
  "error": "Token tidak ditemukan"
}
```

Jika token tidak valid:

```json
{
  "success": false,
  "error": "Token tidak valid"
}
```

## Ringkasan Endpoint

| Method | Endpoint | Auth | Deskripsi |
|---|---|---:|---|
| `POST` | `/auth/register` | Tidak | Register user baru |
| `POST` | `/auth/login` | Tidak | Login user dan mendapatkan JWT token |
| `GET` | `/users` | Ya | Ambil semua user |
| `GET` | `/users/{id}` | Ya | Ambil user berdasarkan ID |
| `DELETE` | `/users/{id}` | Ya | Hapus user berdasarkan ID |
| `GET` | `/employees` | Ya | Ambil semua data detail karyawan |
| `POST` | `/employees/{uid}/data` | Ya | Simpan data detail karyawan |
| `GET` | `/employees/{uid}/data` | Ya | Ambil data detail karyawan berdasarkan UID |
| `POST` | `/employees/{uid}/photo` | Ya | Upload foto profil karyawan |
| `POST` | `/absensi/clock-in` | Ya | Clock in |
| `POST` | `/absensi/clock-out` | Ya | Clock out |
| `GET` | `/absensi/me` | Ya | Ambil riwayat absensi user login |
| `GET` | `/absensi/today` | Ya | Ambil semua absensi hari ini |
| `POST` | `/cuti` | Ya | Ajukan cuti |
| `GET` | `/cuti/me` | Ya | Ambil riwayat cuti user login |
| `GET` | `/cuti` | Ya | Ambil semua pengajuan cuti |
| `PUT` | `/cuti/{id}/approve` | Ya | Approve atau reject cuti |
| `POST` | `/penggajian` | Ya | Tambah data gaji karyawan |
| `GET` | `/penggajian/me` | Ya | Ambil riwayat gaji user login |
| `GET` | `/penggajian` | Ya | Ambil semua data penggajian |
| `PUT` | `/penggajian/{id}/bayar` | Ya | Update status gaji menjadi paid |
| `GET` | `/divisi` | Ya | Ambil semua divisi |
| `POST` | `/divisi` | Ya | Tambah divisi |
| `PUT` | `/divisi/{id}` | Ya | Update divisi |
| `DELETE` | `/divisi/{id}` | Ya | Hapus divisi |
| `GET` | `/jabatan` | Ya | Ambil semua jabatan |
| `POST` | `/jabatan` | Ya | Tambah jabatan |
| `PUT` | `/jabatan/{id}` | Ya | Update jabatan |
| `DELETE` | `/jabatan/{id}` | Ya | Hapus jabatan |

---

## Auth

### 1. Register User

**Endpoint:** `POST /auth/register`

**Deskripsi:** Membuat akun user baru.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `fullname` | string | Ya | Nama lengkap user |
| `nik` | string | Ya | NIK user |
| `email` | string | Ya | Email valid dan unik |
| `password` | string | Ya | Minimal 8 karakter |

**Contoh request:**

```json
{
  "fullname": "Budi Santoso",
  "nik": "1234567890123456",
  "email": "budi@example.com",
  "password": "password123"
}
```

**Response:**

- `201 Created`: Registrasi berhasil.
- `400 Bad Request`: Validasi gagal, email/NIK sudah terdaftar.

### 2. Login User

**Endpoint:** `POST /auth/login`

**Deskripsi:** Login dan mendapatkan JWT token.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `email` | string | Ya | Email user |
| `password` | string | Ya | Password user |

**Contoh request:**

```json
{
  "email": "budi@example.com",
  "password": "password123"
}
```

**Contoh response sukses:**

```json
{
  "success": true,
  "message": "Login berhasil",
  "data": {
    "token": "<jwt_token>",
    "user": {}
  }
}
```

**Response:**

- `200 OK`: Login berhasil.
- `400 Bad Request`: Validasi body gagal.
- `401 Unauthorized`: Email tidak ditemukan atau password salah.

---

## Users

Semua endpoint Users membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 3. Get Semua User

**Endpoint:** `GET /users`

**Deskripsi:** Ambil daftar semua user.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data user berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 4. Get User by ID

**Endpoint:** `GET /users/{id}`

**Deskripsi:** Ambil data user berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | User ID / UID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data user ditemukan.
- `400 Bad Request`: ID tidak valid.
- `404 Not Found`: User tidak ditemukan.

### 5. Delete User

**Endpoint:** `DELETE /users/{id}`

**Deskripsi:** Hapus user berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | User ID / UID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: User berhasil dihapus.
- `400 Bad Request`: ID tidak valid.
- `500 Internal Server Error`: Terjadi error server.

---

## Employees

Semua endpoint Employees membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 6. Get Semua Data Karyawan

**Endpoint:** `GET /employees`

**Deskripsi:** Ambil semua data detail karyawan.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data karyawan berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 7. Create Data User

**Endpoint:** `POST /employees/{uid}/data`

**Deskripsi:** Simpan data detail karyawan berdasarkan UID user.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `uid` | string | Ya | User ID / UID |

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `divisi_id` | string | Tidak | ID divisi |
| `jabatan_id` | string | Tidak | ID jabatan |
| `status` | string | Tidak | Status karyawan |
| `alamat_tinggal` | string | Tidak | Alamat tinggal |
| `tipe_tinggal` | string | Tidak | Tipe tempat tinggal |
| `domisili_tinggal` | string | Tidak | Domisili tinggal |
| `prov_tinggal` | string | Tidak | Provinsi tinggal |
| `alamat_ktp` | string | Tidak | Alamat sesuai KTP |
| `domisili_ktp` | string | Tidak | Domisili KTP |
| `prov_ktp` | string | Tidak | Provinsi KTP |
| `domisili_lahir` | string | Tidak | Domisili/tempat lahir |
| `tanggal_lahir` | string | Tidak | Format `YYYY-MM-DD` |
| `no_npwp` | string | Tidak | Nomor NPWP |
| `no_bpjs` | string | Tidak | Nomor BPJS |
| `is_bpjs_tk` | boolean | Tidak | Status BPJS TK |
| `no_asuransi` | string | Tidak | Nomor asuransi |
| `jenis_asuransi` | string | Tidak | Jenis asuransi |
| `nama_asuransi` | string | Tidak | Nama asuransi |
| `status_perkawinan` | string | Tidak | Status perkawinan |
| `nomor_telepon` | string | Tidak | Nomor telepon utama |
| `nomor_telepon_second` | string/null | Tidak | Nomor telepon kedua |
| `nomor_telepon_darurat` | string/null | Tidak | Nomor telepon darurat |
| `tanggal_masuk` | string | Tidak | Format `YYYY-MM-DD` |
| `no_rekening` | string | Tidak | Nomor rekening |
| `tipe_bank` | string | Tidak | Nama/tipe bank |

**Contoh request:**

```json
{
  "divisi_id": "uuid-divisi",
  "jabatan_id": "uuid-jabatan",
  "status": "active",
  "alamat_tinggal": "Jl. Melati No. 1",
  "tipe_tinggal": "kontrak",
  "domisili_tinggal": "Jakarta",
  "prov_tinggal": "DKI Jakarta",
  "alamat_ktp": "Jl. Melati No. 1",
  "domisili_ktp": "Jakarta",
  "prov_ktp": "DKI Jakarta",
  "domisili_lahir": "Jakarta",
  "tanggal_lahir": "1998-01-15",
  "no_npwp": "1234567890",
  "no_bpjs": "1234567890",
  "is_bpjs_tk": true,
  "no_asuransi": "ASR001",
  "jenis_asuransi": "kesehatan",
  "nama_asuransi": "Asuransi ABC",
  "status_perkawinan": "belum menikah",
  "nomor_telepon": "081234567890",
  "nomor_telepon_second": "081298765432",
  "nomor_telepon_darurat": "081211112222",
  "tanggal_masuk": "2026-05-01",
  "no_rekening": "123456789",
  "tipe_bank": "BCA"
}
```

**Response:**

- `201 Created`: Data user berhasil disimpan.
- `400 Bad Request`: UID/body tidak valid.
- `500 Internal Server Error`: Terjadi error server.

### 8. Get Data User by UID

**Endpoint:** `GET /employees/{uid}/data`

**Deskripsi:** Ambil data detail karyawan berdasarkan UID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `uid` | string | Ya | User ID / UID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data user ditemukan.
- `400 Bad Request`: UID tidak valid.
- `404 Not Found`: Data user tidak ditemukan.

### 9. Upload Foto Karyawan

**Endpoint:** `POST /employees/{uid}/photo`

**Deskripsi:** Upload foto profil karyawan berdasarkan UID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `multipart/form-data` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `uid` | string | Ya | User ID / UID |

**Query params:** Tidak ada.

**Form-data payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `photo` | file | Ya | File `.jpg`, `.jpeg`, atau `.png`, maksimal 2 MB |

**Response:**

- `200 OK`: Foto berhasil diupload.
- `400 Bad Request`: UID/file tidak valid, format salah, ukuran file lebih dari 2 MB.
- `500 Internal Server Error`: Gagal menyimpan file atau update user.

**Contoh response sukses:**

```json
{
  "success": true,
  "message": "Foto berhasil diupload",
  "data": {
    "photo_url": "/uploads/photos/<filename>"
  }
}
```

---

## Absensi

Semua endpoint Absensi membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 10. Clock In

**Endpoint:** `POST /absensi/clock-in`

**Deskripsi:** Karyawan melakukan clock in. Satu user hanya bisa clock in satu kali per hari.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Tidak |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `lokasi` | string | Tidak | Lokasi clock in |

**Contoh request:**

```json
{
  "lokasi": "Kantor Jakarta"
}
```

**Response:**

- `201 Created`: Clock in berhasil.
- `400 Bad Request`: Sudah clock in hari ini.
- `401 Unauthorized`: Token tidak valid.

### 11. Clock Out

**Endpoint:** `POST /absensi/clock-out`

**Deskripsi:** Karyawan melakukan clock out. User harus sudah clock in hari ini.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Tidak |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `lokasi` | string | Tidak | Lokasi clock out |

**Contoh request:**

```json
{
  "lokasi": "Kantor Jakarta"
}
```

**Response:**

- `200 OK`: Clock out berhasil.
- `400 Bad Request`: Belum clock in hari ini atau sudah clock out hari ini.
- `401 Unauthorized`: Token tidak valid.

### 12. Get Absensi Saya

**Endpoint:** `GET /absensi/me`

**Deskripsi:** Ambil riwayat absensi user yang sedang login.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `bulan` | integer | Tidak | Bulan `1-12` |
| `tahun` | integer | Tidak | Tahun, contoh `2026` |

**Body payload:** Tidak ada.

**Contoh:** `GET /absensi/me?bulan=5&tahun=2026`

**Response:**

- `200 OK`: Riwayat absensi berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 13. Get Absensi Hari Ini

**Endpoint:** `GET /absensi/today`

**Deskripsi:** Ambil semua absensi hari ini.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data absensi hari ini berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

---

## Cuti

Semua endpoint Cuti membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 14. Ajukan Cuti

**Endpoint:** `POST /cuti`

**Deskripsi:** Karyawan mengajukan permohonan cuti.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `tipe_cuti` | string | Ya | Tipe cuti |
| `tanggal_mulai` | string | Ya | Format `YYYY-MM-DD` |
| `tanggal_akhir` | string | Ya | Format `YYYY-MM-DD` |
| `alasan` | string | Ya | Alasan cuti |

**Contoh request:**

```json
{
  "tipe_cuti": "tahunan",
  "tanggal_mulai": "2026-05-20",
  "tanggal_akhir": "2026-05-22",
  "alasan": "Keperluan keluarga"
}
```

**Catatan:** `total_hari` dihitung otomatis dari `tanggal_mulai` sampai `tanggal_akhir`, termasuk hari awal dan akhir. Status awal adalah `pending`.

**Response:**

- `201 Created`: Pengajuan cuti berhasil.
- `400 Bad Request`: Validasi body gagal.
- `500 Internal Server Error`: Terjadi error server.

### 15. Get Cuti Saya

**Endpoint:** `GET /cuti/me`

**Deskripsi:** Ambil riwayat cuti user yang sedang login.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Riwayat cuti berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 16. Get Semua Cuti

**Endpoint:** `GET /cuti`

**Deskripsi:** Ambil semua pengajuan cuti.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `status` | string | Tidak | Filter status, contoh `pending`, `approved`, `rejected` |

**Body payload:** Tidak ada.

**Contoh:** `GET /cuti?status=pending`

**Response:**

- `200 OK`: Data cuti berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 17. Approve atau Reject Cuti

**Endpoint:** `PUT /cuti/{id}/approve`

**Deskripsi:** HR menyetujui atau menolak pengajuan cuti.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Cuti ID |

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `status` | string | Ya | Status baru, biasanya `approved` atau `rejected` |
| `catatan_hr` | string/null | Tidak | Catatan dari HR |

**Contoh request:**

```json
{
  "status": "approved",
  "catatan_hr": "Disetujui"
}
```

**Response:**

- `200 OK`: Status cuti berhasil diupdate.
- `400 Bad Request`: ID/body tidak valid.
- `500 Internal Server Error`: Terjadi error server.

---

## Penggajian

Semua endpoint Penggajian membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 18. Create Penggajian

**Endpoint:** `POST /penggajian`

**Deskripsi:** Tambah data gaji karyawan.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `uid` | string | Ya | User ID / UID karyawan |
| `periode` | string | Ya | Periode gaji, format contoh `2026-05` |
| `gaji_pokok` | number | Ya | Gaji pokok |
| `tunjangan` | number | Tidak | Tunjangan |
| `potongan` | number | Tidak | Potongan |
| `keterangan` | string/null | Tidak | Keterangan |

**Contoh request:**

```json
{
  "uid": "uuid-user",
  "periode": "2026-05",
  "gaji_pokok": 5000000,
  "tunjangan": 500000,
  "potongan": 100000,
  "keterangan": "Gaji bulan Mei"
}
```

**Catatan:** `potongan_bpjs` dihitung otomatis sebesar `1%` dari `gaji_pokok`. `total_gaji` dihitung otomatis dengan rumus `gaji_pokok + tunjangan - potongan - potongan_bpjs`. Status bayar awal mengikuti default model yaitu `unpaid`.

**Response:**

- `201 Created`: Data gaji berhasil disimpan.
- `400 Bad Request`: Validasi body gagal.
- `500 Internal Server Error`: Terjadi error server.

### 19. Get Gaji Saya

**Endpoint:** `GET /penggajian/me`

**Deskripsi:** Ambil riwayat gaji user yang sedang login.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Riwayat gaji berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 20. Get Semua Penggajian

**Endpoint:** `GET /penggajian`

**Deskripsi:** Ambil semua data penggajian.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `periode` | string | Tidak | Filter periode, contoh `2026-05` |

**Body payload:** Tidak ada.

**Contoh:** `GET /penggajian?periode=2026-05`

**Response:**

- `200 OK`: Data penggajian berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 21. Bayar Gaji

**Endpoint:** `PUT /penggajian/{id}/bayar`

**Deskripsi:** Update status gaji menjadi `paid` dan mengisi `tanggal_bayar`.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Gaji ID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Gaji berhasil dibayar.
- `400 Bad Request`: ID tidak valid.
- `500 Internal Server Error`: Terjadi error server.

---

## Divisi

Semua endpoint Divisi membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 22. Get Semua Divisi

**Endpoint:** `GET /divisi`

**Deskripsi:** Ambil daftar semua divisi.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data divisi berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 23. Create Divisi

**Endpoint:** `POST /divisi`

**Deskripsi:** Tambah divisi baru.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `name_divisi` | string | Ya | Nama divisi |

**Contoh request:**

```json
{
  "name_divisi": "Engineering"
}
```

**Response:**

- `201 Created`: Divisi berhasil dibuat.
- `400 Bad Request`: Validasi body gagal.
- `500 Internal Server Error`: Terjadi error server.

### 24. Update Divisi

**Endpoint:** `PUT /divisi/{id}`

**Deskripsi:** Update nama divisi berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Divisi ID |

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `name_divisi` | string | Ya | Nama divisi baru |

**Contoh request:**

```json
{
  "name_divisi": "Human Resource"
}
```

**Response:**

- `200 OK`: Divisi berhasil diupdate.
- `400 Bad Request`: ID/body tidak valid.
- `500 Internal Server Error`: Terjadi error server.

### 25. Delete Divisi

**Endpoint:** `DELETE /divisi/{id}`

**Deskripsi:** Hapus divisi berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Divisi ID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Divisi berhasil dihapus.
- `400 Bad Request`: ID tidak valid.
- `500 Internal Server Error`: Terjadi error server.

---

## Jabatan

Semua endpoint Jabatan membutuhkan header:

| Header | Value |
|---|---|
| `Authorization` | `Bearer <jwt_token>` |

### 26. Get Semua Jabatan

**Endpoint:** `GET /jabatan`

**Deskripsi:** Ambil daftar semua jabatan.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Data jabatan berhasil diambil.
- `500 Internal Server Error`: Terjadi error server.

### 27. Create Jabatan

**Endpoint:** `POST /jabatan`

**Deskripsi:** Tambah jabatan baru.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:** Tidak ada.

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `nama_jabatan` | string | Ya | Nama jabatan |

**Contoh request:**

```json
{
  "nama_jabatan": "Backend Developer"
}
```

**Response:**

- `201 Created`: Jabatan berhasil dibuat.
- `400 Bad Request`: Validasi body gagal.
- `500 Internal Server Error`: Terjadi error server.

### 28. Update Jabatan

**Endpoint:** `PUT /jabatan/{id}`

**Deskripsi:** Update nama jabatan berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |
| `Content-Type` | `application/json` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Jabatan ID |

**Query params:** Tidak ada.

**Body payload:**

| Field | Type | Wajib | Keterangan |
|---|---|---:|---|
| `nama_jabatan` | string | Ya | Nama jabatan baru |

**Contoh request:**

```json
{
  "nama_jabatan": "Senior Backend Developer"
}
```

**Response:**

- `200 OK`: Jabatan berhasil diupdate.
- `400 Bad Request`: ID/body tidak valid.
- `500 Internal Server Error`: Terjadi error server.

### 29. Delete Jabatan

**Endpoint:** `DELETE /jabatan/{id}`

**Deskripsi:** Hapus jabatan berdasarkan ID.

**Headers:**

| Header | Value | Wajib |
|---|---|---:|
| `Authorization` | `Bearer <jwt_token>` | Ya |

**Path params:**

| Param | Type | Wajib | Keterangan |
|---|---|---:|---|
| `id` | string | Ya | Jabatan ID |

**Query params:** Tidak ada.

**Body payload:** Tidak ada.

**Response:**

- `200 OK`: Jabatan berhasil dihapus.
- `400 Bad Request`: ID tidak valid.
- `500 Internal Server Error`: Terjadi error server.

---

## Referensi Model Response

Bagian ini menjelaskan field umum yang muncul di `data` response.

### User

| Field | Type | Keterangan |
|---|---|---|
| `uid` | string | ID user |
| `fullname` | string | Nama lengkap |
| `nik` | string | NIK |
| `email` | string | Email |
| `photo_url` | string/null | URL foto profil |
| `created_at` | string | Timestamp dibuat |
| `updated_at` | string | Timestamp diupdate |
| `data_user` | object | Detail data karyawan, jika dipreload |

### Data User

| Field | Type | Keterangan |
|---|---|---|
| `da_id` | string | ID data user |
| `uid` | string | ID user |
| `divisi_id` | string/null | ID divisi |
| `jabatan_id` | string/null | ID jabatan |
| `status` | string | Status karyawan |
| `tanggal_lahir` | string/null | Tanggal lahir |
| `tanggal_masuk` | string/null | Tanggal masuk |
| `divisi` | object | Data divisi, jika dipreload |
| `jabatan` | object | Data jabatan, jika dipreload |

### Absensi

| Field | Type | Keterangan |
|---|---|---|
| `absensi_id` | string | ID absensi |
| `uid` | string | ID user |
| `tanggal` | string | Tanggal absensi |
| `jam_masuk` | string/null | Waktu clock in |
| `jam_keluar` | string/null | Waktu clock out |
| `status` | string | Status absensi, default `hadir` |
| `lokasi_masuk` | string/null | Lokasi clock in |
| `lokasi_keluar` | string/null | Lokasi clock out |

### Cuti

| Field | Type | Keterangan |
|---|---|---|
| `cuti_id` | string | ID cuti |
| `uid` | string | ID user |
| `tipe_cuti` | string | Tipe cuti |
| `tanggal_mulai` | string | Tanggal mulai |
| `tanggal_akhir` | string | Tanggal akhir |
| `total_hari` | integer | Total hari cuti |
| `alasan` | string | Alasan cuti |
| `status` | string | Status cuti |
| `catatan_hr` | string/null | Catatan HR |
| `approved_by` | string/null | UID approver |
| `approved_at` | string/null | Waktu approval |
| `created_at` | string | Timestamp dibuat |

### Penggajian

| Field | Type | Keterangan |
|---|---|---|
| `gaji_id` | string | ID gaji |
| `uid` | string | ID user |
| `periode` | string | Periode gaji |
| `gaji_pokok` | number | Gaji pokok |
| `tunjangan` | number | Tunjangan |
| `potongan` | number | Potongan |
| `potongan_bpjs` | number | Potongan BPJS otomatis |
| `total_gaji` | number | Total gaji otomatis |
| `status_bayar` | string | Status pembayaran, default `unpaid` |
| `tanggal_bayar` | string/null | Waktu pembayaran |
| `keterangan` | string/null | Keterangan |
| `created_at` | string | Timestamp dibuat |

### Divisi

| Field | Type | Keterangan |
|---|---|---|
| `divisi_id` | string | ID divisi |
| `name_divisi` | string | Nama divisi |

### Jabatan

| Field | Type | Keterangan |
|---|---|---|
| `jabatan_id` | string | ID jabatan |
| `nama_jabatan` | string | Nama jabatan |
