CREATE TABLE users (
    userid INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100) ,
    yob YEAR,
    address VARCHAR(255),
    phone INT(10)
);

CREATE TABLE accounts (
    accountid INT PRIMARY KEY AUTO_INCREMENT,
    userid INT,
    account_type VARCHAR(50),
    balance DECIMAL(10, 2) DEFAULT 0.00,
    FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE transactions (
    transaction_id INT PRIMARY KEY AUTO_INCREMENT,
    accountid INT,
    amount DECIMAL(10, 2),
    transaction_type VARCHAR(20) , 
    FOREIGN KEY (accountid) REFERENCES accounts(accountid) ON DELETE CASCADE
);

CREATE TABLE atms (
    atm_id INT PRIMARY KEY AUTO_INCREMENT,
    location_name VARCHAR(100)
);
