-- Modify "fuel_records" table
ALTER TABLE "fuel_records" DROP CONSTRAINT "fuel_records_cars_fuel_records", DROP COLUMN "car_fuel_records", ADD COLUMN "car_id" uuid NOT NULL, ADD CONSTRAINT "fuel_records_cars_fuel_records" FOREIGN KEY ("car_id") REFERENCES "cars" ("car_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
