-- Create the metrics table
CREATE TABLE metrics (
    id SERIAL PRIMARY KEY,
    app_id INTEGER NOT NULL,
    component VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    hours INTEGER NOT NULL,
    metric_type VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);

-- Create the apps table
CREATE TABLE apps (
    id SERIAL PRIMARY KEY,
    app_id VARCHAR(255) UNIQUE NOT NULL,
    app_name VARCHAR(255) NOT NULL,
    app_type VARCHAR(255) NOT NULL,
    component VARCHAR(255) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the signups table
CREATE TABLE signups (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_no BIGINT NOT NULL,
    password VARCHAR(255) NOT NULL,
    confirm_password VARCHAR(255) NOT NULL
);

-- Create the metric_types table
CREATE TABLE metric_types (
    id SERIAL PRIMARY KEY,
    metric_type VARCHAR(255) NOT NULL,
    app_id INTEGER NOT NULL,
    component VARCHAR(255) NOT NULL
);

-- Create the daily_metrics table
CREATE TABLE daily_metrics (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    hours INTEGER NOT NULL,
    metric_type VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);

-- Create the weekly_metrics table
CREATE TABLE weekly_metrics (
    id SERIAL PRIMARY KEY,
    week DATE NOT NULL,
    hours INTEGER NOT NULL,
    metric_type VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);

-- Create the monthly_metrics table
CREATE TABLE monthly_metrics (
    id SERIAL PRIMARY KEY,
    month DATE NOT NULL,
    hours INTEGER NOT NULL,
    metric_type VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);

-- Create the yearly_metrics table
CREATE TABLE yearly_metrics (
    id SERIAL PRIMARY KEY,
    year DATE NOT NULL,
    hours INTEGER NOT NULL,
    metric_type VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);