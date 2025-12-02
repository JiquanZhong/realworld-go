create table public.users
(
    id         bigserial
        primary key,
    name       text                         not null,
    email      text                         not null
        constraint uni_users_email
            unique,
    age        bigint,
    password   text                         not null,
    role       text    default 'user'::text not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    delete_at  timestamp with time zone,
    icon_url   varchar default 'https://img.icons8.com/?size=100&id=20750&format=png&color=000000'::character varying
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


-- ============================================
-- 2. 板块表 forum_boards
-- ============================================
CREATE TABLE forum_boards (
    id          BIGSERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    slug        VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    sort_order  INT NOT NULL DEFAULT 0,
    is_hidden   BOOLEAN NOT NULL DEFAULT FALSE
);

-- ============================================
-- 3. 主题状态枚举
-- ============================================
CREATE TYPE forum_topic_status AS ENUM ('normal', 'locked', 'archived');

-- ============================================
-- 4. 主题表 forum_topics
-- ============================================
CREATE TABLE forum_topics (
    id                  BIGSERIAL PRIMARY KEY,
    board_id            BIGINT NOT NULL REFERENCES forum_boards(id) ON DELETE CASCADE,
    user_id             BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    title               VARCHAR(255) NOT NULL,
    status              forum_topic_status NOT NULL DEFAULT 'normal',
    view_count          INT NOT NULL DEFAULT 0,
    reply_count         INT NOT NULL DEFAULT 0,
    last_post_at        TIMESTAMPTZ,
    last_post_user_id   BIGINT REFERENCES users(id),

    is_deleted          BOOLEAN NOT NULL DEFAULT FALSE,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_forum_topics_board_created_at
    ON forum_topics(board_id, created_at DESC);

CREATE INDEX idx_forum_topics_board_last_post_at
    ON forum_topics(board_id, last_post_at DESC);

-- ============================================
-- 5. 帖子表 forum_posts（支持楼中楼）
-- ============================================
CREATE TABLE forum_posts (
    id              BIGSERIAL PRIMARY KEY,
    topic_id        BIGINT NOT NULL REFERENCES forum_topics(id) ON DELETE CASCADE,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE RESTRICT,

    -- 父楼层（楼中楼/引用楼层）
    parent_post_id  BIGINT REFERENCES forum_posts(id) ON DELETE SET NULL,

    content         TEXT NOT NULL,
    is_deleted      BOOLEAN NOT NULL DEFAULT FALSE,
    edit_count      INT NOT NULL DEFAULT 0,
    last_edited_at  TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_forum_posts_topic_created_at
    ON forum_posts(topic_id, created_at ASC);

CREATE INDEX idx_forum_posts_user_created_at
    ON forum_posts(user_id, created_at DESC);

-- ============================================
-- 6. 点赞表 forum_post_likes
-- ============================================
CREATE TABLE forum_post_likes (
    user_id     BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id     BIGINT NOT NULL REFERENCES forum_posts(id) ON DELETE CASCADE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, post_id)
);

CREATE INDEX idx_forum_post_likes_post_id
    ON forum_post_likes(post_id);

-- ============================================
-- 7. 通知类型枚举（可选）
-- ============================================
CREATE TYPE forum_notification_type AS ENUM ('reply', 'mention', 'system');

-- ============================================
-- 8. 通知表 forum_notifications（可选）
-- ============================================
CREATE TABLE forum_notifications (
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type            forum_notification_type NOT NULL,
    topic_id        BIGINT REFERENCES forum_topics(id) ON DELETE CASCADE,
    post_id         BIGINT REFERENCES forum_posts(id) ON DELETE CASCADE,
    is_read         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_forum_notifications_user_is_read
    ON forum_notifications(user_id, is_read, created_at DESC);