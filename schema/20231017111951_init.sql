-- Create the 'cars' table
CREATE TABLE cars (
    car_id UUID PRIMARY KEY,
    make VARCHAR,
    model VARCHAR,
    year INT8,
    currentFuelRecordId UUID,
    currentOdometerRecordId UUID
);

-- Create the 'fuel_records' table
CREATE TABLE fuel_records (
    id UUID PRIMARY KEY,
    previousRecordId UUID,
    nextRecordId UUID,
    current_fuel FLOAT8,
    difference FLOAT8,
    created_at TIMESTAMPTZ
);

-- Create the 'odometer_records' table
CREATE TABLE odometer_records (
    id UUID PRIMARY KEY,
    previousRecordId UUID,
    nextRecordId UUID,
    currentKilometers INT8 CHECK (currentKilometers >= 0),
    difference INT8,
    created_at TIMESTAMPTZ
);
