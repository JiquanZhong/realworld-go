create table public.users
(
    id         bigserial
        primary key,
    name       text                      not null,
    email      text                      not null
        constraint uni_users_email
            unique,
    age        bigint,
    password   text                      not null,
    role       text default 'user'::text not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    delete_at  timestamp with time zone
);

alter table public.users
    owner to postgres;

create index idx_users_delete_at
    on public.users (delete_at);

create table public.mcp_services
(
    id            bigserial
        primary key,
    name          text not null,
    icon_url      text,
    description   text,
    endpoint      text,
    json_schema   text,
    category      text,
    install_count bigint,
    rating_avg    numeric,
    is_visible    boolean default true,
    submitter_id  bigint,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    delete_at     timestamp with time zone
);

alter table public.mcp_services
    owner to postgres;

create index idx_mcp_services_delete_at
    on public.mcp_services (delete_at);

create table public.mcp_tags
(
    id         bigserial
        primary key,
    name       text not null
        constraint uni_mcp_tags_name
            unique,
    created_at timestamp with time zone
);

alter table public.mcp_tags
    owner to postgres;

create table public.mcp_service_tags
(
    mcp_id bigint not null
        constraint fk_mcp_service_tags_mcp_service
            references public.mcp_services
            on update cascade on delete cascade,
    tag_id bigint not null
        constraint fk_mcp_service_tags_tag
            references public.mcp_tags
            on update cascade on delete cascade,
    primary key (mcp_id, tag_id)
);

alter table public.mcp_service_tags
    owner to postgres;

create table public.mcp_ratings
(
    id             bigserial
        primary key,
    user_id        bigint,
    mcp_service_id bigint,
    stars          smallint,
    comment        text,
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone
);

alter table public.mcp_ratings
    owner to postgres;

create unique index idx_user_mcp_service
    on public.mcp_ratings (user_id, mcp_service_id);

create table public.mcp_favorites
(
    user_id        bigint not null,
    mcp_service_id bigint not null,
    created_at     timestamp with time zone,
    primary key (user_id, mcp_service_id)
);

alter table public.mcp_favorites
    owner to postgres;

create table public.mcp_installations
(
    id             bigserial
        primary key,
    mcp_service_id bigint,
    user_id        bigint,
    last_used_at   timestamp with time zone,
    created_at     timestamp with time zone
);

alter table public.mcp_installations
    owner to postgres;

create unique index idx_install_user_mcp
    on public.mcp_installations (mcp_service_id, user_id);

create table public.mcp_tokens
(
    id             bigserial
        primary key,
    token          text   not null
        constraint uni_mcp_tokens_token
            unique,
    user_id        bigint not null,
    mcp_service_id bigint not null,
    scope          text   not null,
    expires_at     timestamp with time zone,
    revoked        boolean default false,
    created_at     timestamp with time zone,
    last_used_at   timestamp with time zone
);

alter table public.mcp_tokens
    owner to postgres;

