# OGWC - OGame Win Calculator

OGWC is a web based utility to compute the fair distribution of an ACS combat report in the browsergame [OGame](https://lobby.ogame.gameforge.com). It allows users to sum up multple combat reports into one, in order to compute the total distribution. Missilie Reports and Recycler reports can also be added.

The tool in its current version is only localized in german, however Pull Requests in this area are highly welcome.

## Used technologies

This programm is consists of two main components: the frontend written in TypeScript using Angular and the backend written in Go using the Gin web framework. To store calculations on the backend either the embedded Bolt DB or Redis is used depending on the configuration. To obtain the API data from the OGame servers the tool uses the API proxy provided by https://ogapi.rest/ The architecuture of the application is based upon the idea of a convetional web application to potentially allow the use of direct access to the API which has to be done using a predefined set of IP addresses. However this was never implemented, due to lack of documanetation.

## Running OGWC

To run OGWC in the most basic configuration using BoltDB just download the latest version from the Releases section for your platorm and open a command prompt, e.g. cmd.exe or a Fish shell. After navigating to the directory where the binary is located just run:

```
./ogwc server
```

To see how to run the application using a different Database path for BoltDB or using Redis please consult the Help page using `./ogwc --help` and `./ogwc server --help`.

## Building OGWC

To build the application run `npm install` in the UI directory and make sure the Rice utility is in the path. Next run the `build.fish` script. 

## License

The source code of the application is licensed under AGPL v3. See [LICENSE](/LICENSE) for more information.

## Legal Notices

- OGame is a registered trademark of Gameforge AG in Karlsruhe, Germany
- This tool is not affiliated with Gameforge AG
