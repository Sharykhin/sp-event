var db = connect("mongodb://localhost/admin");

db.createUser(
    {
        user: "root",
        pwd: "root",
        roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
    }
)