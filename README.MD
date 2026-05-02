Why Go uses []byte instead of string?

Because:

Efficient for I/O (files, network, APIs)
Flexible (you can send bytes directly)
Avoids unnecessary conversions

In real systems:

APIs send/receive bytes, not strings

json.Marshal() // for APIs (compact)
json.MarshalIndent() // for logs/debugging



