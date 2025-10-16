-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create enums
CREATE TYPE order_status AS ENUM (
    'DRAFT', 'CONFIRMED', 'SOURCING', 'PICKING', 
    'PACKED', 'SHIPPED', 'DELIVERED', 'CANCELLED', 'ON_HOLD'
);

CREATE TYPE stock_movement_reason AS ENUM (
    'PURCHASE_RECEIVE', 'SALE', 'WASTE', 'DAMAGED', 
    'EXPIRED', 'COUNT_CORRECTION', 'RETURN_FROM_CLIENT', 
    'RETURN_TO_SUPPLIER', 'INTERNAL_TRANSFER'
);

CREATE TYPE location_type AS ENUM (
    'WAREHOUSE', 'SUPPLIER', 'IN_TRANSIT', 'CLIENT'
);

CREATE TYPE client_status AS ENUM (
    'ACTIVE', 'SUSPENDED', 'INACTIVE'
);

CREATE TYPE audit_action AS ENUM (
    'CREATE', 'UPDATE', 'DELETE', 'ADJUST_QUANTITY', 
    'ADD_PHOTO', 'STATUS_CHANGE'
);

-- Create tables (your full schema here)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    role VARCHAR NOT NULL,
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_name VARCHAR NOT NULL,
    yacht_name VARCHAR,
    contact_person VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    phone VARCHAR,
    billing_address TEXT,
    default_delivery_address TEXT,
    account_status client_status NOT NULL DEFAULT 'ACTIVE',
    payment_terms VARCHAR,
    credit_limit NUMERIC(12,2),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Add indexes
CREATE INDEX idx_clients_company ON clients(company_name);
CREATE INDEX idx_clients_status ON clients(account_status);

-- Continue with all other tables...
