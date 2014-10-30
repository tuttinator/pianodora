pianodora
=========

This project is designed to log events generated by Pianobar, and post notifications on song changes.

Pianobar allows for a shell command to be executed when an event occurs, such as song changes, liking songs, banning songs or changing stations.

Our use case is a Raspberry Pi, and previously we had a Ruby script posting Hipchat notifications. There is a noticeable delay between songs where the Ruby VM starts up, the script making an HTTP request and then terminates.

The idea behind this project is to have a client that boots quickly to communicate with a continuously running daemon. This daemon makes the actual API calls to your chat service of choice.

This is still a work in progress, and pre-alpha standard code.

It does not currently work.
