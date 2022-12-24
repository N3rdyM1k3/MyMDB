- Active: Zip search results and make repos async
    - We got things working, but we aren't done. Next steps:
        - Search owned movies (by title?)
        - Verify merge behavior
        - Maybe provide a single API? Idk... the merge function feels like buisness logic to me, so maybe keep it out of the repo

- Split repo directory into two packages to hopefully create repositories/mongo/ and repositories/omdb/ 
- Individual Get: Check Mongo First? Check both and zip? 
- Save: Add isOwned
- Save: should merge??? 
- Getting individual results from omdb returns much more detail
- Make repo methods async

- Organize a bit more: create additional repos, maybe create helper to write data? 
- Zip search results with results from db 
- Add tests? 
- Go deployed to GCP hello world 


- Polish ideas
    - Create actual structs with names even if they are empty interfaces? This would *feel* more strongly typed? I think it just *feels* more go based on examples I've seen. 