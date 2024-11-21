-- Create snippet table
CREATE TABLE snippets (
                          id integer NOT NULL PRIMARY KEY AUTO_INCREMENT,
                          title VARCHAR(100) NOT NULL,
                          content TEXT NOT NULL,
                          created DATETIME NOT NULL,
                          expires DATETIME NOT NULL
);
-- Create index on created field for snippets...
CREATE INDEX idx_snippets_created ON snippets(created);

-- Insert test data
-- Add test data's to the database.
INSERT INTO snippets (title, content, created, expires)
VALUES (
           'An old silent pond',
           'An old silent pond....\nA frog jumps into the pond.\nsplash! Silence again.\n\n-Matsuo Basho',
           UTC_TIMESTAMP(),
           DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
       );

-- Add test data's to the database.
INSERT INTO snippets (title, content, created, expires)
VALUES (
           'Over the wintry forest',
           'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n- Natsume Soseki',
           UTC_TIMESTAMP(),
           DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
       );

-- Add test data's to the database.
INSERT INTO snippets (title, content, created, expires)
VALUES (
           'First autumn morning',
           'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n- Murakami Kijo',
           UTC_TIMESTAMP(),
           DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
       );
