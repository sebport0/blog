-- 1. Create the table.
create table internet_speed_2022 (
    country text,
    broadband text,
    mobile text
);

-- 2. Import data from the original CSV source.
copy internet_speed_2022 (country, broadband, mobile)
from '/tmp/internet_speed_2022.csv'
delimiter ','
csv header;

-- 3. Check that the import worked as expected.
select * from internet_speed_2022 limit 10;

-- 4. Compute percentiles for the broadband column.
select
    percentile_cont(.5) 
    within group (order by broadband::numeric) as p50
from internet_speed_2022;

select
    percentile_cont(.25) 
    within group (order by broadband::numeric) as p25
from internet_speed_2022;

select
    percentile_cont(array[0.25, 0.5, 0.75]) 
    within group (order by broadband::numeric) as percentiles
from internet_speed_2022;

select
    unnest(
        percentile_cont(array[0.25, 0.5, 0.75]) 
        within group (order by broadband::numeric)
    ) as percentiles
from internet_speed_2022;
