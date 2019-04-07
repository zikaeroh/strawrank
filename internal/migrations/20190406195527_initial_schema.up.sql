BEGIN;

CREATE TABLE polls (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    question text,
    choices text[]
);

CREATE TABLE ballots (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    poll_id int REFERENCES polls(id),
    user_xid text,
    user_ip text,
    votes int[]
);

COMMIT;