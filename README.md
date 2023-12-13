# QPlayerData
This is just some format for saving Minecraft Classic player data, nothing crazy.

## Format
* int16 = little endian int16 (2 bytes)
* string = string prefixed by an int16

`
string | name - the player's name
int16  | x - the player's X position
int16  | y - the player's Y position
int16  | z - the player's Z position
byte   | yaw - the player's horizontal orientation (yaw)
byte   | pitch - the player's vertical orientation (pitch)
`