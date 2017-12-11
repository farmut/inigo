# Inigo

Golang written **.ini files parser** with disabled (_commented_) options support.

           o8o               o8o
           `"'               `"'
          oooo  ooo. .oo.   oooo   .oooooooo  .ooooo.
          `888  `888P"Y88b  `888  888' `88b  d88' `88b
           888   888   888   888  888   888  888   888
           888   888   888   888  `88bod8P'  888   888
          o888o o888o o888o o888o `8oooooo.  `Y8bod8P'
          ___.ini files parser___ d"     YD
                                  "Y88888P'

1. Description
    1.1 App
    1.2 Format spec.
2. Install
3. Usage
5. Contrib
6. TODO

## Description

### App

### Ini file format

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


