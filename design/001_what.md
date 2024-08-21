## Project

Overall I wanted to build a project that looks at two things:

 - sync
 - async

We all know what this means inherently but I wanted to define what this impact has for consumers and the experience this means for how they interact with the service. The assumptions they make and how we can make this more visible.

- What does a 200 response actually mean?
- Latency requirements and mapping?
- Guarantees and delivery?
- Patterns for ensuring messages are persisted without loss, are there are any proofs we can use?

So this is what we came up


-> Sync Processor
-> Async Processor

We will design these that we could replace the specific streaming service or datastore at a later point.


