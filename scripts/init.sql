CREATE TABLE "public.queries" (
                                  "id" serial NOT NULL,
                                  "query" TEXT NOT NULL,
                                  "filter_price_low" integer,
                                  "filter_price_high" integer,
                                  "filter_city" TEXT,
                                  CONSTRAINT "queries_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "public.items" (
                                "id" serial NOT NULL,
                                "name" TEXT NOT NULL,
                                "description" TEXT NOT NULL,
                                "price" integer NOT NULL,
                                "img" TEXT NOT NULL,
                                "location" TEXT NOT NULL,
                                "url" TEXT NOT NULL,
                                "blacklisted" BOOLEAN NOT NULL,
                                "date_added" DATE NOT NULL,
                                "last_updated" TIMESTAMP NOT NULL,
                                "query_id" integer NOT NULL,
                                CONSTRAINT "items_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );




ALTER TABLE "items" ADD CONSTRAINT "items_fk0" FOREIGN KEY ("query_id") REFERENCES "queries"("id");



