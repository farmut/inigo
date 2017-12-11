# Inigo

Golang written `INI` files **parser** with commented (_disabled_) options support.

           o8o               o8o
           `"'               `"'
          oooo  ooo. .oo.   oooo   .oooooooo  .ooooo.
          `888  `888P"Y88b  `888  888' `88b  d88' `88b
           888   888   888   888  888   888  888   888
           888   888   888   888  `88bod8P'  888   888
          o888o o888o o888o o888o `8oooooo.  `Y8bod8P'
          ___.ini files parser___ d"     YD
                                  "Y88888P'

1. **Description**
    * _App_
    * _Format spec_
2. **Install**
3. **Usage**
5. **Contrib**
6. **TODO**

## Description

### App

### Ini file format

As there is not strict specification for `INI` format, so **Inigo** is based on common description of it in 
[Wikipedia](https://en.wikipedia.org/wiki/INI_file). Here is the basic notes from there.

#### Keys or Parameters

The basic element contained in an `INI` file is the **key** or property. Every **key** has a _name_ and a 
_value_, delimited by an **equals sign** `=`. The _name_ appears to the left of the equals sign.

**Note:** In the Windows implementation the key cannot contain the characters equal sign `=` or semicolon
 `;` as these are reserved characters. The _value_ can contain **any character**.

    name=value

#### Sections

**Keys** may (but need not) be grouped into arbitrarily named **sections**. The section name appears on a
line by itself, in square brackets `[` and `]`. All **keys** after the **section** declaration are associated
with that **section**. There is no explicit **"end of section" delimiter**; **sections** end at the next 
**section** declaration, or the end of the file. **Sections** may not be nested.

**Note:** In the Windows implementation the **section** cannot contain the character closing bracket `]`.

    [section]
    a=a
    b=b

#### Case insensitivity

**Section** and **key** _names_ are **not case sensitive** in the Windows implementation.

#### Comments

**Semicolons** `;` at the beginning of the line indicate a **comment**. **Comment** lines are ignored.

    ; comment text

See simple example **ini file** in `example/ex.ini`

    ; begin of the animals section

    [Animals]

    animal1=dog
    animal2=cat
    animal3=dino # may be Tyranozaurus or Diplodocus 

    ; begin of the birds section

    [Birds]

    bird1=cock, crown, eagle
    bird2=cuckoo

    ; Lets say pterodactyl
    ; is not a dynozaurus
    bird3=ptero ; just a pterodactyl

    ; bird4 is disabled
    : bird4=crocodile

    ; this is the end


