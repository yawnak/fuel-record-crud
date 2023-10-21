-- Modify "fuel_records" table
ALTER TABLE "fuel_records" DROP CONSTRAINT "fuel_records_fuel_records_next", ADD CONSTRAINT "fuel_records_fuel_records_prev" FOREIGN KEY ("next_fuel_record_id") REFERENCES "fuel_records" ("fuel_record_id") ON UPDATE NO ACTION ON DELETE SET NULL;
