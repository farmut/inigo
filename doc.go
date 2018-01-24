// Copyright 2018 hIMEI
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//////////////////////////////////////////////////////////////////////////

/*
Package inigo is an INI files parsing library for Golang applications. It use wide range of recognised
options trying to enhance common INI syntax.

Short info about INI files from Wikipedia (https://en.wikipedia.org/wiki/INI_file):

"The INI files format is an informal standard for configuration files for some platforms or software.
INI files are simple text files with a basic structure composed of sections, properties, and values.
The name "INI file" comes from the commonly used filename extension .INI, which stands for
_"initialization"_.

In MS-DOS and 16-bit Windows platforms up through Windows ME, the INI files served as the primary
mechanism to configure operating system and installed applications features, such as device drivers,
fonts, startup launchers, and things that needed to be initialized in booting Windows. INI files
were also generally used by applications to store their individual settings.

Starting with Windows NT, Microsoft favored the use of the registry, and began to steer
developers away from using INI files for configuration. But the APIs still exist in Windows, however,
and developers may still use them.

Linux and Unix systems also use a similar file format for system configuration. In addition,
platform-agnostic software may use this file format for configuration. It is human-readable and simple
to parse, so it is a usable format for configuration files that do not require much greater complexity.
For example, the platform-agnostic PHP uses the "INI" format for its php.ini configuration file in
both Windows and Linux systems".

Ini file format

Inigo use next conception and rules of INI syntax:

1. Ini file structure

        <GLOBAL PARAMETERS>

        ; commented line

        [SECTION1]

            PARAM1=value1

            PARAM1=value2

        [SECTION2]

            PARAM3=value3

            ...

2. Sections

Keys MAY BE grouped into arbitrarily named sections. The section name appears on a
line by itself, in square brackets "[" and "]". All keys after the section declaration are associated
with that section. There is no explicit "end of section" delimiter so sections end at the next
section declaration, or the end of the file. Sections may not be nested.

Note: In the Windows implementation the section cannot contain the character closing bracket "]".

    [section]
    a=a
    b=b

3. Parameters

In common cases parameter names parses as string always. But Inigo breaks this rule and parse
underlying value of parameter. So parsed value may be:

- Int
- Uint
- Int in hexademical, octal or binary representation
- Boolean
- String (default)
- Array (e.g Golang slice) of strings
- Map [string]string

Parameter name may contains any symbols except space.

    parameter1="Iggy Pop"
    new parameter=20.6465 // Error

Note: In the Windows implementation the key cannot contain the characters equal sign "=" or semicolon
 ";" as these are reserved characters. The _value_ can contain any character.

Case insensitivity

Section and key _names_ are not case sensitive in the Windows implementation.

Comments

_Semicolons_ ";" at the beginning of the line indicate a comment. Comment lines are ignored.

    ; commented string
    # Unix-style commented string

4. Varying features

The INI files format is not well defined. Many programs support features beyond the basics described
above. The following is a list of some common features, which may or may not be implemented in any
given program.

Blank lines

Some rudimentary programs do not allow blank lines. Every line must therefore be a section head, a
parameter, or a comment. Inigo ignores blank lines founded.

Unix-style comments

Some software supports the use of the _number sign_ "#" as an alternative to the _semicolon_ for indicating
comments. Practically speaking, using it to begin a line may effectively change a variable name on that
line. For instance, the following line creates a variable named "#var", but not one named "var"; this
is sometimes used to create a pseudo-implementation of a comment.

    #var=a

More generally, use of the _number sign_ is unpredictable, as in these following lines (note the _space_
after the number sign in the second line). For this reason, the _number sign_ character should not be used
to begin comments.

Unlike this, Inigo parses Unix-style comment as ordinary comment with ";".

    #[section]
    # var=a
    # Unix-style comment

Duplicate names

Most implementations only support having one parameter with a given name in a section. The second
occurrence of a parameter _name_ may cause an abort, it may be ignored (and the _value_ discarded), or
it may override the first occurrence (with the first _value_ discarded). Some programs use duplicate
parameters names to implement multi-valued properties.

Interpretation of multiple section declarations with the same name also varies. In some implementations,
duplicate sections simply merge their _properties_ together, as if they occurred contiguously. Others may
abort, or ignore some aspect of the INI files.

Arrays

Russian version https://ru.wikipedia.org/wiki/.ini of Wikipedia article shows an examples of
working with arrays in INI files as Zend Framework do:

    ; in Zend Framework array is determined as in next statement
    [Section3]
    var1[]=значение_1_1
    var1[]=значение_1_2
    var1[]=значение_1_3
    var2=значение_2

But Inigo don't recognize such syntax.

Examples of INI files

See examples of INI files in "example" folder.

    ;;;;;;;;;;;;;;;;;;;;;
    ;; ex.ini file begin
    ;;;;;;;;;;;;;;;;;;;;;

    ; global parameters

    animals=animals
    birds=birds
    insects=insects

    ; begin of the animals section

    [Animals]

    animal1=dog
    animal2=cat
    animal3=dino # may be Tyranozaurus or Diplodocus

    ; Add some info
    [DetailOfAnimal]

    NumOfLegs = 4
    NumOfYeys=2

    ColorOfPolarBear="dirty white"
    ColorOfTiger=yellow black striped # error

    ; Insects it is a very socialized creatures

    [insects]

    ; Number of spider's legs

    spiderLegs=0x8

    ; Other insects have 6 legs

    other_insect_legs=00000110

    ; Case of ants
    ColonyPopulation=137540000

    ; Case of bees
    Colonypopulation=1200005.3544345435

    ; Total amounts of insects
    Total="very big number"

    ; begin of the birds section

    [Birds]

    bird1=cock, crown, eagle
    bird2=cuckoo
    angryByrds=first:bird00x2, second:antibird, third:hyperBird

    ; Lets say pterodactyl
    ; is not a dynozaurus
    bird3=ptero ; just a pterodactyl

    ; bird4 is disabled
    : bird4=crocodile

    ; this is the end

*/
package inigo
