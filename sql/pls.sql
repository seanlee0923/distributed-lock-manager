CREATE TABLE pls_clients (
    id varchar(36) primary key,
    deployment_name varchar(100) not null,
    created_at datetime default current_timestamp
)

CREATE TABLE distribute_locks (
    deployment_name varchar(100) primary key,
    leader_id varchar(36) not null
)