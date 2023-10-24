# INVOICES

CREATE TABLE keel_invoice.invoice_status(
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE keel_invoice.invoice_note(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    business_id VARCHAR(50) NOT NULL,
    reference VARCHAR(50) NOT NULL UNIQUE,
    content VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_name (name ASC)
)

CREATE TABLE keel_invoice.invoice_client(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    client_id VARCHAR(50) NULL,
    name VARCHAR(100) NULL,
    document DECIMAL(20) NULL,
    phone DECIMAL(20) NULL,
    email VARCHAR(50) NULL,
    INDEX idx_nick (nickname ASC),
    INDEX idx_document (document ASC),
    INDEX idx_email (email ASC),
    INDEX idx_phone (phone ASC)
);

CREATE TABLE keel_invoice.invoice(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    reference VARCHAR(10) NOT NULL UNIQUE,
    business_id VARCHAR(50) NOT NULL,
    customer_id VARCHAR(50) NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    date DATE NOT NULL,
    due DATE NOT NULL,
    note_id VARCHAR(50) NULL,
    status_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_reference (reference ASC),
    INDEX idx_business_id (client_id ASC),
    INDEX idx_customer_id (customer_id ASC),
    INDEX idx_date (date ASC),
    INDEX idx_due (due ASC),
    INDEX idx_status (status_id ASC),
    FOREIGN KEY (status_id) REFERENCES keel_invoice.invoice_status(id),
    FOREIGN KEY (notes_id) REFERENCES keel_invoice.invoice_note (id),
    FOREIGN KEY (business_id) REFERENCES keel_invoice.invoice_client (id),
    FOREIGN KEY (customer_id) REFERENCES keel_invoice.invoice_client (id)
);

CREATE TABLE keel_invoice.invoice_item(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    invoice_id VARCHAR(50) NOT NULL,
    product_reference VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NULL,
    description VARCHAR(200) NULL,
    value DECIMAL(20, 2) NOT NULL,
    quantity DECIMAL(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_service_id (service_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES invoice(id)
);

CREATE TABLE keel_invoice.invoice_payment(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    invoice_id VARCHAR(50) NOT NULL,
    reference VARCHAR(50) NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_invoice_id (invoice_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id)
);

create table keel_invoice.invoice_status_log(
    id int not null primary key,
    invoice_id varchar(50) not null,
    status_id int not null,
    created_at timestamp not null,
    name varchar(50) not null,
    INDEX idx_invoice_id (invoice_id ASC)
);

create table keel_invoice.invoice_send_log(
    id int not null primary key,
    invoice_id varchar(50) not null,
    send_type varchar(50) not null,
    send_to varchar(50) not null,
    created_at timestamp not null,
    INDEX idx_invoice_id (invoice_id ASC)
);
