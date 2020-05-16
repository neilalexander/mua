# Matrix-Distributed Lua

This is a mad-science project to turn Matrix rooms into a kind of
distributed Lua environment. It's probably crazy and I don't really know
why anyone would want to, but here we are.

## Starting mdlc

The defaults will create a guest account on `matrix.org` and join you
to `#mdl-test:matrix.org`. 

It'll load the Lua contents from the `com.github.neilalexander.mdl.source`
event with the state key `main` and run it. 

Give it a go:

```
go run github.com/neilalexander/mdl/cmd/mdlc
```

If all goes well, you should get an interactive prompt in which you can
type Lua.

Try importing something from a different event state key:

```
$ go run github.com/neilalexander/mdl/cmd/mdlc
Using https://matrix.org
HELLO!
This code is running from the room state. Neat!
>> load("test")
I am going to set 'foo'!
OK, give it a go.
>> print(foo)
bar
>> 
```

## Encode some Lua

You can also create a Lua file, say `test.lua`:

```
$ cat test.lua
print("I am going to set 'foo'!")
foo = "bar"
print("OK, give it a go.")
```

and encode it with the `-encode test.lua` command line parameters:

```
$ go run github.com/neilalexander/mdl/cmd/mdlc -encode test.lua
{
  "type": 0,
  "source": "cHJpbnQoIkkgYW0gZ29pbmcgdG8gc2V0ICdmb28nISIpCmZvbyA9ICJiYXIiCnByaW50KCJPSywgZ2l2ZSBpdCBhIGdvLiIpCg"
}
```

You can now use your favourite client to send a custom event with this encoded blob as the content.

## Why would you do this?

No idea.