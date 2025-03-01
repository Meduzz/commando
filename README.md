# commando
Perhaps a great lib when coding nude?

## Goals

* Define the commands and flags somewhere else to existing logic (within boundaries).
* Simplify all helper/starters with a hyper extensible DSL.

### Commands and flags defined somewhere else with local logic

This could be used to build clis with permission management or where similar logic is executed in a similar fashion.

Ex:

A command that points to a file that is then read and uploaded somewhere.

Another command does the same thing but uploads it somewhere else.

**How?**

Some kind of json/config structure that can be turned into commands somewhat easily.

Ex:

I define a set of handlers in my binary but allow someone else to define the commands that then execute them.

**Todo**

* Add some kind of handler factory abstraction. Prolly by extending the command definition with a setting field of type json.RawMessage.

### Simplify all starters with extensible DSL

With cobra commands there's some plumbing, I want to get rid of it.

**How?**

A DSL, thathat generate straight into the libs registry.

**Todo**

* Not sure this goal was achieved yet. Time will tell.
* The read flags part (ExecuteCommand) also needs a converter abstraction most likely.