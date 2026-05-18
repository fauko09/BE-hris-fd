# HRIS Frontend Integration Brief

Dokumen ini adalah ringkasan backend untuk kebutuhan pembuatan frontend HRIS. Detail endpoint lengkap tetap ada di `API_DOCUMENTATION.md`; file ini fokus ke cara konsumsi API, bentuk data penting, flow halaman, dan catatan yang perlu diketahui saat vibe coding FE.

## Stack Backend

- Bahasa/framework: Go + Gin.
- Database: MySQL via GORM.
- Auth: JWT HS256, token berisi claim `uid` dan `exp`.
- Token login berlaku 24 jam.
- Base URL lokal: `http://localhost:8080/api/v1`.
- Swagger lokal: `http://localhost:8080/swagger/index.html`.
- File statis upload: `http://localhost:8080/uploads/...`.

## Environment Backend

Backend membaca `.env` dari root project saat `go run main.go`.

```env
APP_PORT=8080
DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
JWT_SECRET=
```

Perintah umum:

```bash
go mod download
go run cmd/migrate/main.go
go run main.go
```

Catatan FE: belum terlihat konfigurasi CORS di backend. Jika frontend berjalan di origin berbeda, misalnya `http://localhost:3000`, browser bisa terkena CORS. Solusi sementara untuk development adalah pakai proxy dari dev server FE ke backend, atau tambahkan middleware CORS di backend.

## Format Response

Semua response memakai wrapper berikut.

Sukses:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

Error:

```json
{
  "success": false,
  "error": "Pesan error"
}
```

Implikasi FE:

- Selalu baca `success`.
- Untuk sukses, data utama ada di `data`.
- Untuk gagal, tampilkan `error`.
- Jangan mengandalkan `message` untuk error karena error memakai field `error`.
- Untuk endpoint delete/update tertentu, `data` bisa `null` atau tidak ada.

## Auth FE

Endpoint publik:

- `POST /auth/register`
- `POST /auth/login`

Endpoint lain butuh header:

```http
Authorization: Bearer <jwt_token>
```

Login sukses mengembalikan:

```json
{
  "success": true,
  "message": "Login berhasil",
  "data": {
    "token": "<jwt_token>",
    "user": {
      "uid": "uuid",
      "fullname": "Nama User",
      "nik": "1234567890123456",
      "email": "user@example.com",
      "photo_url": null,
      "created_at": "2026-05-18T10:00:00+07:00",
      "updated_at": "2026-05-18T10:00:00+07:00"
    }
  }
}
```

Rekomendasi FE:

- Simpan `token` dan `user` setelah login.
- Pakai interceptor/fetch wrapper untuk menyisipkan header Authorization.
- Jika status HTTP `401`, hapus session lokal dan arahkan ke login.
- Backend saat ini tidak punya endpoint `/me`; untuk profil user login gunakan `user` dari response login, atau panggil `GET /users/{uid}` dengan UID dari session.
- JWT hanya membawa `uid`. Belum ada role/permission di backend, jadi UI role admin/employee belum bisa divalidasi server-side.

## Endpoint Map

Gunakan semua path di bawah dengan prefix `/api/v1`.

| Area | Method | Path | FE Use Case |
|---|---|---|---|
| Auth | `POST` | `/auth/register` | Form register user |
| Auth | `POST` | `/auth/login` | Form login |
| Users | `GET` | `/users` | List akun user |
| Users | `GET` | `/users/{id}` | Detail akun/user profile |
| Users | `DELETE` | `/users/{id}` | Hapus user |
| Employees | `GET` | `/employees` | List karyawan lengkap |
| Employees | `POST` | `/employees/{uid}/data` | Isi detail karyawan |
| Employees | `GET` | `/employees/{uid}/data` | Detail data karyawan |
| Employees | `POST` | `/employees/{uid}/photo` | Upload foto karyawan |
| Absensi | `POST` | `/absensi/clock-in` | Tombol clock in |
| Absensi | `POST` | `/absensi/clock-out` | Tombol clock out |
| Absensi | `GET` | `/absensi/me` | Riwayat absensi user login |
| Absensi | `GET` | `/absensi/today` | Dashboard absensi hari ini |
| Cuti | `POST` | `/cuti` | Form pengajuan cuti |
| Cuti | `GET` | `/cuti/me` | Riwayat cuti user login |
| Cuti | `GET` | `/cuti` | List approval cuti |
| Cuti | `PUT` | `/cuti/{id}/approve` | Approve/reject cuti |
| Penggajian | `POST` | `/penggajian` | Input payroll |
| Penggajian | `GET` | `/penggajian/me` | Slip/riwayat gaji user login |
| Penggajian | `GET` | `/penggajian` | List payroll |
| Penggajian | `PUT` | `/penggajian/{id}/bayar` | Mark paid |
| Divisi | `GET` | `/divisi` | Master divisi |
| Divisi | `POST` | `/divisi` | Tambah divisi |
| Divisi | `PUT` | `/divisi/{id}` | Edit divisi |
| Divisi | `DELETE` | `/divisi/{id}` | Hapus divisi |
| Jabatan | `GET` | `/jabatan` | Master jabatan |
| Jabatan | `POST` | `/jabatan` | Tambah jabatan |
| Jabatan | `PUT` | `/jabatan/{id}` | Edit jabatan |
| Jabatan | `DELETE` | `/jabatan/{id}` | Hapus jabatan |

## Data Shape Penting

### User

```ts
type User = {
  uid: string;
  fullname: string;
  nik: string;
  email: string;
  photo_url: string | null;
  created_at: string;
  updated_at: string;
  data_user?: DataUser | null;
};
```

### DataUser / Employee Detail

```ts
type DataUser = {
  da_id: string;
  uid: string;
  divisi_id: string | null;
  jabatan_id: string | null;
  status: string;
  alamat_tinggal: string;
  tipe_tinggal: string;
  domisili_tinggal: string;
  prov_tinggal: string;
  alamat_ktp: string;
  domisili_ktp: string;
  prov_ktp: string;
  domisili_lahir: string;
  tanggal_lahir: string | null;
  no_npwp: string;
  no_bpjs: string;
  is_bpjs_tk: boolean;
  no_asuransi: string;
  jenis_asuransi: string;
  nama_asuransi: string;
  status_perkawinan: string;
  id_family: string | null;
  nomor_telepon: string;
  nomor_telepon_second: string | null;
  nomor_telepon_darurat: string | null;
  tanggal_masuk: string | null;
  no_rekening: string;
  tipe_bank: string;
  user?: User;
  divisi?: Divisi;
  jabatan?: Jabatan;
};
```

Untuk `POST /employees/{uid}/data`, tanggal dikirim sebagai `YYYY-MM-DD`. Jika `divisi_id` atau `jabatan_id` kosong string, backend menyimpan `null`.

### Absensi

```ts
type Absensi = {
  absensi_id: string;
  uid: string;
  tanggal: string;
  jam_masuk: string | null;
  jam_keluar: string | null;
  status: "hadir" | string;
  keterangan: string | null;
  lokasi_masuk: string | null;
  lokasi_keluar: string | null;
  created_at: string;
  user?: User;
};
```

Rules:

- Clock in hanya bisa sekali per user per tanggal.
- Clock out butuh data clock in hari yang sama.
- Request body clock in/out opsional:

```json
{
  "lokasi": "Kantor Jakarta"
}
```

### Cuti

```ts
type Cuti = {
  cuti_id: string;
  uid: string;
  tipe_cuti: string;
  tanggal_mulai: string;
  tanggal_akhir: string;
  total_hari: number;
  alasan: string;
  status: "pending" | "approved" | "rejected" | string;
  catatan_hr: string | null;
  approved_by: string | null;
  approved_at: string | null;
  created_at: string;
  user?: User;
};
```

Request pengajuan cuti:

```json
{
  "tipe_cuti": "tahunan",
  "tanggal_mulai": "2026-05-20",
  "tanggal_akhir": "2026-05-22",
  "alasan": "Keperluan keluarga"
}
```

Request approve/reject:

```json
{
  "status": "approved",
  "catatan_hr": "Disetujui"
}
```

Catatan: backend belum membatasi nilai `status`, jadi FE sebaiknya hanya menyediakan pilihan `approved` dan `rejected`.

### Penggajian

```ts
type Penggajian = {
  gaji_id: string;
  uid: string;
  periode: string;
  gaji_pokok: number;
  tunjangan: number;
  potongan: number;
  potongan_bpjs: number;
  total_gaji: number;
  status_bayar: "unpaid" | "paid" | string;
  tanggal_bayar: string | null;
  keterangan: string | null;
  created_at: string;
  user?: User;
};
```

Request create payroll:

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

Rules:

- `periode` format praktis: `YYYY-MM`.
- `potongan_bpjs` dihitung backend sebesar `1%` dari `gaji_pokok`.
- `total_gaji` dihitung backend: `gaji_pokok + tunjangan - potongan - potongan_bpjs`.
- `PUT /penggajian/{id}/bayar` tidak butuh body.

### Divisi dan Jabatan

```ts
type Divisi = {
  divisi_id: string;
  name_divisi: string;
};

type Jabatan = {
  jabatan_id: string;
  nama_jabatan: string;
};
```

## Query Parameters

- `GET /absensi/me?bulan=5&tahun=2026`
- `GET /cuti?status=pending`
- `GET /penggajian?periode=2026-05`

Jika `bulan` atau `tahun` kosong pada `/absensi/me`, backend mengembalikan semua riwayat user login.

## Upload Foto

Endpoint:

```http
POST /api/v1/employees/{uid}/photo
Content-Type: multipart/form-data
Authorization: Bearer <token>
```

Form-data:

- field file: `photo`
- ekstensi: `.jpg`, `.jpeg`, `.png`
- maksimal: 2 MB

Response sukses:

```json
{
  "success": true,
  "message": "Foto berhasil diupload",
  "data": {
    "photo_url": "/uploads/photos/<filename>"
  }
}
```

Untuk render gambar di FE, gabungkan dengan origin backend:

```ts
const imageUrl = `${BACKEND_ORIGIN}${photo_url}`;
```

Contoh: `http://localhost:8080/uploads/photos/file.png`.

## Rekomendasi Struktur Halaman FE

Minimal halaman yang cocok dengan backend saat ini:

- Login: email, password.
- Register: fullname, nik, email, password.
- Dashboard: ringkasan user, absensi hari ini, cuti pending, payroll status.
- Users: list akun, detail user, delete user.
- Employees: list karyawan dari `/employees`, detail karyawan, form data karyawan, upload foto.
- Absensi Saya: clock in, clock out, filter bulan/tahun, riwayat.
- Absensi Hari Ini: tabel semua absensi hari ini.
- Cuti Saya: form pengajuan, riwayat status.
- Approval Cuti: filter status, approve/reject dengan catatan.
- Penggajian Saya: riwayat gaji user login.
- Payroll Admin: create payroll, filter periode, mark paid.
- Master Divisi: CRUD divisi.
- Master Jabatan: CRUD jabatan.

## Flow FE yang Disarankan

### Login Flow

1. Submit `POST /auth/login`.
2. Simpan `data.token` dan `data.user`.
3. Arahkan ke dashboard.
4. Setiap request protected memakai Authorization header.
5. Jika menerima `401`, clear session dan kembali ke login.

### Employee Setup Flow

1. Buat akun lewat register atau data user sudah tersedia dari admin.
2. Ambil master data: `GET /divisi` dan `GET /jabatan`.
3. Submit detail: `POST /employees/{uid}/data`.
4. Upload foto opsional: `POST /employees/{uid}/photo`.
5. Ambil detail lagi: `GET /employees/{uid}/data`.

### Attendance Flow

1. Halaman absensi user memanggil `GET /absensi/me`.
2. Jika belum ada absensi hari ini, tampilkan tombol clock in.
3. Setelah clock in, tampilkan jam masuk dan tombol clock out.
4. Setelah clock out, disable tombol clock out.
5. Error yang mungkin tampil: `sudah clock in hari ini`, `belum clock in hari ini`, `sudah clock out hari ini`.

### Leave Flow

1. User submit `POST /cuti`.
2. Status awal selalu `pending`.
3. User melihat riwayat di `GET /cuti/me`.
4. Admin/HR melihat list `GET /cuti?status=pending`.
5. Admin/HR submit `PUT /cuti/{id}/approve` dengan status `approved` atau `rejected`.

### Payroll Flow

1. Admin pilih user dari `GET /users` atau employee dari `GET /employees`.
2. Admin submit payroll `POST /penggajian`.
3. Backend menghitung potongan BPJS dan total gaji.
4. Admin mark paid via `PUT /penggajian/{id}/bayar`.
5. User melihat gaji miliknya via `GET /penggajian/me`.

## Fetch Wrapper Contoh

```ts
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080/api/v1";
const BACKEND_ORIGIN = API_BASE_URL.replace("/api/v1", "");

type ApiResponse<T> = {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
};

async function apiFetch<T>(path: string, options: RequestInit = {}) {
  const token = localStorage.getItem("token");
  const headers = new Headers(options.headers);

  if (!(options.body instanceof FormData)) {
    headers.set("Content-Type", "application/json");
  }

  if (token) {
    headers.set("Authorization", `Bearer ${token}`);
  }

  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...options,
    headers,
  });

  const payload = (await response.json()) as ApiResponse<T>;

  if (!response.ok || !payload.success) {
    if (response.status === 401) {
      localStorage.removeItem("token");
      localStorage.removeItem("user");
    }
    throw new Error(payload.error || payload.message || "Request gagal");
  }

  return payload.data as T;
}
```

## Validasi Form FE

Validasi yang sebaiknya dilakukan di FE sebelum submit:

- Register: `fullname`, `nik`, `email`, `password`; password minimal 8 karakter; email valid.
- Login: `email`, `password`.
- Employee data: tanggal `YYYY-MM-DD`; nomor telepon sebagai string; `is_bpjs_tk` boolean.
- Cuti: `tipe_cuti`, `tanggal_mulai`, `tanggal_akhir`, `alasan`; tanggal akhir tidak lebih awal dari tanggal mulai.
- Payroll: `uid`, `periode`, `gaji_pokok`; angka tidak negatif untuk gaji/tunjangan/potongan.
- Upload foto: tipe file jpg/jpeg/png, ukuran maksimal 2 MB.

## Catatan Risiko untuk FE

- Belum ada role/permission di backend. Semua user yang punya token bisa mengakses endpoint protected termasuk master data, approval cuti, payroll, dan delete user.
- Belum ada endpoint update user atau update employee data meskipun ada fungsi repository. FE saat ini hanya bisa create detail employee, bukan edit lewat endpoint.
- `POST /employees/{uid}/data` memakai field `uid` unik, sehingga submit kedua untuk user yang sama kemungkinan gagal dari database.
- Parsing tanggal di beberapa handler mengabaikan error parse. FE harus menjaga format tanggal valid.
- Delete memakai soft delete untuk model yang punya `DeletedAt` seperti User dan DataUser; master Divisi/Jabatan tidak punya timestamp/delete marker.
- Response middleware auth terbaru mengembalikan error tanpa wrapper `success/message` saat token tidak ada/tidak valid: `{"error":"Token tidak ditemukan"}` atau `{"error":"Token tidak valid"}`. Fetch wrapper harus toleran terhadap bentuk ini.

## Prioritas Implementasi FE

Urutan paling efisien untuk membangun frontend:

1. Setup API client + auth storage + route guard.
2. Login/register.
3. Layout dashboard protected.
4. Master divisi dan jabatan.
5. Users dan employees.
6. Absensi user.
7. Cuti user dan approval.
8. Payroll admin dan payroll user.
9. Upload foto dan render avatar.
10. Polish error/loading/empty states.
