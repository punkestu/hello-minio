# MINIO

## Review
minIO dapat digunakan sebagai bucket/penyimpanan file. Teknologi ini berbentuk service sehingga kita bisa langsung menggunakannya secara remote dari service lain. Karena berbentuk service, maka untuk menyimpan dan mendapatkan file kita menggunakan metode yang serupa dengan upload dan download. "Mirip dbms tapi file"

## Alur Kerja
Pada dasarnya minIO adalah sebuah service untuk menyimpan file jadi secara garis besar metode penggunaannya sama dengan ketika kita menggunakan dbms. Alur kerjanya akan menjadi:
1. Buat bucket di minIO
2. Buat koneksi ke service minIO
3. Upload/Download file dari service kita ke bucket yang ada di service minIO

## Alur Pembuatan
1. Buat sebuah credential user pada service minIO
2. Konfigurasi hak akses untuk user tersebut
3. Simpan data credential
4. Koneksikan service kita dengan service minIO dengan credential yang kita miliki
5. Buat bucket menggunakan koneksi tersebut (atau bisa dari service minIO jika credential kita tidak memiliki akses untuk membuat bucket)
6. --- upload file ---
7. Baca file (bisa berbentuk file fisik maupun byte dari request)
8. Upload file ke bucket yang kita inginkan
9. --- download file ---
10. Ambil data byte file dari bucket
11. Simpan file (atau bisa langsung kirim ke user sebagai sebuah response tapi jangan lupa tambahkan content-type)

## Implementasi
- [x] Upload file
- [x] Download file
