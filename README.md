# mau-pesen-final — 2025 Production Ready

Project konsultasi online Go terlengkap, tercepat, dan terkuat di Indonesia.

## Fitur
- Register + OTP (disimpan di DB buat audit)
- JWT + Refresh Token + Logout
- Consultant full profile + rating auto avg
- Slot per jam (hour INT) — super cepat
- Booking anti double (SELECT FOR UPDATE)
- Approve → Google Calendar + Meet link otomatis
- Session notes terpisah
- Rating + auto update avg
- Admin FULL CRUD + ban user
- Logs audit trail
- Docker + Postgres + Redis

## Cara Jalanin (60 detik)
```bash
git clone https://github.com/rubiagatra/mau-pesen-final.git
cd mau-pesen-final
docker-compose up -d
go run cmd/api/main.go