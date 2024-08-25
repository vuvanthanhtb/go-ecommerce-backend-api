## DUMP Database
```docker exec -it mysql mysqldump -uroot -pThanh2024 --databases shopdevgo --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/shopdevgo.sql```
