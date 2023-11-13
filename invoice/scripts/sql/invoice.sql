# INVOICES

CREATE TABLE keel_invoice.invoice_status_vertex(
    id INT NOT NULL PRIMARY KEY,
    internal_name VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NULL
);

insert into keel_invoice.invoice_status_vertex values (0, 'creating', 'Creating', 'System is creating invoice. Wait please.');
insert into keel_invoice.invoice_status_vertex values (10, 'consulting', 'Client Consult', 'System is Consulting clients information. Wait please.');
insert into keel_invoice.invoice_status_vertex values (20, 'waiting_client', 'Waiting Client Data', 'Client information missing. Plase update client information.');
insert into keel_invoice.invoice_status_vertex values (30, 'created', 'Created', 'Invoice was created sucessfully.');
insert into keel_invoice.invoice_status_vertex values (40, 'delivered', 'Delivered', 'Invoice was delivered to client.');
insert into keel_invoice.invoice_status_vertex values (50, 'saw' , 'Saw', 'Invoice was saw by client.');
insert into keel_invoice.invoice_status_vertex values (60, 'canceled', 'Canceled', 'Invoice was canceled.');


 create table keel_invoice.invoice_status_edge(
	id INT NOT NULL PRIMARY KEY,
    from_id INT NULL,
    to_id INT NOT NULL,
    description VARCHAR(100) NOT NULL,
    index idx_from_id (from_id),
    index idx_to_id (to_id),
    unique (from_id, to_id),
    FOREIGN KEY (from_id) REFERENCES keel_invoice.invoice_status_vertex(id),
    FOREIGN KEY (to_id) REFERENCES keel_invoice.invoice_status_vertex(id)
);

insert into keel_invoice.invoice_status_edge values (1, null, 0, 'creating invoice');
insert into keel_invoice.invoice_status_edge values (2, 0, 10, 'consulting client information');
insert into keel_invoice.invoice_status_edge values (3, 10, 20, 'waiting client information');
insert into keel_invoice.invoice_status_edge values (4, 10, 30, 'created after consulting client information');
insert into keel_invoice.invoice_status_edge values (5, 20, 30, 'created after client information was updated');
insert into keel_invoice.invoice_status_edge values (6, 20, 60, 'calceled after waiting client information');
insert into keel_invoice.invoice_status_edge values (7, 30, 60, 'canceled after created');
insert into keel_invoice.invoice_status_edge values (8, 30, 40, 'delivered after created');
insert into keel_invoice.invoice_status_edge values (9, 40, 50, 'saw after delivered');
insert into keel_invoice.invoice_status_edge values (10, 40, 60, 'canceled after delivered');
insert into keel_invoice.invoice_status_edge values (11, 50, 60, 'canceled after saw');


CREATE TABLE keel_invoice.invoice_payment_status_vertex(
    id INT NOT NULL PRIMARY KEY,
    internal_name VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NULL
);

insert into keel_invoice.invoice_payment_status_vertex values (0, 'opened', 'Opened', 'Waiting payment');
insert into keel_invoice.invoice_payment_status_vertex values (10, 'underpaid', 'Underpaid', 'Payment was partial paid');
insert into keel_invoice.invoice_payment_status_vertex values (20, 'paid', 'Paid', 'Payment was received');
insert into keel_invoice.invoice_payment_status_vertex values (30, 'overpaid', 'Overpaid', 'Payment was overpaid');
insert into keel_invoice.invoice_payment_status_vertex values (40, 'reversed', 'Reversed', 'Payment was canceled');

CREATE TABLE keel_invoice.invoice_payment_status_edge(
    id INT NOT NULL PRIMARY KEY,
    from_id INT NULL,
    to_id INT NOT NULL,
    description VARCHAR(100) NOT NULL,
    index idx_from_id (from_id),
    index idx_to_id (to_id),
    FOREIGN KEY (from_id) REFERENCES keel_invoice.invoice_payment_status_vertex(id),
    FOREIGN KEY (to_id) REFERENCES keel_invoice.invoice_payment_status_vertex(id)
);


insert into keel_invoice.invoice_payment_status_edge values (1, null, 0, 'creating invoice');
insert into keel_invoice.invoice_payment_status_edge values (2, 0, 10, 'partially paid');
insert into keel_invoice.invoice_payment_status_edge values (3, 0, 20, 'paid');
insert into keel_invoice.invoice_payment_status_edge values (4, 0, 30, 'overpaid');
insert into keel_invoice.invoice_payment_status_edge values (5, 10, 20, 'paid after partial payment');
insert into keel_invoice.invoice_payment_status_edge values (6, 10, 30, 'overpaid after partial payment');
insert into keel_invoice.invoice_payment_status_edge values (7, 20, 30, 'overpaid after payment');
insert into keel_invoice.invoice_payment_status_edge values (8, 10, 40, 'reversed after underpayment');
insert into keel_invoice.invoice_payment_status_edge values (9, 20, 40, 'reversed after payment');
insert into keel_invoice.invoice_payment_status_edge values (10, 30, 40, 'reversed after overpayment');


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
    invoice_status_vertex_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    INDEX idx_reference (reference ASC),
    INDEX idx_business_id (business_id ASC),
    INDEX idx_customer_id (customer_id ASC),
    INDEX idx_date (date ASC),
    INDEX idx_due (due ASC),
    INDEX idx_invoice_status_vertex_id (invoice_status_vertex_id ASC),
    INDEX idx_note_id (note_id ASC),
    FOREIGN KEY (invoice_status_vertex_id) REFERENCES keel_invoice.invoice_status_vertex(id),
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
    invoice_status_edge_id int not null,
    created_at timestamp not null,
    author varchar(100) null,
    description varchar(500) null,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_invoice_status_edge_id (invoice_status_edge_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id),
    FOREIGN KEY (invoice_status_edge_id) REFERENCES keel_invoice.invoice_status_edge(id)
);

create table keel_invoice.invoice_payment_status_log (
    id int not null primary key,
    invoice_id varchar(50) not null,
    invoice_payment_status_edge_id int not null,
    created_at timestamp not null,
    author varchar(100) null,
    description varchar(500) null,
    INDEX idx_invoice_id (invoice_id ASC),
    INDEX idx_invoice_payment_status_edge_id (invoice_payment_status_edge_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id),
    FOREIGN KEY (invoice_payment_status_edge_id) REFERENCES keel_invoice.invoice_payment_status_edge(id)
);
