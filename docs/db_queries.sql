insert into products (uuid,name,brand) values (uuid(),'retailer1','brand1');
insert into products (uuid,name,brand) values (uuid(),'retailer2','brand2');
insert into products (uuid,name,brand) values (uuid(),'retailer3','brand3');
insert into products (uuid,name,brand) values (uuid(),'retailer4','brand4');
insert into products (uuid,name,brand) values (uuid(),'retailer5','brand5');
insert into products (uuid,name,brand) values (uuid(),'retailer6','brand6');


insert into retailers (uuid,name,email) values (uuid(),'retailer1','retailer1@gmail.com');
insert into retailers (uuid,name,email) values (uuid(),'retailer2','retailer2@gmail.com');
insert into retailers (uuid,name,email) values (uuid(),'retailer3','retailer3@gmail.com');
insert into retailers (uuid,name,email) values (uuid(),'retailer4','retailer4@gmail.com');
insert into retailers (uuid,name,email) values (uuid(),'retailer5','retailer5@gmail.com');
insert into retailers (uuid,name,email) values (uuid(),'retailer6','retailer6@gmail.com');


insert into products_retailers (uuid,product_id,retailer_id,price,quantity) values (uuid(),FLOOR(RAND()*(10-5+1)+1),FLOOR(RAND()*(10-5+1)+1),RAND()*(100)+1,FLOOR(RAND()*(100)));
