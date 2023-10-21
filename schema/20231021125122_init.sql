-- Create "cars" table
CREATE TABLE "cars" (
    "car_id" uuid NOT NULL,
    "make" character varying NOT NULL,
    "model" character varying NOT NULL,
    "year" integer NOT NULL,
    PRIMARY KEY ("car_id")
);
-- Create "fuel_records" table
CREATE TABLE "fuel_records" (
    "fuel_record_id" uuid NOT NULL,
    "current_fuel_liters" double precision NOT NULL,
    "difference" double precision NOT NULL,
    "created_at" timestamptz NOT NULL,
    "car_fuel_records" uuid NOT NULL,
    "fuel_record_prev" uuid NULL,
    PRIMARY KEY ("fuel_record_id"),
    CONSTRAINT "fuel_records_cars_fuel_records" FOREIGN KEY ("car_fuel_records") REFERENCES "cars" ("car_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "fuel_records_fuel_records_prev" FOREIGN KEY ("fuel_record_prev") REFERENCES "fuel_records" ("fuel_record_id") ON UPDATE NO ACTION ON DELETE
    SET NULL
);
-- Create index "fuel_records_fuel_record_prev_key" to table: "fuel_records"
CREATE UNIQUE INDEX "fuel_records_fuel_record_prev_key" ON "fuel_records" ("fuel_record_prev");
-- Create "odometer_records" table
CREATE TABLE "odometer_records" (
    "odometer_record_id" uuid NOT NULL,
    "current_fuel_liters" double precision NOT NULL,
    "difference" double precision NOT NULL,
    "created_at" timestamptz NOT NULL,
    "car_odometer_records" uuid NOT NULL,
    "odometer_record_prev" uuid NULL,
    PRIMARY KEY ("odometer_record_id"),
    CONSTRAINT "odometer_records_cars_odometer_records" FOREIGN KEY ("car_odometer_records") REFERENCES "cars" ("car_id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    CONSTRAINT "odometer_records_odometer_records_prev" FOREIGN KEY ("odometer_record_prev") REFERENCES "odometer_records" ("odometer_record_id") ON UPDATE NO ACTION ON DELETE
    SET NULL
);
-- Create index "odometer_records_odometer_record_prev_key" to table: "odometer_records"
CREATE UNIQUE INDEX "odometer_records_odometer_record_prev_key" ON "odometer_records" ("odometer_record_prev");