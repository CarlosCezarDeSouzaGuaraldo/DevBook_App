    INSERT INTO users (name, nick, email, password)
    VALUES
    ('Carlos User 1', 'Carlos1', 'carlos1@carlos.carlos', '$2a$10$OM6Vn1byTfBKqQnO6aCbrexqUXQuHPFE3MqhRid4LnjvCgv/77OY6'),
    ('Carlos User 2', 'Carlos2', 'carlos2@carlos.carlos', '$2a$10$OM6Vn1byTfBKqQnO6aCbrexqUXQuHPFE3MqhRid4LnjvCgv/77OY6'),
    ('Carlos User 3', 'Carlos3', 'carlos3@carlos.carlos', '$2a$10$OM6Vn1byTfBKqQnO6aCbrexqUXQuHPFE3MqhRid4LnjvCgv/77OY6');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

INSERT INTO publications(title, content, author_id)
VALUES
("User's 1 publications", "It's the user's 1 publications!", 1),
("User's 2 publications", "It's the user's 2 publications!", 2),
("User's 3 publications", "It's the user's 3 publications!", 3);