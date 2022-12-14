CREATE TABLE book_user_stamps (
    id CHAR(36) NOT NULL,
    book_id CHAR(36) NOT NULL,
    book_series_id CHAR(36) NOT NULL,
    page_num INT NOT NULL,
    x INT NOT NULL,
    y INT NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    stamp_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (book_series_id) REFERENCES book_series(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (stamp_id) REFERENCES stamps(id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
