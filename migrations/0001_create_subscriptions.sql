CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,         
    price INTEGER NOT NULL,                     
    user_id UUID NOT NULL,                      
    start_date DATE NOT NULL,                   
    end_date DATE,                              
    created_at TIMESTAMP DEFAULT NOW(),         
    updated_at TIMESTAMP DEFAULT NOW()          
);

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_service
    ON subscriptions (user_id, service_name);

