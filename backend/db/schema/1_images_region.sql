START TRANSACTION;

ALTER TABLE images 
    ADD column region text DEFAULT NULL,
    ADD column display_order int DEFAULT 0;

COMMIT;
