- Active: 
    - Zip was overly complex. I want this to be simple. 
    - FindMovie -> Calls OMDB
    - AddMovies  (POST)
    - OverwriteMovies (Patch)
    - RemoveMovies (DELETE)
    - SearchMovies (GET)
    - Review use of []interface{} for movie data; consider alternatives for flexibility and type safety.

- Vibe Safe
    - Cannot store passwords inside directory

- AWS Function. AWS DB. 

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