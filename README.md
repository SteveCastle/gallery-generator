# Gallery Generator

A simple GO program to serve a gallery of images from a folder of images.

## Build
```bash
go build .
```

## Usage

Place all of your images in the `images` subdirectory in the same directory as the executable. Then run the executable. The gallery will be served on the port provided

```
./gallery-generator -t="Gallery Title" -p=80 -r=.5 -c=4
```

## Options

        -t="Gallery Title"  Title of the gallery. Will be used for the page title and the h1 header.
        -p=8888             Percentage of the image to use.
        -r=1                Ratio of the grid images. 1 is square, 2 is 2:1, 0.5 is 1:2
        -c=6                Number of columns to use.
