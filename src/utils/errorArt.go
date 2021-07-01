package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var errorId = "PID"

// ErrorMsg displays cute error ASCI Art for PID Errors
func ErrorMsg(id string) {
	switch id {
	case "cid":
		errorId = "CID"
	case "cname":
		errorId = "CNAME"

	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(7) //gives a pseudorandom number in the range [0, n) (n not included).
	switch num {
	case 0:
		errorDoggo()
	case 1:
		errorBunny()
	case 2:
		errorDolphy()
	case 3:
		errorOwl()
	case 4:
		errorBeaver()
	case 5:
		errorElephant()
	}
}

func errorDoggo() {
	dog := `
	ID no exist, done doggo a sad
		\
		 \
		/^-----^\
		V  o o  V
		 |  Y  |
		  \ ⌓ /
		  / - \
		  |    \
		  |     \     )
		  || (___\====
	`
	fmt.Println(strings.ReplaceAll(dog, "ID", errorId))
}

func errorOwl() {
	goobes := `

   /\_/\  The council of wise owls are confused!  /\_/\
  ((@v@))      Please provide a valid ID!       ((@v@))
 ():::::()                                      ():::::()
   VV-VV          /\_/\         /\_/\             VV-VV
                 ((@v@))       ((@v@))
                ():::::()     ():::::()
                  VV-VV         VV-VV
  `
	fmt.Println(strings.ReplaceAll(goobes, "ID", errorId))
}

func errorDolphy() {
	dolphy := `
                               _.-~  )
                    _..--~~~~,'   ,-/     _
                 .-'. . . .'   ,-','    ,' )
               ,'. . . _   ,--~,-'__..-'  ,'
             ,'. . .  (@)' ---~~~~      ,'
            /. . . . '~~             ,-'
           /. . . . .             ,-'
          ; . . . .  - .        ,'
         : . . . .      \_     /       ID did Dolphy a daze,
        . . . . .          \-.:        Please enter valid ID
       . . . ./  - .          )
      .  . . |  _____..---.._/ _____________
~---~~~~----~~~~             ~~
`
	fmt.Println(strings.ReplaceAll(dolphy, "ID", errorId))
}

func errorBunny() {
	bunny := `             ,
            /|      __
           / |   ,-~ /
          Y :|  //  /
          | jj /( .^
          >-"~"-v"
         /       Y
        jo  o    |
       ( ~T~     j
        >._-' _./
       /   "~"  |
      Y     _,  |
     /| ;-"~ _  l
    / l/ ,-"~    \
    \//\/      .- \
     Y        /    Y     Bunny couldn't recognise that ID.
     l       I     !      Done bunny a confuse.
     ]\      _\    /"\     Please give bunny a valid ID.
    (" ~----( ~   Y.  )
~~~~~~~~~~~~~~~~~~~~~~~~~
  `

	fmt.Println(strings.ReplaceAll(bunny, "ID", errorId))
}

func errorBeaver() {
	beaver := `
    /   \          /   \
    \_   \        /  __/
    _\   \      /  /__
    \___  \____/   __/
        \_       _/
          | @ @  \_
          |
        _/     /\
        /o)  (o/\ \_
        \_____/ /
          \____/
  Whoooopsssss, invalid ID. 
  Please enter a valid ID.
  `
	fmt.Println(strings.ReplaceAll(beaver, "ID", errorId))
}

func errorElephant() {
	elephant := `
                         ____
                    ---'-    \
      .-----------/           \
     /           (         ^  |   __
&   (             \        O  /  / .'
'._/(              '-'  (.   (_.' /
     \                    \     ./
     |    |       |    |/ '._.'
     )   @).____\|  @ |
 .  /    /       (    |           Pawoo. Pawoo. Pawoo!
\|, '_:::\  . ..  '_:::\ ..\).    Plz give Elephant a valid ID.
  `
	fmt.Println(strings.ReplaceAll(elephant, "ID", errorId))
}
