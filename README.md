# MyIP Snap

Snap to find out which IP address a box has. You need to connect a USB LED Dream Cheeky and for the first two minutes after the system boots, it will every 20 seconds scroll all the interfaces and their v4 IP addresses.

Build:
git clone ...
snapcraft

Run:

sudo snap install myip_0.1_amd64.snap --force-dangerous --devmode

When dcled can be put behind a serial interface, this snap would be able to be uploaded to the store but for the moment there is no functionality to automatically detect a USB device that gets plugged in and to connect to it.
