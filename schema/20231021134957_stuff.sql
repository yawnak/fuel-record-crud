-- Modify "fuel_records" table
ALTER TABLE "fuel_records" DROP COLUMN "fuel_record_prev", ADD COLUMN "next_fuel_record_id" uuid NULL, ADD CONSTRAINT "fuel_records_fuel_records_next" FOREIGN KEY ("next_fuel_record_id") REFERENCES "fuel_records" ("fuel_record_id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Create index "fuel_records_next_fuel_record_id_key" to table: "fuel_records"
CREATE UNIQUE INDEX "fuel_records_next_fuel_record_id_key" ON "fuel_records" ("next_fuel_record_id");
