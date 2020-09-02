# rename

Cross compatible rename cli, similar to sed.

## Usage

```
rename [-hn] s/pattern/replace/ path
  -h    Print help
  -n    Dry run, will not actually rename any files
```

## Example

```shell
tree
.
├── 1.yaml
├── 2.yaml
├── 3.yaml
├── 4.yaml
└── 5.yaml

rename -n 's/(\d)\.yaml/$1.yml/' .                                                                                                                                                                                                16:00:09
1.yaml   -->   1.yml
2.yaml   -->   2.yml
3.yaml   -->   3.yml
4.yaml   -->   4.yml
5.yaml   -->   5.yml
```

