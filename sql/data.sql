INSERT INTO users (name, nick, email, password)
VALUES
('Carlos User 1', 'Carlos1', 'carlos@1.com', '$2a$10$/ozol93qvf4Kga.WKsbH9uOH5JXtK776F6BhMry6Qeiyw4kfQ.6s2'),
('Carlos User 2', 'Carlos2', 'carlos@2.com', '$2a$10$/ozol93qvf4Kga.WKsbH9uOH5JXtK776F6BhMry6Qeiyw4kfQ.6s2'),
('Carlos User 3', 'Carlos3', 'carlos@3.com', '$2a$10$/ozol93qvf4Kga.WKsbH9uOH5JXtK776F6BhMry6Qeiyw4kfQ.6s2'),
('Carlos User 4', 'Carlos4', 'carlos@4.com', '$2a$10$/ozol93qvf4Kga.WKsbH9uOH5JXtK776F6BhMry6Qeiyw4kfQ.6s2');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);