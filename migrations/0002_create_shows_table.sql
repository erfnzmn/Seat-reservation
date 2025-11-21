CREATE TABLE shows (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    movie_title VARCHAR(255) NOT NULL,
    description TEXT NULL,
    hall_id BIGINT UNSIGNED NOT NULL,
    starts_at DATETIME NOT NULL,
    ends_at DATETIME NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    CONSTRAINT fk_shows_hall
        FOREIGN KEY (hall_id) REFERENCES halls(id)
        ON DELETE CASCADE,

    INDEX idx_shows_hall (hall_id),
    INDEX idx_shows_start_time (starts_at)
);
