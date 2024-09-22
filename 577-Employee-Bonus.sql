SELECT name, bonus FROM Employee
LEFT OUTER JOIN Bonus ON Bonus.empId = Employee.empId
WHERE bonus < 1000 OR Bonus.empId is null