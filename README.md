# film-fetch

Utility written in Go to provide a GUI to retrieve (fetch) films from a server. Very much a work in progress and I'm using the project to learn Go as I.. go.. but a very slapdash roadmap of what I want to achieve is below

## Objective
I have a media server with (legal, of course) videos that my Dad watches through Plex. Unfortunately my parents Internet connection isn't good enough to maintain a connection without buffering every few minutes. My parents can't legally download videos themselves without riddling their computer with junk so I figured I'd enable them to copy individual films to their computer via an easy to use GUI. I can see use of this outside of my situation however, so I figured I'd open source anything I manage to get done. 

I'm aware there are probably better ways of doing this, but I wanted a clearly defined project to learn Go against.

## Roadmap

##### Phase 1

 - Terminal based
 - Prompt user for server details, then make a connection
 - Prompt user for local download location
 - Recursively list films in a parent directory
 - Allow a film to be selected and then retrieved to the local machine via SSH

##### Phase 2
  - Executable which opens default browser and navigates to web app
  - Implement an API to show poster and proper title for any films
