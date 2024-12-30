package slices

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {

	err := errors.New("unsupported plan")

	if plan == planPro {
		return messages[:], nil
	}
	if plan == planFree {
		return messages[:2], nil
	}

	return nil, err

}

/*

9 times out of 100 you will use a slice instead of an array when working with ordered lists.

Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

The zero value of slice is nil.

Slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:

primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}

The syntax is:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

Where lowIndex is inclusive and highIndex is exclusive.

lowIndex, highIndex, or both can be omitted to use the entire array on that side of the colon.

Assignment
Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.

*/

//More on slices

/*
Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, any changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer (we'll cover pointers later) to the underlying array. A Read function can therefore accept a slice argument rather than a pointer and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read() method of the File type in package os:

Referenced from Effective Go

func (f *File) Read(buf []byte) (n int, err error)

*/

// Make function

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

/*
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

/ func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

/ the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}

Notice the square brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.

Length
The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3
Copy icon
Capacity
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3

Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice because it will automatically grow as needed.

Assignment
We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.

*/
