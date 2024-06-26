# WowWee MiP

![MiP](../../images/mip.png)

About the MiP...

## What you need

    - MiP robot
    - Personal computer with Go installed, and a Bluetooth 4.0 radio.
    - Linux, macOS, or Windows

## Installation

The code in this activity code uses the TinyGo Bluetooth package http://tinygo.org/bluetooth

It also uses the https://github.com/hybridgroup/tinygo-mip package which is where the code wrappers for the WowWee MiP Bluetooth API are located.

Change directories into this directory where the needed Go modules files are located, and all dependencies will be installed.

## Running the code

When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the MiP using the Bluetooth interface.

On Linux and Windows you will use the MAC address of the device to connect.

On macOS you must use the Bluetooth ID of the device to connect.

Therefore, you must know the correct name and then MAC address or ID for that robot in order to connect to it.

To find out the unique Bluetooth ID assigned to that device, you can use the Bluetooth scanner located in the `tools/blescanner` directory of this repo:

```
cd tools
go run ./blescanner
```

Press "Control-C" to stop the `blescanner` program. 

## Code

### step1

This tests that the WowWee MiP is connected correctly to your computer, by turning on the chest LED.

```
go run ./step1/ [MAC address or Bluetooth ID]
```

### step2

Now lets play with the LEDs, and put all of our actions into a single function. Look at the `performActions()` function in `actions.go`.

```
go run ./step2/ [MAC address or Bluetooth ID]
```

### step3

Now lets try turning around. Stand up the MiP so it balances itself, then run this program. It will turn around left and right in place.

```
go run ./step3/ [MAC address or Bluetooth ID]
```

### step4

Now lets move forwards and backward. Stand up the MiP so it balances itself, then run this programs. It will turn around and go forward, then turn around and come back.

```
go run ./step4/ [MAC address or Bluetooth ID]
```

Now you should have the basics for moving MiP around. Try modifying the `performActions()` function to combine movement and LEDs however you like.

### step5

Now it is time for free driving controlled by you, the human pilot. Plug in the DS3 controller to your computer. The controls are as follows:

* Triangle    - Standup
* Right stick - direction and speed


IMPORTANT NOTE: you must press the "P3" button when your program first runs for the "clone" DS3 joysticks we are using to fully turn on.

## What now?

Try adding some new features to the https://github.com/hybridgroup/tinygo-mip repo. There are a number of interesting unimplemented API functions to explore!

## License

Copyright (c) 2015-2024 The Hybrid Group and friends. Licensed under the MIT license.
