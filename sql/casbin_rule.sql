create table casbin_rule
(
    id     int unsigned auto_increment
        primary key,
    p_type varchar(100) null,
    v0     varchar(100) null,
    v1     varchar(100) null,
    v2     varchar(100) null,
    v3     varchar(100) null,
    v4     varchar(100) null,
    v5     varchar(100) null,
    v6     varchar(25)  null,
    v7     varchar(25)  null,
    effect varchar(100) null,
    constraint p_type
        unique (p_type, v0, v1, v2, v3, v4, v5, v6, v7)
);

INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (1, 'p', ' alice', ' /users', ' GET', null, null, null, null, null, null, null);
INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (2, 'p', ' bob', ' /users', ' POST', null, null, null, null, null, null, null);
INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (3, 'p', ' role::admin', ' /users', ' GET', null, null, null, null, null, null, null);
INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (4, 'p', ' role::admin', ' /users', ' POST', null, null, null, null, null, null, null);
INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (5, 'p', ' role::admin', ' /users', ' DELETE', null, null, null, null, null, null, null);
INSERT INTO casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5, v6, v7, effect, ptype) VALUES (6, 'g', ' alice', ' role::admin', null, null, null, null, null, null, null, null);
