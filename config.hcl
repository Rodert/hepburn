log {
    log_level = "info"
    log_path = "./web.log"
}

web {
    address = "0.0.0.0:20001"
}

mysql {
    dns = "root:root@tcp(127.0.0.1:3306)/hepburn?charset=utf8mb4&parseTime=True&loc=Local"
}

