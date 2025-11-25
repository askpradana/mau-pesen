CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT,
    email TEXT UNIQUE,
    phone TEXT UNIQUE NOT NULL,
    role TEXT NOT NULL DEFAULT 'client' CHECK (role IN ('client', 'consultant', 'admin')),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Consultants (terpisah, lengkap)
CREATE TABLE IF NOT EXISTS consultants (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    speciality TEXT,
    bio TEXT,
    rating_avg NUMERIC(3,2) DEFAULT 0,
    total_rating INT DEFAULT 0,
    price_per_session INT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS otp_codes (
    id SERIAL PRIMARY KEY,
    phone TEXT,
    code TEXT,
    expires_at TIMESTAMPTZ NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS consultant_slots (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slot_id TEXT UNIQUE,
    consultant_id UUID REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    hour INT NOT NULL CHECK (hour >= 0 AND hour <= 23),
    available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(consultant_id, date, hour)
);

CREATE TABLE IF NOT EXISTS bookings (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id TEXT UNIQUE NOT NULL,
    client_id UUID REFERENCES users(id) ON DELETE SET NULL,
    consultant_id UUID REFERENCES users(id) ON DELETE SET NULL,
    slot_id UUID REFERENCES consultant_slots(id) ON DELETE SET NULL,
    date DATE NOT NULL,
    hour INT NOT NULL,
    purpose TEXT,
    status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected', 'cancelled')),
    google_event_id TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    update_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (consultant_id, date, hour)
);

CREATE TABLE IF NOT EXISTS session_notes (
    booking_id UUID PRIMARY KEY REFERENCES bookings(id) ON DELETE CASCADE,
    consultant_id UUID REFERENCES users(id),
    notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ratings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    booking_id UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES users(id),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    activity TEXT NOT NULL,
    metadata JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);