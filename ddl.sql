drop table if exists bulk_update;
drop table if exists bulk_insert;

create table bulk_update (
  id bigint primary key
  , val text
  , created_at timestamp with time zone
  , updated_at timestamp with time zone
);

create table bulk_insert (
  id bigint primary key
  , val text
  , created_at timestamp with time zone
);
