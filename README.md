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

The basic element contained in an `INI file` is the **key** or property. Every **key** has a _name_ and a 
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

_Semicolons_ `;` at the beginning of the line indicate a **comment**. **Comment** lines are ignored.

    ; comment text

#### Varying features

The `INI file` format is not well defined. Many programs support features beyond the basics described above. The following is a list of some common features, which may or may not be implemented in any given program.

##### Blank lines

Some rudimentary programs do not allow **blank lines**. Every line must therefore be a **section** head, a 
**parameter**, or a **comment**. **Inigo** ignores **blank lines** founded.

##### Unix-style comments

Some software supports the use of the _number sign_ `#` as an alternative to the _semicolon_ for indicating
**comments**. Practically speaking, using it to begin a line may effectively change a variable name on that
line. For instance, the following line creates a variable named `#var`, but not one named `var`; this 
is sometimes used to create a pseudo-implementation of a **comment**. 

    #var=a

More generally, use of the _number sign_ is unpredictable, as in these following lines (note the _space_ 
after the number sign in the second line). For this reason, the _number sign_ character should not be used 
to begin comments.

Unlike this, **Inigo** parses **Unix-style comment** such as ordinary comment with `;`.

    #[section]
    # var=a
    # Unix-style comment

##### Duplicate names

Most implementations only support having one **parameter** with a given name in a section. The second 
occurrence of a **parameter** _name_ may cause an abort, it may be ignored (and the _value_ discarded), or 
it may override the first occurrence (with the first _value_ discarded). Some programs use duplicate 
**parameters** names to implement multi-valued properties.

Interpretation of multiple **section** declarations with the same name also varies. In some implementations,
duplicate **sections** simply merge their _properties_ together, as if they occurred contiguously. Others may
abort, or ignore some aspect of the `INI file`.

#### Examples of INI files

See examples of `INI files` in `example` folder of app's repo.
One file, `example/ex.ini`, is the simple `INI file` for illustration of base `INI` features:

    ; example/ex.ini file begin

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

Other file, `examples/ex_php.ini` it realy `php.ini` file representing more **complex case**.

