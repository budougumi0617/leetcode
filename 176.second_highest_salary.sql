# https://leetcode.com/problems/second-highest-salary/
SELECT MAX(Salary) AS SecondHighestSalary FROM Employee
WHERE Salary NOT IN (
	SELECT MAX(Salary) FROM Employee
);
