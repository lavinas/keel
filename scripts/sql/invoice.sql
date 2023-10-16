# INVOICES

CREATE TABLE keel_invoice.invoice_status(
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE keel_invoice.invoice(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    reference VARCHAR(10) NOT NULL,
    client_id VARCHAR(50) NOT NULL,
    client_email VARCHAR(50) NOT NULL,
    client_nickname VARCHAR(50) NULL,
    client_name VARCHAR(100) NULL,
    client_document DECIMAL(20) NULL,
    client_phone DECIMAL(20) NULL,
    date DATE NOT NULL,
    due DATE NOT NULL,
    notes VARCHAR(200) NOT NULL,
    status_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_reference (reference ASC),
    INDEX idx_client_id (client_id ASC),
    INDEX idx_date (date ASC),
    INDEX idx_due (due ASC),
    INDEX idx_status (status_id ASC),
    FOREIGN KEY (client_id) REFERENCES keel_client.client(id),
    FOREIGN KEY (status_id) REFERENCES keel_invoice.invoice_status(id)
);

CREATE TABLE keel_invoice.invoice_item(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    invoice_id VARCHAR(50) NOT NULL,
    service_id VARCHAR(50) NOT NULL,
    service_name VARCHAR(50) NOT NULL,
    description VARCHAR(200) NULL,
    value DECIMAL(20, 2) NOT NULL,
    quantity DECIMAL(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_service_id (service_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES invoice(id)
);


