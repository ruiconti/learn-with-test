# Note

## TDD

> "When to use iterative development? You should use iterative development only on projects that you want to succeed."
> Martin Fowler

1. Always split your problem in the thinnest possible "vertical slice". In a way that you continually make progress on **working** software. This is a muscle that needs training.
2. Once you have some working software it should be easier to iterate with small steps until you arrive at the software you need.
2. When we deal with IO that take some time to process, mocking (or faking) comes a long way!

## Mocking

Takeaways

1. We can use dependency injection to insert mocks to our functions.
2. Design fakes in a way that we can spy if they were called and used i.e. behaviored as expected. 

Rules of thumb

- Refactor only **public behavior**. Not implementation details.
- I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design
- Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. Be sure you actually care about these details if you're going to spy on them
