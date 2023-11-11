# INVOICES

CREATE TABLE keel_invoice.invoice_status(
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NULL
);

insert into keel_invoice.invoice_status values (0, 'Creating', 'System is creating invoice. Wait please.');
insert into keel_invoice.invoice_status values (10, 'Consulting client information', 'System is Consulting clients information. Wait please.');
insert into keel_invoice.invoice_status values (20, 'Waiting client information', 'Client information missing. Plase update client information.');
insert into keel_invoice.invoice_status values (30, 'Created', 'Invoice was created sucessfully.');
insert into keel_invoice.invoice_status values (40, 'Delivered', 'Invoice was delivered to client.');
insert into keel_invoice.invoice_status values (50, 'Saw', 'Invoice was saw by client.');
insert into keel_invoice.invoice_status values (60, 'Canceled', 'Invoice was canceled.');


 create table keel_invoice.invoice_status_flow(
	id INT NOT NULL PRIMARY KEY,
    from_id INT NULL,
    to_id INT NOT NULL,
    description VARCHAR(100) NOT NULL,
    index idx_from_id (from_id),
    index idx_to_id (to_id),
    FOREIGN KEY (from_id) REFERENCES keel_invoice.invoice_status(id),
    FOREIGN KEY (to_id) REFERENCES keel_invoice.invoice_status(id)
);

insert into keel_invoice.invoice_status_flow values (1, null, 0, 'creating invoice');
insert into keel_invoice.invoice_status_flow values (2, 0, 10, 'consulting client information');
insert into keel_invoice.invoice_status_flow values (3, 10, 20, 'waiting client information');
insert into keel_invoice.invoice_status_flow values (4, 10, 30, 'created after consulting client information');
insert into keel_invoice.invoice_status_flow values (5, 20, 30, 'created after client information was updated');
insert into keel_invoice.invoice_status_flow values (6, 20, 60, 'calceled after waiting client information');
insert into keel_invoice.invoice_status_flow values (7, 30, 60, 'canceled after created');
insert into keel_invoice.invoice_status_flow values (8, 30, 40, 'delivered after created');
insert into keel_invoice.invoice_status_flow values (9, 40, 50, 'saw after delivered');
insert into keel_invoice.invoice_status_flow values (10, 40, 60, 'canceled after delivered');
insert into keel_invoice.invoice_status_flow values (11, 50, 60, 'canceled after saw');


CREATE TABLE keel_invoice.invoice_payment_status(
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NULL
);

insert into keel_invoice.invoice_payment_status values (0, 'Opened', 'Waiting payment');
insert into keel_invoice.invoice_payment_status values (10, 'Underpaid', 'Payment was partial paid');
insert into keel_invoice.invoice_payment_status values (20, 'Paid', 'Payment was received');
insert into keel_invoice.invoice_payment_status values (30, 'Overpaid', 'Payment was overpaid');
insert into keel_invoice.invoice_payment_status values (40, 'Reversed', 'Payment was canceled');

CREATE TABLE keel_invoice.invoice_payment_status_flow(
    id INT NOT NULL PRIMARY KEY,
    from_id INT NULL,
    to_id INT NOT NULL,
    description VARCHAR(100) NOT NULL,
    index idx_from_id (from_id),
    index idx_to_id (to_id),
    FOREIGN KEY (from_id) REFERENCES keel_invoice.invoice_payment_status(id),
    FOREIGN KEY (to_id) REFERENCES keel_invoice.invoice_payment_status(id)
);


insert into keel_invoice.invoice_payment_status_flow values (1, null, 0, 'creating invoice');
insert into keel_invoice.invoice_payment_status_flow values (2, 0, 10, 'partially paid');
insert into keel_invoice.invoice_payment_status_flow values (3, 0, 20, 'paid');
insert into keel_invoice.invoice_payment_status_flow values (4, 0, 30, 'overpaid');
insert into keel_invoice.invoice_payment_status_flow values (5, 10, 20, 'paid after partial payment');
insert into keel_invoice.invoice_payment_status_flow values (6, 10, 30, 'overpaid after partial payment');
insert into keel_invoice.invoice_payment_status_flow values (7, 20, 30, 'overpaid after payment');
insert into keel_invoice.invoice_payment_status_flow values (8, 10, 40, 'reversed after underpayment');
insert into keel_invoice.invoice_payment_status_flow values (9, 20, 40, 'reversed after payment');
insert into keel_invoice.invoice_payment_status_flow values (10, 30, 40, 'reversed after overpayment');


CREATE TABLE keel_invoice.invoice_note(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    business_id VARCHAR(50) NOT NULL,
    reference VARCHAR(50) NOT NULL UNIQUE,
    content VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_name (reference ASC)
);

CREATE TABLE keel_invoice.invoice_client(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    nickname VARCHAR(50) NOT NULL,
    client_id VARCHAR(50) NULL,
    name VARCHAR(100) NULL,
    document DECIMAL(20) NULL,
    phone DECIMAL(20) NULL,
    email VARCHAR(50) NULL,
    created_at TIMESTAMP NOT NULL,
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
    INDEX idx_business_id (business_id ASC),
    INDEX idx_customer_id (customer_id ASC),
    INDEX idx_date (date ASC),
    INDEX idx_due (due ASC),
    INDEX idx_status_id (status_id ASC),
    INDEX idx_note_id (note_id ASC),
    FOREIGN KEY (status_id) REFERENCES keel_invoice.invoice_status(id),
    FOREIGN KEY (note_id) REFERENCES keel_invoice.invoice_note (id),
    FOREIGN KEY (business_id) REFERENCES keel_invoice.invoice_client (id),
    FOREIGN KEY (customer_id) REFERENCES keel_invoice.invoice_client (id)
);

CREATE TABLE keel_invoice.invoice_item(
    id VARCHAR(50) NOT NULL PRIMARY KEY,
    invoice_id VARCHAR(50) NOT NULL,
    service_reference VARCHAR(50) NOT NULL,
    description VARCHAR(200) NULL,
    amount DECIMAL(20, 2) NOT NULL,
    quantity DECIMAL(10) NOT NULL,
    INDEX idx_invoice_id (invoice_id ASC),
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

create table keel_invoice.invoice_delivery(
    id int not null primary key,
    invoice_id varchar(50) not null,
    method varchar(50) not null,
    address varchar(100) not null,
    created_at timestamp not null,
    INDEX idx_invoice_id (invoice_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id)
);

create table keel_invoice.invoice_status_log(
    id int not null primary key,
    invoice_id varchar(50) not null,
    status_from int null,
    status_to int not null,
    created_at timestamp not null,
    author varchar(100) null,
    description varchar(500) null,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_status_from (status_from ASC),
    INDEX idx_status_to (status_to ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id),
    FOREIGN KEY (status_from) REFERENCES keel_invoice.invoice_status(id),
    FOREIGN KEY (status_to) REFERENCES keel_invoice.invoice_status(id)
);

create table keel_invoice.invoice_payment_status_log (
    id int not null primary key,
    invoice_id varchar(50) not null,
    status_from int null,
    status_to int not null,
    created_at timestamp not null,
    author varchar(100) null,
    description varchar(500) null,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_status_from (status_from ASC),
    INDEX idx_status_to (status_to ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id),
    FOREIGN KEY (status_from) REFERENCES keel_invoice.invoice_payment_status(id),
    FOREIGN KEY (status_to) REFERENCES keel_invoice.invoice_payment_status(id)
);


