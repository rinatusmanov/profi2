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
where count_of_works >= 6
