-- TASK 6 - SQL: Billing Software Report
-- You are working on a billing application and need to get a list of customers who have a positive
-- balance.
-- The result should have the following columns: iban | amount.
-- The result should be sorted in descending order by amount.
-- Note:
-- • Only customers whose balance is positive should be included in the report.
-- • For MS SQL Server users - in order to match the required format,
-- you can use FORMAT(amount, 'N') in your query.

-- ASSUMING THE STRUCTURE OF THE DATABASE IS :
CREATE TABLE customers (
    iban VARCHAR(255) PRIMARY KEY,  
    amount DECIMAL(15, 2)          
);

SELECT iban, amount
FROM customers
WHERE amount > 0
ORDER BY amount DESC;
