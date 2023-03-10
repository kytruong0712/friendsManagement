package constants

const (
	GetAllUsers                     = `select id, name, email, friends, subscribe, created_at, updated_at from public.user order by id`
	GetUser                         = `select id, name, email, friends, subscribe, created_at, updated_at from public.user where email = $1`
	AddFriendToExistingFriendsArray = `update public.user
										set    friends = (select array_agg(distinct e) from unnest(friends || ARRAY[$2]) e)
										where  not friends @> ARRAY[$2] and email = $1;`
	AddFriendToNullFriendsArray          = `UPDATE public.user SET friends = ARRAY[$2] where  email = $1 and friends IS NULL;`
	AddSubscribeToExistingSubscribeArray = `update public.user
											set    subscribe = (select array_agg(distinct e) from unnest(subscribe || ARRAY[$2]) e)
											where  not subscribe @> ARRAY[$2] and email = $1;`
	AddSubscribeToNullSubscribeArray = `UPDATE public.user SET subscribe = ARRAY[$2] where  email = $1 and subscribe IS NULL;`
)
