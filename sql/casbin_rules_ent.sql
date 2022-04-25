create table casbin_rules
(
    id    bigint auto_increment
        primary key,
    ptype varchar(255) default '' not null,
    v0    varchar(255) default '' not null,
    v1    varchar(255) default '' not null,
    v2    varchar(255) default '' not null,
    v3    varchar(255) default '' not null,
    v4    varchar(255) default '' not null,
    v5    varchar(255) default '' not null
)
    collate = utf8mb4_bin;

INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (1, 'p', 'alice', '/users', 'GET', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (2, 'p', 'bob', '/users', 'POST', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (3, 'p', 'role::admin', '/users', 'GET', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (4, 'p', 'role::admin', '/users', 'POST', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (5, 'p', 'role::admin', '/users', 'DELETE', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (6, 'g', 'alice', 'role::admin', '', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (7, 'g', 'bob', 'admin', '', '', '', '');
INSERT INTO casbin_rules (id, ptype, v0, v1, v2, v3, v4, v5) VALUES (8, 'g', 'randy', 'admin', '', '', '', '');
