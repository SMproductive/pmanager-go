# Pmanager-go
## The most straight forward password manager

### Installation
#### Gnu/Linux
* Use precompiled pmanager-go.tar.xz with Makefile
* Or build it yourself using go

#### Windows
* Just execute pmanager-go.exe
* Or compile it yourself using go

### Usage

#### Title entries
* Right click: enables it for change (submit change with enter -> save change with save button)
* Left click: opens content
* Empty: will be deleted on save
[![\nCan not display image!](https://github.com/SMproductive/pmanager-go/pics/titleEntries.png)]

#### Content entries
* Right click: enables it for change (submit change with enter -> save change with save button)
* Double left click: toggles visibility of text
* Left click: copies the text to clipboard
* If empty: will be deleted on save
[![\nCan not display image!](https://github.com/SMproductive/pmanager-go/pics/contentEntries.png)]

#### Buttons
* Add: adds a new entry
* Save: saves the database with the provided password
* Change MPW: changes the password for database encryption(save is needed to confirm the new one)
* Generate: generates a random string with provided characters and length in "Random Generator"(default: 21 ASCII characters)
