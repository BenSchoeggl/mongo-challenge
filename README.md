# Mongo Take Home Challenge

I solved the challenge with a simple depth first search through the JSON. Each leaf node in the JSON object needs to have a key in the result so you can DFS through to each leaf node, add an entry in the result for it and build its key as you return back up the stack.

Considerations:
* Before building this, I would check with stakeholders that the JSON objects aren't going to be deep enough to encounter a stack overflow error and build an iterative solution instead if this was the case. I left a comment about this in the method.
* Normally I stray away from recursive solutions in production code since they can be harder to grasp to outside readers, but this is such a simple DFS, I think it's acceptable in this case.

Testing strategy:
* I spent the large majority of my time building the testing infrastructure. I've worked extensively on [golden testing](https://ro-che.info/articles/2017-12-04-golden-tests) frameworks and I think it's an extremely useful development concept.
* With a golden testing workflow, developers can make changes to the code, and never have to make changes to testing code because unit test data is stored on files on disk instead of in code.
   * For example, in this challenge, if developers wanted to change the "." in keys to a "-", they could make that change in code, see the change in diff of the golden test files, and be confident the change worked without ever making testing code changes.
   * Developers can also add test cases without ever writing code because they can just add a file in the test data input directory.
   * PR reviewers benefit similarly, they only have to look at diff in unit test output files when reviewing a PR.


### Prerequisites

You should have go installed to this challenge

```
brew install golang
```

### Installing

```
go get -u github.com/BenSchoeggl/mongo-challenge
```

### Running

```
cat <json file to flatten> | go run github.com/BenSchoeggl/mongo-challenge
```

## Running the tests

```
go test github.com/BenSchoeggl/mongo-challenge/jsonutils
```