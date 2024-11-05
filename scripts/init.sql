GRANT ALL PRIVILEGES ON k8s_db.* TO 'k8s_user'@'%';
FLUSH PRIVILEGES;

DROP TABLE IF EXISTS person;

CREATE TABLE person (
                              id INT AUTO_INCREMENT PRIMARY KEY,
                              name VARCHAR(255),
                              email VARCHAR(255),
                              phone VARCHAR(255)
);
