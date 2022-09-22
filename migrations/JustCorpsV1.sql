CREATE TABLE "User" (
  "user_id" int PRIMARY KEY,
  "fname" string,
  "lname" string,
  "email" string,
  "password" string,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Favorite" (
  "favorite_id" int PRIMARY KEY,
  "user_id" [fk],
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Product" (
  "product_id" int PRIMARY KEY,
  "url" string,
  "description" string,
  "created_at" datetime,
  "last_modified" datetime,
  "user_id" int,
  "img_id" string
);

CREATE TABLE "ProductImage" (
  "image_id" int PRIMARY KEY,
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Image" (
  "image_id" int PRIMARY KEY,
  "image_location" string,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Certification" (
  "certification_id" int PRIMARY KEY,
  "certifying_company" string,
  "cert_name" string,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Company" (
  "company_id" int PRIMARY KEY,
  "department" string,
  "url" string,
  "description" text,
  "created_at" datetime,
  "last_modified" datetime,
  "user_id" int,
  "verified" string,
  "img_id" string
);

CREATE TABLE "ProductCertification" (
  "productcert_id" int PRIMARY KEY,
  "certification_id" int,
  "product_id" int,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "CompanyCertification" (
  "companycert_id" int PRIMARY KEY,
  "certification_id" int,
  "company_id" int,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "CompanyImage" (
  "image_id" int PRIMARY KEY,
  "company_id_id" [fk],
  "image_location" string,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "ProductDepartment" (
  "department_id" int PRIMARY KEY,
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Department" (
  "department_id" int PRIMARY KEY,
  "title" string,
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "ProductCategory" (
  "category_id" int PRIMARY KEY,
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Category" (
  "category_id" int PRIMARY KEY,
  "title" string,
  "created_at" datetime,
  "last_modified" datetime,
  "department_id" [fk]
);

CREATE TABLE "ProductType" (
  "type_id" int PRIMARY KEY,
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Type" (
  "type_id" int PRIMARY KEY,
  "title" string,
  "created_at" datetime,
  "last_modified" datetime,
  "category_id" [fk]
);

CREATE TABLE "ProductStyle" (
  "style_id" int PRIMARY KEY,
  "product_id" [fk],
  "created_at" datetime,
  "last_modified" datetime
);

CREATE TABLE "Style" (
  "style_id" int PRIMARY KEY,
  "title" string,
  "created_at" datetime,
  "last_modified" datetime,
  "type_id" [fk]
);

ALTER TABLE "Product" ADD FOREIGN KEY ("product_id") REFERENCES "ProductCertification" ("product_id");

ALTER TABLE "Certification" ADD FOREIGN KEY ("certification_id") REFERENCES "ProductCertification" ("certification_id");

ALTER TABLE "Favorite" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("product_id");

ALTER TABLE "Favorite" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "Company" ADD FOREIGN KEY ("company_id") REFERENCES "CompanyImage" ("company_id_id");

ALTER TABLE "Product" ADD FOREIGN KEY ("product_id") REFERENCES "ProductImage" ("product_id");

ALTER TABLE "Image" ADD FOREIGN KEY ("image_id") REFERENCES "CompanyImage" ("image_id");

ALTER TABLE "Image" ADD FOREIGN KEY ("image_id") REFERENCES "ProductImage" ("image_id");

ALTER TABLE "ProductType" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("product_id");

ALTER TABLE "ProductDepartment" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("product_id");

ALTER TABLE "Department" ADD FOREIGN KEY ("department_id") REFERENCES "ProductDepartment" ("department_id");

ALTER TABLE "ProductCategory" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("product_id");

ALTER TABLE "Category" ADD FOREIGN KEY ("category_id") REFERENCES "ProductCategory" ("category_id");

ALTER TABLE "ProductStyle" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("product_id");

ALTER TABLE "ProductStyle" ADD FOREIGN KEY ("style_id") REFERENCES "Style" ("style_id");

ALTER TABLE "ProductType" ADD FOREIGN KEY ("type_id") REFERENCES "Type" ("type_id");

ALTER TABLE "Company" ADD FOREIGN KEY ("company_id") REFERENCES "CompanyCertification" ("company_id");

ALTER TABLE "Certification" ADD FOREIGN KEY ("certification_id") REFERENCES "CompanyCertification" ("certification_id");
