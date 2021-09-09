CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "examination_requests" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"vin" TEXT NOT NULL,
"request_date" DATETIME NOT NULL,
"request_number" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE sqlite_sequence(name,seq);
