select
    employee.name
from
    employee
where
    employee.id in (
        select
            managerid
        from
            employee
        group by
            managerid
        having
            count(managerid) > 4
    );