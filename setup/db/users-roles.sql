REVOKE CREATE ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON DATABASE tenantwin FROM PUBLIC;

CREATE ROLE user_service_readwrite;
GRANT CONNECT ON DATABASE tenantwin TO user_service_readwrite;
GRANT USAGE ON SCHEMA public TO user_service_readwrite;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO user_service_readwrite;
CREATE USER user_service WITH PASSWORD '1234';
GRANT user_service_readwrite TO user_service;

CREATE ROLE lease_service_readwrite;
GRANT CONNECT ON DATABASE tenantwin TO lease_service_readwrite;
GRANT USAGE ON SCHEMA public TO lease_service_readwrite;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE leases, units TO lease_service_readwrite;
CREATE USER lease_service WITH PASSWORD '1234';
GRANT lease_service_readwrite TO lease_service;


CREATE ROLE payment_service_readwrite;
GRANT CONNECT ON DATABASE tenantwin TO payment_service_readwrite;
GRANT USAGE ON SCHEMA public TO payment_service_readwrite;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE payment_instruments, ledger_entries TO payment_service_readwrite;
CREATE USER payment_service WITH PASSWORD '1234';
GRANT payment_service_readwrite TO payment_service;

CREATE ROLE request_service_readwrite;
GRANT CONNECT ON DATABASE tenantwin TO request_service_readwrite;
GRANT USAGE ON SCHEMA public TO request_service_readwrite;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE service_requests TO request_service_readwrite;
CREATE USER request_service WITH PASSWORD '1234';
GRANT request_service_readwrite TO request_service;