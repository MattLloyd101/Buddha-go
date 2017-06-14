# Buddha-go

A Buddhabrot renderer in Go.

# Todo List

1. Write a priority heat map pass to pre-generate what areas of the image are interesting. This should be pickled and re-used for future passes. It should be low resolution and scaled up to speed up high resolution passes. The heatmap should give us a method of then distributing the iterations such that we hit more interesting areas more often than boring areas.
2. A Tiling renderer to handle output larger than 4gb.
3. Loop detector to cut down on the burn in on strange loops.