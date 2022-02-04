# Video Ripping Tools: SortByFolder

For convenience, I recently started ripping my DVDs to digital videos, using Handbrake to perform the rip, and MetaX to add appropriate metadata so they can be added to iTunes and not look out of place.

For movies this is fairly straightforward. However, for TV shows, juggling all the output files and simplifying running them through MetaX took a bit of thought. Ultimately I came up with a couple of utilities to help reduce the pain of ripping a full season (or more) of a show.

These were originally written in Perl, but it seemed worthwhile rewriting them in Go and making them available for anyone else who might want them.

## CreateFolders

See the [CreateFolders](https://github.com/PJSoftware/go-create-folders) repo.

## SortByFolder

Once all disks of a season have been ripped into the folder structure described above, SortByFolder.exe will rename them all to the format "S#E##". This makes labelling them via MetaX much easier, since it recognises this naming convention.

## History

Both of these tools were initially developed in my `Command-Line-Util` repo -- an approach guided by my SVN experience. However, multiple projects in a single repo is not the `git` way, so I've finally decided to split them.
