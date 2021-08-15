CREATE TABLE notes (
  id INTEGER PRIMARY KEY,
  user_id INTEGER
    FOREIGN KEY(user_id) 
	  REFERENCES users(id),
  note VARCHAR(255),
);
