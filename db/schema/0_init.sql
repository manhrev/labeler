START TRANSACTION;

-- User 
CREATE TABLE users (
    id              BIGSERIAL,
    username        text UNIQUE NOT NULL,
    password        text NOT NULL,
    display_name    text NOT NULL,
    created_at      timestamptz DEFAULT now(),
    updated_at      timestamptz DEFAULT now(),
    email           text NOT NULL,

    PRIMARY KEY (id)
);

CREATE TYPE category AS ENUM ('C_BASKETBALL_COMPETITION', 'C_BASKETBALL_COMPETITOR', 'C_FOOTBALL_COMPETITION', 'C_FOOTBALL_COMPETITOR');
CREATE TYPE background_type AS ENUM ('BT_NONE','BT_OUTSIDE','BT_FULL');
-- url selected: 0 1 2 3 null, 0 is no fit
CREATE TABLE images (
    id                  BIGINT,
    category            category NOT NULL,     
    background_type      background_type DEFAULT 'BT_NONE',           
    labeler_id          BIGINT,
    name                text NOT NULL,
    display_name        text,
    url1                text NOT NULL,
    url2                text NOT NULL,
    url3                text NOT NULL,
    url_selected        smallint,
    created_at          timestamptz DEFAULT now(),
    updated_at          timestamptz DEFAULT now(),

    PRIMARY KEY (id, category)
);


COMMIT;