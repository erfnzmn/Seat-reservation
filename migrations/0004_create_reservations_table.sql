CREATE TABLE reservations (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    show_id BIGINT UNSIGNED NOT NULL,
    seat_id BIGINT UNSIGNED NOT NULL,

    -- Customer information (no user account)
    customer_name VARCHAR(100) NOT NULL,
    customer_phone VARCHAR(20) NOT NULL,

    status ENUM('confirmed', 'cancelled') NOT NULL DEFAULT 'confirmed',

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    CONSTRAINT fk_res_show
        FOREIGN KEY (show_id) REFERENCES shows(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_res_seat
        FOREIGN KEY (seat_id) REFERENCES seats(id)
        ON DELETE CASCADE,

    -- Prevent double-booking of the same seat for the same show
    UNIQUE KEY uq_reservation_seat (show_id, seat_id),

    INDEX idx_res_show (show_id),
    INDEX idx_res_seat (seat_id)
);
