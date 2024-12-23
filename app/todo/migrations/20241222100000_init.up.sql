CREATE TABLE IF NOT EXISTS todo (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    user_id INT NOT NULL,
    is_done BOOLEAN NOT NULL DEFAULT FALSE,
    done_at TIMESTAMP without time zone,
    created_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS todo_user_id ON todo (user_id);

CREATE TABLE IF NOT EXISTS item (
    id SERIAL PRIMARY KEY,
    is_done BOOLEAN NOT NULL DEFAULT FALSE,
    title TEXT NOT NULL,
    todo_id INT NOT NULL REFERENCES todo(id),
    done_at TIMESTAMP without time zone,
    created_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at TIMESTAMP without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS item_todo_id ON item (todo_id);