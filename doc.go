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
Package inigo is an INI files parsing library for Golang applications with wide
range of parsed data types.

Inigo's features:

* Inigo parses disabled (commented) parameters of INI file as well as enabled ones.
* Unlike some others INI parsers, Inigo recognizes parameter's value not just as string,
but tries to parse its native data type. Here is the types that Inigo understands:

    - Int64
    - Uint64
    - Float64
    - Boolean
    - Octal representation of Int
    - Hexadecimal representation of Int
    - String (default)
    - Array (as **Go**'s slice of strings)
    - Map (currently _[string]string_ only)

* Inigo also parses ini file's parameters with syntax errors and stores its as special data type.
* Inigo it is INI parser, not editor. For INI files editing use your preferred
 text editor (Vim is great!)

About INI files

Short info about INI files from Wikipedia (https://en.wikipedia.org/wiki/INI_file):

"The INI files format is an informal standard for configuration files for some platforms or software.
INI files are simple text files with a basic structure composed of sections, properties, and values.
The name "INI file" comes from the commonly used filename extension .INI, which stands for
"initialization".

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

The INI files format is not well defined. Many programs support features beyond the basics described
above. The following is a list of some common features, which may or may not be implemented in any
given program.

Inigo use next conception and rules of INI syntax:

1. Ini file structure

        <GLOBAL PARAMETERS>

        ; commented line

            param0=value0

        [SECTION1]

        ; commented line

            PARAM1=value1

        ; commented line

            PARAM1=value2

        [SECTION2]

        ; commented line

            PARAM3=value3

            ...

2. Sections

Keys MAY BE grouped into arbitrarily named sections. The section name appears on a
line by itself, in square brackets "[" and "]". All keys after the section declaration are associated
with that section. There is no explicit "end of section" delimiter so sections end at the next
section declaration, or the end of the file. Sections may not be nested.

    [section]
    a=a
    b=b

3. Parameters

In common cases parameter names parses as string.

Parameter name may contains any symbols except space. Trailing spaces at the end of parameter name
are ignored.

    parameter1="Iggy Pop"
    new parameter=20.6465 // Error

The value can contain any character but if STRING value contains the spaces, such value must
be doublequotted. Otherwise value will be parsed as error.

    parameter1=Iggy Pop // Error also

Leading space at begin of parameter value are ignored.

    parameter1 = "Iggy Pop"
    parameter= 20.6465

4. Comments

Semicolons ";" and sharp "#" (the UNIX style comments) at the beginning of the line indicate a comment.
Commented lines are ignored.

    ; commented string
    # Unix-style commented string

Inline comments with ";" and "#" is also recognized by Inigo.

    param=value ; inline comment are ignored
    param=value # inline comment are ignored

5. Blank lines

Some rudimentary programs do not allow blank lines. Every line must therefore be a section head, a
parameter, or a comment. Inigo ignores blank lines.

6. Compound data types

Inigo parses next compound data types:
* arrays as Go's slices of strings. In this case parameter's value must looks as

    slice=first, second, third

* map[string]string. In this case parameter's value must looks as

    map=first:value1, second:value2, third:value3

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

    other_insect_legs=06

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
