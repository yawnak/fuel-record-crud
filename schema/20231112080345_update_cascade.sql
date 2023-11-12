-- Modify "fuel_records" table
ALTER TABLE "fuel_records" DROP CONSTRAINT "fuel_records_cars_fuel_records", ADD CONSTRAINT "fuel_records_cars_fuel_records" FOREIGN KEY ("car_id") REFERENCES "cars" ("car_id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "odometer_records" table
ALTER TABLE "odometer_records" DROP CONSTRAINT "odometer_records_cars_odometer_records", ADD CONSTRAINT "odometer_records_cars_odometer_records" FOREIGN KEY ("car_id") REFERENCES "cars" ("car_id") ON UPDATE NO ACTION ON DELETE CASCADE;
