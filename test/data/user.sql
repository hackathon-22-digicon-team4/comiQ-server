INSERT
    IGNORE INTO `users` (id, password)
VALUES
    (
        'user1',
        '$2a$10$.pn8EY6zCytzgV3JW5dXYeZ2xnsUI2cmuCsFbbYlsuGGotKx4qOhO' -- password="password"
    ),
    (
        'user2',
        '$2a$10$777RHmBriDm7ilr64wim8OaQCAMiqzVS.Dwn.UnqlOSJDwcWxUi1m' -- password="password2",bcrypt
    );
