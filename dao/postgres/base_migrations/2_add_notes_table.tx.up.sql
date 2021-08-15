CREATE TABLE notes (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  note VARCHAR(255),
  CONSTRAINT user_fk
      FOREIGN KEY(user_id) 
  	  REFERENCES users(id)
      ON DELETE cascade
      
);
