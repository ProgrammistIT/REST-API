CREATE TABLE IF NOT EXISTS threats (
    md5 VARCHAR(32) PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    severity VARCHAR(20) NOT NULL DEFAULT 'low',
    detected_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT md5_format
        CHECK (md5 ~ '^[a-fA-F0-9]{32}$'),

    CONSTRAINT name_not_empty
        CHECK (length(trim(name)) > 0),

    CONSTRAINT severity_valid
        CHECK (severity IN ('low', 'medium', 'high', 'critical'))
);
