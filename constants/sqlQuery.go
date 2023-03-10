package constants

const (
	AddFriendToExistingFriendsArray = `update public.user
		set    friends = (select array_agg(distinct e) from unnest(friends || ARRAY[$2]) e)
		where  not friends @> ARRAY[$2] and email = $1;`
	AddFriendToNullFriendsArray = `UPDATE public.user SET friends = ARRAY[$2]
		where  email = $1 and friends IS NULL;`
)
