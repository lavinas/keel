# INVOICES

CREATE TABLE keel_invoice.invoice_vertex(
    class VARCHAR(50) NOT NULL,
    id VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(500) NULL,
    PRIMARY KEY (class, id)
);

insert into keel_invoice.invoice_vertex values ('invoice', 'none', 'Started', 'System is started.');
insert into keel_invoice.invoice_vertex values ('invoice', 'getting', 'Consulting Client Data', 'System is Consulting clients information. Wait please.');
insert into keel_invoice.invoice_vertex values ('invoice', 'waiting', 'Waiting Client Data', 'Client information missing. Plase update client information.');
insert into keel_invoice.invoice_vertex values ('invoice', 'created', 'Created', 'Invoice was created sucessfully.');
insert into keel_invoice.invoice_vertex values ('invoice', 'delivered', 'Delivered', 'Invoice was delivered to client.');
insert into keel_invoice.invoice_vertex values ('invoice', 'saw' , 'Saw', 'Invoice was saw by client.');
insert into keel_invoice.invoice_vertex values ('invoice', 'canceled', 'Canceled', 'Invoice was canceled.');
insert into keel_invoice.invoice_vertex values ('payment', 'none', 'Started', 'System is started.');
insert into keel_invoice.invoice_vertex values ('payment', 'opened', 'Opened', 'Waiting payment');
insert into keel_invoice.invoice_vertex values ('payment', 'underpaid', 'Underpaid', 'Payment was partial paid');
insert into keel_invoice.invoice_vertex values ('payment', 'paid', 'Paid', 'Payment was received');
insert into keel_invoice.invoice_vertex values ('payment', 'overpaid', 'Overpaid', 'Payment was overpaid');
insert into keel_invoice.invoice_vertex values ('payment', 'reversed', 'Reversed', 'Payment was canceled');

 create table keel_invoice.invoice_edge(
    class VARCHAR(50) NOT NULL,
    from_invoice_vertex_id VARCHAR(50) NOT NULL,
    to_invoice_vertex_id VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    PRIMARY key (class, from_invoice_vertex_id, to_invoice_vertex_id),
    FOREIGN KEY (class, from_invoice_vertex_id) REFERENCES keel_invoice.invoice_vertex(class, id),
    FOREIGN KEY (class, to_invoice_vertex_id) REFERENCES keel_invoice.invoice_vertex(class, id)
);


insert into keel_invoice.invoice_edge values ('invoice',  'none', 'getting', 'creating invoice');
insert into keel_invoice.invoice_edge values ('invoice', 'getting', 'waiting', 'waiting client information');
insert into keel_invoice.invoice_edge values ('invoice', 'getting', 'created', 'created after consulting client information');
insert into keel_invoice.invoice_edge values ('invoice', 'waiting', 'created', 'created after client information was updated');
insert into keel_invoice.invoice_edge values ('invoice', 'waiting', 'canceled', 'calceled after waiting client information');
insert into keel_invoice.invoice_edge values ('invoice', 'created', 'canceled', 'canceled after created');
insert into keel_invoice.invoice_edge values ('invoice', 'created', 'delivered', 'delivered after created');
insert into keel_invoice.invoice_edge values ('invoice', 'delivered', 'saw', 'saw after delivered');
insert into keel_invoice.invoice_edge values ('invoice', 'delivered', 'canceled', 'canceled after delivered');
insert into keel_invoice.invoice_edge values ('invoice', 'saw', 'canceled', 'canceled after saw');
insert into keel_invoice.invoice_edge values ('payment',  'none', 'opened', 'creating invoice');
insert into keel_invoice.invoice_edge values ('payment', 'opened', 'underpaid', 'partially paid');
insert into keel_invoice.invoice_edge values ('payment', 'opened', 'paid', 'paid');
insert into keel_invoice.invoice_edge values ('payment', 'opened', 'overpaid', 'overpaid');
insert into keel_invoice.invoice_edge values ('payment', 'underpaid', 'paid', 'paid after partial payment');
insert into keel_invoice.invoice_edge values ('payment', 'underpaid', 'overpaid', 'overpaid after partial payment');
insert into keel_invoice.invoice_edge values ('payment', 'paid', 'overpaid', 'overpaid after payment');
insert into keel_invoice.invoice_edge values ('payment', 'underpaid', 'reversed', 'reversed after underpayment');
insert into keel_invoice.invoice_edge values ('payment', 'paid', 'reversed', 'reversed after payment');
insert into keel_invoice.invoice_edge values ('payment', 'overpaid', 'reversed', 'reversed after overpayment');

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
    business_id VARCHAR(50) NOT NULL,
    reference VARCHAR(10) NOT NULL UNIQUE,
    customer_id VARCHAR(50) NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    date DATE NOT NULL,
    due DATE NOT NULL,
    note_id VARCHAR(50) NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    FOREIGN KEY (note_id) REFERENCES keel_invoice.invoice_note (id),
    FOREIGN KEY (business_id) REFERENCES keel_invoice.invoice_client (id),
    FOREIGN KEY (customer_id) REFERENCES keel_invoice.invoice_client (id),
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


CREATE TABLE keel_invoice.invoice_status (
    invoice_id VARCHAR(50) NOT NULL,
    invoice_vertex_class VARCHAR(50) NOT NULL,
    invoice_vertex_id VARCHAR(50) NOT NULL,
    PRIMARY KEY (invoice_id, invoice_vertex_class),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice (id),
    FOREIGN KEY (invoice_vertex_class, invoice_vertex_id) REFERENCES keel_invoice.invoice_vertex (class, id)
);

create table keel_invoice.invoice_status_log(
    id varchar(50) not null primary key,
    invoice_id varchar(50) not null,
    class VARCHAR(50) NOT NULL,
    from_invoice_vertex_id VARCHAR(50) NOT NULL,
    to_invoice_vertex_id VARCHAR(50) NOT NULL,
    created_at timestamp not null,
    author varchar(100) null,
    description varchar(500) null,
    INDEX idx_invoice_id (invoice_id ASC),
    FOREIGN KEY (invoice_id) REFERENCES keel_invoice.invoice(id),
    FOREIGN KEY (class, from_invoice_vertex_id, to_invoice_vertex_id) REFERENCES keel_invoice.invoice_edge(class, from_invoice_vertex_id, to_invoice_vertex_id)
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