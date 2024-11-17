CREATE TABLE IF NOT EXISTS showcase.bikes (
    id uuid PRIMARY KEY,
    lat float4,
    lon float4,
    image_url text,
    available boolean DEFAULT TRUE
)
