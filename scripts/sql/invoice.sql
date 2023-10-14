# INVOICES

CREATE TABLE keel_invoice.invoice_status(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
)

CREATE TABLE keel_invoice.invoice(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    sequence VARCHAR(10) NOT NULL,
    client_reference VARCHAR(50) NOT NULL,
    client_name VARCHAR(100) NOT NULL,
    client_document DECIMAL(20) NOT NULL,
    client_email VARCHAR(50) NOT NULL,
    date DATE NOT NULL,
    due DATE NOT NULL,
    notes VARCHAR(200) NOT NULL,
    status_id VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_number (number ASC),
    INDEX idx_client (client_id ASC),
    FOREIGN KEY (status_id) REFERENCES keel_invoice.invoice_status(id),
)

CREATE TABLE keel_invoice.invoice_item(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    invoice_id VARCHAR(50) NOT NULL,
    service_reference VARCHAR(50) NOT NULL,
    service_description VARCHAR(200) NOT NULL,
    service_value DECIMAL(20, 2) NOT NULL,
    service_quantity DECIMAL(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (invoice_id) REFERENCES invoice(id),
);

