CREATE TABLE IF NOT EXISTS showcase.rides (
    id uuid PRIMARY KEY,
    user_id uuid REFERENCES showcase.users(id),
    bike_id uuid REFERENCES showcase.bikes(id),
    start_date timestamptz NOT NULL DEFAULT now(),
    end_date timestamptz NULL DEFAULT NULL
)
