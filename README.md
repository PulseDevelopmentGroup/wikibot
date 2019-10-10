# WikiBot
A bot to find the shortest path between two links on Wikipedia (Pretty much brute forcing it).

## Design Ideaa

- [ ] Initially command line input: `./wikibot -start https://some.wiki.page -end https://some.other.wiki.page`
- [ ] Starts two threads and starts at both ends, searching all links for a connection
- [ ] If no connections are found, the first link of each page is used to find connections
- [ ] Some sort of map or data store will be used to contain the history of pages and their connections (so that backreference can be performed)
- [ ] Will initially run entirely in memory (RIP garbage collection)
