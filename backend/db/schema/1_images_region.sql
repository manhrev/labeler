START TRANSACTION;

alter table images add column region text default null;
alter table images add column display_order int default 0;

COMMIT;
