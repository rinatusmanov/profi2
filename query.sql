select work_work_id from (
	select work_work_id,count(work_work_id) as count_of_works from order_works where order_order_id in (
		select order_id from orders where order_room_id in (
			select rooms.room_id from rooms
			join room_types on
			 	room_type_name != 'Малое помещение' and
			 	room_volume > room_type_from_volume and
			 	room_volume < room_type_to_volume
	 	)
	)
	group by work_work_id
) as count_of_works_query
where count_of_works >= 6;


select * from employees where employee_id not in (
	select employee_id from order_gantts where order_gantt_from_date > '2021-03-14 00:00:00' and order_gantt_to_date < '2021-03-14 23:59:59'
);
