# speedreader

### Usage

```bash
cat your_doc | sread -w 200
```

or 

```bash
sread my_file.txt --no-highlight
```

| flag | short | default | purpose |
| --- | --- | --- | --- |
| --words-per-minute | -w | 300 | the speed the words are shown |
| --highlight-color | -c | #FF0088 | color of the 'center' character |
| --no-highlight | -n | false | don't highlight the 'center' character |
