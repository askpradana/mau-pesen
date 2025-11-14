# Online Booking System

A lightweight Go-based booking system with basic calendar integration.

## Actors

- **Admin**: Manages consultants
- **Consultant**: Sets availability, approves/rejects bookings
- **Client**: Registers and books sessions

## Core Requirements

### Authentication

- Register with: **email, phone number, name**
- Login via **phone number + OTP** (alphanumeric 6 digits)
- Use **JWT** for auth sessions

### Booking System

- Browse consultants
- Book available slots
- Approve or reject

## User Roles

### Admin

- Add / manage consultants

### Consultant

- Set monthly availability (1-hour slots)
- Approve or reject booking requests

### Client

- Register
- Browse consultants
- Book available slots

## Booking Flow

1. Client selects consultant + available time
2. Sends booking request (includes purpose)
3. Consultant approves/rejects
4. On approval → create Google Calendar event

## Google Calendar Integration

- Use service account
- Sync **approved bookings only**
- Time slots fixed to **1-hour** between **09:00–17:00 (GMT+7)**
- Event includes:
  - Client name
  - Booking purpose

## Required Features

### Authentication (user)

- [ ] Register
- [ ] Login with OTP
- [ ] JWT token

### Booking System (user)

- [ ] Browse consultants
- [ ] Book available slots

### Booking System (consultant)

- [ ] Set availability
- [ ] Approve/reject

### Google Calendar (consultant)

- [ ] Integrate with google calendar
- [ ] Create event on approval

### Management (admin)

- [ ] Add/remove consultants

## Plus Points

- [ ] Session notes from consultant to users
- [ ] Logs on admin dashboard (booking requests)
- [ ] OTP via telegram/whatsapp
- [ ] Ratings post-session 1 to 5 with messages

## Example API

### GET consultant availability

```json
{
  "consultant_id": "123",
  "month": "2025-12",
  "timezone": "Asia/Jakarta",
  "slots": [
    { "date": "2024-12-15", "time": "09:00", "slot_id": "slot_001" },
    { "date": "2024-12-15", "time": "13:00", "slot_id": "slot_002" }
  ]
}
```

### POST booking

```json
{
  "consultant_id": "123",
  "slot_id": "slot_001",
  "booking_purpose": "Career counseling"
}
```