CREATE INDEX IF NOT EXISTS vendors_title_idx ON vendors USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS vendors_genres_idx ON vendors USING GIN (genres);