CREATE TABLE waiting_list (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    show_id BIGINT UNSIGNED NOT NULL,

    -- Customer info (similar to reservation)
    customer_name VARCHAR(100) NOT NULL,
    customer_phone VARCHAR(20) NOT NULL,

    requested_seat_hint VARCHAR(50) NULL,
    
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    CONSTRAINT fk_waiting_show
        FOREIGN KEY (show_id) REFERENCES shows(id)
        ON DELETE CASCADE,

    -- Each customer can only be in the waiting list once per show
    UNIQUE KEY uq_waiting_phone_show (show_id, customer_phone),

    INDEX idx_waiting_order (show_id, created_at)
);
