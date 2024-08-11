START TRANSACTION;

ALTER TABLE images 
    ADD column url1_title text NOT NULL DEFAULT '',
    ADD column url2_title text NOT NULL DEFAULT '',
    ADD column url3_title text NOT NULL DEFAULT '';
    
COMMIT;