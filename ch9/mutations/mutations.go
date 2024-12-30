package mutations

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	err = errors.New("not found")

	value, ok := users[name]

	if !ok {
		return false, err
	}

	if ok && value.scheduledForDeletion == false {
		return false, nil
	}

	delete(users, name)

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

/*

Insert an element
m[key] = elem
Copy icon
Get an element
elem = m[key]
Copy icon
Delete an element
delete(m, key)
Copy icon
Check if a key exists
elem, ok := m[key]
Copy icon
If key is in m, then ok is true and elem is the value as expected.

If key is not in the map, then ok is false and elem is the zero value for the map's element type.

Assignment
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
Note on passing maps
Like slices, maps are also passed by reference into functions.
This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.

*/
