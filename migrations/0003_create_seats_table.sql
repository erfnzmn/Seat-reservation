CREATE TABLE seats (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    hall_id BIGINT UNSIGNED NOT NULL,
    row_number INT NOT NULL,
    seat_number INT NOT NULL,
    label VARCHAR(20) NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    CONSTRAINT fk_seats_hall
        FOREIGN KEY (hall_id) REFERENCES halls(id)
        ON DELETE CASCADE,

    -- A seat position inside a hall must be unique
    UNIQUE KEY uq_seat_position (hall_id, row_number, seat_number),

    INDEX idx_seats_hall (hall_id)
);
