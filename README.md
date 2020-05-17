# Mua

This is a mad-science project to turn Matrix rooms into a kind of
distributed Lua environment. It's probably crazy and I don't really know
why anyone would want to, but here we are.

## Starting muac

The defaults will create a guest account on `matrix.org` and join you
to `#mua-test:matrix.org`. 

It'll load the Lua contents from the `mua.source` event with the state
key `main` and run it. 

Give it a go:

```
go run github.com/neilalexander/mua/cmd/muac
```

If all goes well, you should get an interactive prompt in which you can
type Lua.

## Loading Lua from the room

Using `importstate()`, you can import something from a specific event state key:

```
$ go run github.com/neilalexander/mua/cmd/muac
Using https://matrix.org
HELLO!
This code is running from the room state. Neat!
>> importstate("!TdSVXZoEcLugVpglQn:matrix.org", "mua.source", "main")
HELLO!
This is running from bytecodes stored in the room state.
>> 
```

... or the same event by event ID instead:

```
>> importevent("!TdSVXZoEcLugVpglQn:matrix.org", "$Z_brwOSlxvtgp24_g_N4QmzBQtSEtl7rI4IQIp02S94")
HELLO!
This is running from bytecodes stored in the room state.
```

## Loading room events

You can load an event from the room:

```
>> ev = event.new("!TdSVXZoEcLugVpglQn:matrix.org", "$Z_brwOSlxvtgp24_g_N4QmzBQtSEtl7rI4IQIp02S94")
>> print(ev:json())
{"state_key":"main","sender":"@neilalexander:matrix.org","type":"mua.source","origin_server_ts":1589709947669,"event_id":"$Z_brwOSlxvtgp24_g_N4QmzBQtSEtl7rI4IQIp02S94","room_id":"!TdSVXZoEcLugVpglQn:matrix.org","unsigned":{"age":1079280},"content":{"source":"cHJpbnQoIkhFTExPISIpCnByaW50KCJUaGlzIGlzIHJ1bm5pbmcgZnJvbSBieXRlY29kZXMgc3RvcmVkIGluIHRoZSByb29tIHN0YXRlLiIpCmxvYWRlZCA9IHRydWUK","type":0}}
>> print(ev:type())
mua.source
```

## Run local scripts

Run a local script, using the room state as the preloaded state:

```
$ cat getevent.lua
ev = event.new("$vnNJvIz7PStSIOE1NDEUxLmHUprPjrwrb556S-I6bjI")
print(ev:type())
print(ev:content())

$ go run github.com/neilalexander/mua/cmd/muac test.lua
HELLO!
This code is running from the room state. Neat!
m.room.member
{"avatar_url":"","displayname":"neilalexander (Dendrite)","membership":"join"}
```

You can `load()` in your script as normal in order to import and run Lua from the room.

## Encode some Lua

You can also take a compiled Lua file from `luac`, say `luac.out`, and
encode it with the `-encode luac.out` command line parameters:

```
$ go run github.com/neilalexander/mua/cmd/muac -encode test.lua
{
  "type": 0,
  "source": "G0x1YVIAAQQIBAgAGZMNChoKAAAAAAAAAAAAAQQVAAAABgBAAEFAAACBgAAAHUCAAQbAQABBQAAAgQABAMFAAQAdQAACBoBBAEHAAQCGAEIAmwAAABeAAICBQAIAm0AAABcAAICBgAIAVoCAAB1AAAEfAIAACwAAAAQMAAAAAAAAAGltcG9ydGV2ZW50AAQfAAAAAAAAACFUZFNWWFpvRWNMdWdWcGdsUW46bWF0cml4Lm9yZwAELQAAAAAAAAAkWl9icndPU2x4dnRncDI0X2dfTjRRbXpCUXRTRXRsN3JJNElRSXAwMlM5NAAEDAAAAAAAAABpbXBvcnRzdGF0ZQAECwAAAAAAAABtdWEuc291cmNlAAQFAAAAAAAAAG1haW4ABAYAAAAAAAAAcHJpbnQABAkAAAAAAAAATG9hZGVkPyAABAcAAAAAAAAAbG9hZGVkAAQEAAAAAAAAAHllcwAEAwAAAAAAAABubwAAAAAAAQAAAAEACgAAAAAAAABAdGVzdC5sdWEAFQAAAAEAAAABAAAAAQAAAAEAAAACAAAAAgAAAAIAAAACAAAAAgAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAQAAAAEAAAABAAAAAAAAAABAAAABQAAAAAAAABfRU5WAA"
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