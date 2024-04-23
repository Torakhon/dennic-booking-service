CREATE TABLE "patients"(
                           "id" UUID NOT NULL,
                           "first_name" VARCHAR(50) NOT NULL,
                           "last_name" VARCHAR(50) NOT NULL,
                           "birth_date" DATE NOT NULL,
                           "gender" VARCHAR(255) NOT NULL CHECK ("gender" IN ('Male', 'Female', 'Other')),
                           "blood_group" VARCHAR(255) NOT NULL CHECK ("blood_group" IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-')),
                           "phone_number" VARCHAR(15) NOT NULL,
                           "city" VARCHAR(50) NOT NULL,
                           "country" VARCHAR(50) NOT NULL,
                           "patient_problem" TEXT NOT NULL,
                           "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                           "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

ALTER TABLE "patients" ADD PRIMARY KEY("id");

CREATE TABLE "archive"(
                          "id" SERIAL NOT NULL,
                          "doctor_availability_id" INTEGER NOT NULL,
                          "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "patient_problem" TEXT NOT NULL,
                          "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                          "status" VARCHAR(255) NOT NULL CHECK ("status" IN ('Active', 'Inactive')),
                          "payment_type" VARCHAR(255) NOT NULL CHECK ("payment_type" IN ('Cash', 'Card', 'Insurance')),
                          "payment_amount" DOUBLE PRECISION NOT NULL,
                          "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                          "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

ALTER TABLE "archive" ADD PRIMARY KEY("id");

CREATE TABLE "doctor_notes"(
                               "id" SERIAL NOT NULL,
                               "appointment_id" INTEGER NOT NULL,
                               "doctor_id" UUID NOT NULL,
                               "patient_id" UUID NOT NULL,
                               "prescription" TEXT NOT NULL,
                               "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                               "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE,
                               CONSTRAINT "doctor_notes_appointment_id_foreign" FOREIGN KEY("appointment_id") REFERENCES "archive"("id"),
                               CONSTRAINT "doctor_notes_patient_id_foreign" FOREIGN KEY("patient_id") REFERENCES "patients"("id")
);

ALTER TABLE "doctor_notes" ADD PRIMARY KEY("id");

ALTER TABLE "booked_appointments" ADD PRIMARY KEY("id");

CREATE TABLE "booked_appointments"(
                                      "id" SERIAL NOT NULL,
                                      "department_id" UUID NOT NULL,
                                      "doctor_id" UUID NULL,
                                      "patient_id" UUID NOT NULL,
                                      "appointment_date" DATE NOT NULL,
                                      "appointment_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "duration" BIGINT NOT NULL,
                                      "key" VARCHAR(20) NOT NULL,
                                      "expires_at" BIGINT NOT NULL,
                                      "patient_status" BOOLEAN NOT NULL DEFAULT TRUE,
                                      "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                                      "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

CREATE TABLE "doctor_availability"(
                                      "id" SERIAL NOT NULL,
                                      "department_id" UUID NOT NULL,
                                      "doctor_id" UUID NOT NULL,
                                      "doctor_date" DATE NOT NULL,
                                      "start_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "end_time" TIME(0) WITHOUT TIME ZONE NOT NULL,
                                      "status" VARCHAR(255) NOT NULL CHECK ("status" IN ('Available', 'Unavailable')),
                                      "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE ,
                                      "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE
);

ALTER TABLE "doctor_availability" ADD PRIMARY KEY("id");

ALTER TABLE "archive" ADD CONSTRAINT "archive_doctor_availability_id_foreign" FOREIGN KEY("doctor_availability_id") REFERENCES "doctor_availability"("id");
