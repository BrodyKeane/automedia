DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id          INT AUTO_INCREMENT NOT NULL,
  name        VARCHAR(128) NOT NULL,
  bio         VARCHAR(128) NOT NULL,
  avatar_path VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO user
  (name, bio, avatar_path)
VALUES
  ('Test', 'Hello World!', 'assets/avatars/default-avatar.png'),
  ('John', 'Hello!', 'assets/avatars/default-avatar.png'),
  ('Jane', 'World!', 'assets/avatars/default-avatar.png'),
  ('Jack', 'Bye!', 'assets/avatars/default-avatar.png');
