- Active: Zip search results and make repos async
    - Current issue is that I can't seem to get the results into the json array. See HandleSeach_Two and GetOmdbMovies
    - https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
    - Current thought is maybe you are serializing it into a dictionary? Try key value of "Search"? Research further into code base for example of dealing with json payloads. Save? 

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