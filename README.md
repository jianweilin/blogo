# Blogo

## Synopsis

Blogo is an util to generate blog template, you give a blog title, and generates title and tile and category intelligent by your title.
This is great when you write a blog and don't want to specific an category.
Those features are provided for now:

* Support hexo blog templates;
* Support automatically generate different language templates, if you are writing English, **blogo** will generate English summary
if you are writing Chinese, a Chinese version will be generate;

Here is the screenshot of blogo.
![PicName](http://ofwzcunzi.bkt.clouddn.com/nYQZeGXPh6dFj1MS.png)


## Usage

blogo is very simple to use, just give your title, all the things will be done!

```
mv blogo /usr/bin
blogo -title "your title of blog"
```

That's it, and your *.md file will generates in your current directory.
However, you can also specific your directory by given `-path` option. For more usage, type:
```
blogo -h
```

## Future Work

Maybe this tool will becomes more intelligent by applying AI classifier and divided your article into different categories. This could be
a interesting work to explore AI in Golang!

## Terminate Goal

Well, the final goal would be, I think, let blogo automatically generates a blog by your title given. This souds realy good, not only in blog, in the future, the writer, AI can helps them to structure a article in seconds!


