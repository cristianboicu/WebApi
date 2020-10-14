# WebApi

->Routes

ip = localhost

port = 8080

https://ip:port/api/v1/items    ->  HTTP.GET     (get all)

https://ip:port/api/v1/item     ->  HTTP.POST    (post)

https://ip:port/api/v1/item/id  ->  HTTP.PUT     (put(update) by id)

https://ip:port/api/v1/item/id  ->  HTTP.DELETE  (delete by id)


->Item details

Task {
int ID;
string Name;
string Details;
time Date;
bool Done;
}
