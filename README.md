# sizedir
Simple program to size a directory and all its contents

## Usage

    Usage of sizedir:
      -dot
            Include dot files
      -ext
            Aggregate by extension
      -files string
            Sets a file pattern to use (default "*")
      -path string
            Set the path to check (default ".")

## Example

Default path:

    $ sizedir
    1 folders, 5 files, 2212706 bytes, 2160KB, 2MB, 0GB
    Scanned in 205.199µs

Specifying a path:

    $ sizedir -path ~/Documents/
    3306 folders, 19351 files, 36943765320 bytes, 36077895KB, 35232MB, 34GB
    Scanned in 1.837590047s

Aggregating by extension:

    $ sizedir -path ~/Downloads -ext
    182 folders, 498 files, 2096761904 bytes, 2047619KB, 2000MB, 2GB
    Scanned in 17.881438ms

    Extens Files      Bytes      KB   MB GB   Avg KB
    ------ ----- ---------- ------- ---- -- --------
               2      20736      20    0  0    10368
    .JPG      18   72052754   70364   68  0  4002930
    .TXT       2       1309       1    0  0      654
    .cfg       4       2068       2    0  0      517
    .docx      1     781390     763    0  0   781390
    .gif       1      60260      58    0  0    60260
    .jpg      37   15375804   15015   14  0   415562
    .mp3      10  143870206  140498  137  0 14387020
    .pdf       6    7281693    7111    6  0  1213615
    .zip      15 1360105660 1328228 1297  1 90673710

(_example output shortened_)
