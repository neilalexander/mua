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

... or by event ID:

```
>> load("$Ecwdeae6qipZA6Dk_fZKCa4NO-iYWs-uDK0E-2fzLdk")
I am going to set 'foo'!
OK, give it a go.
```

You can load an event from the room:

```
>> ev = event.new("$vnNJvIz7PStSIOE1NDEUxLmHUprPjrwrb556S-I6bjI")
>> print(ev:json())
{"state_key":"@neilalexander:dendrite.neilalexander.dev","sender":"@neilalexander:dendrite.neilalexander.dev","type":"m.room.member","origin_server_ts":1589645356974,"event_id":"$vnNJvIz7PStSIOE1NDEUxLmHUprPjrwrb556S-I6bjI","room_id":"!TdSVXZoEcLugVpglQn:matrix.org","unsigned":{},"content":{"avatar_url":"","displayname":"neilalexander (Dendrite)","membership":"join"}}
>> print(ev:type())
m.room.member
>> print(ev:content())
{"avatar_url":"","displayname":"neilalexander (Dendrite)","membership":"join"}
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

You can now use your favourite client to send a custom event with this
encoded blob as the content.

## Is this safe?

You can't access the filesystem, sockets, system calls or such within
the Lua environment. It should be reasonably self-contained as a result,
although this may also make it a bit harder to do anything useful with.

## Why would you do this?

No idea.