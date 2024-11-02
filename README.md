# Template-dir-generater 🛠️

Tool for automatic generation of template directories.

## Usage

```
./template-generater
```

Run the above with the following configuration.

```
.
├── config.yml
├── template-generater
└── templates
    ├── sample.json
    ├── sample.py
    └── sample.txt
```

Place template files to be copied in `templates`.
Configure the template directory in `config.yml`.

### Config Example

```
directory:
  name: template_dir
  path: ../
files:
  # Copy source file name
  - file: sample.txt
  # Output file name (Can be omitted)
    name: template.txt
  - file: sample.json
    name:
  - file: sample.py
```

If `name` is omitted, the output file name will be the same as `file`.
