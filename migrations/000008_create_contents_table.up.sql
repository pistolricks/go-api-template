CREATE TABLE IF NOT EXISTS contents
(
    id text PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    src text NOT NULL,
    type text NOT NULL,
    size float NOT NULL,
    width integer NOT NULL,
    height integer NOT NULL,
    sort_order integer NOT NULL,
    user_id text NOT NULL
)