CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4 (),
	phone VARCHAR(50) NOT NULL UNIQUE,
	email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(25) NOT NULL UNIQUE,
    user_type INT NOT NULL,
    portal_access_status INT NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE units (
    id uuid DEFAULT uuid_generate_v4 (),
    is_occupied BOOLEAN NOT NULL,
    address_line_one VARCHAR(50) NOT NULL,
    address_line_two VARCHAR(50),
    city VARCHAR(50) NOT NULL,
    state VARCHAR(10) NOT NULL,
    zip VARCHAR(30) NOT NULL,
    PRIMARY KEY (id),
    owner_id uuid
        CONSTRAINT fk_user REFERENCES users(id)
);

CREATE TABLE leases (
    id uuid DEFAULT uuid_generate_v4 (),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    rent MONEY NOT NULL,
    is_active BOOLEAN NOT NULL,
    PRIMARY KEY (id),
    tenant_id uuid 
        CONSTRAINT fk_user REFERENCES users(id),
    unit_id uuid
        CONSTRAINT fk_unit REFERENCES units(id)
);

CREATE TABLE service_requests (
    id uuid DEFAULT uuid_generate_v4 (),
    description VARCHAR(500) NOT NULL,
    preferred_date_one TIMESTAMP,
    preferred_date_two TIMESTAMP,
    preferred_date_three TIMESTAMP,
    status INT NOT NULL,
    start_date DATE,
    end_date DATE,
    PRIMARY KEY (id),
    lease_id uuid
        CONSTRAINT fk_lease REFERENCES leases(id)
);

CREATE TABLE payment_instruments (
    id uuid DEFAULT uuid_generate_v4 (),
    card_number VARCHAR(30) NOT NULL UNIQUE,
    card_expiration VARCHAR(8) NOT NULL UNIQUE,
    card_cvv VARCHAR(3) NOT NULL UNIQUE,
    bank VARCHAR(30) NOT NULL UNIQUE,
    routing_number VARCHAR(30) NOT NULL UNIQUE,
    account_type INT NOT NULL,
    next_recurring_payment DATE,
    PRIMARY KEY (id),
    tenant_id uuid
        CONSTRAINT fk_user REFERENCES users(id)
);

CREATE TABLE ledger_entries (
    id uuid DEFAULT uuid_generate_v4 (),
    amount MONEY NOT NULL,
    type INT NOT NULL,
    PRIMARY KEY (id),
    lease_id uuid
        CONSTRAINT fk_lease REFERENCES leases(id),
    instrument_id uuid
        CONSTRAINT fk_payment_instrument REFERENCES payment_instruments(id)
);