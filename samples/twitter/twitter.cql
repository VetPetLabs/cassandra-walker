CREATE KEYSPACE twitter
  WITH REPLICATION = {
   'class' : 'SimpleStrategy',
   'replication_factor' : 1
 };

CREATE TABLE twitter.tweet (
	user_id bigint,
	tweet_id varchar,
	body varchar,
	create_time int,
	PRIMARY KEY (user_id,tweet_id)
);

CREATE TABLE twitter.user (
	user_id int,
	user_name varchar,
	full_name varchar,
	created_time bigint,
	PRIMARY KEY (user_id)
);