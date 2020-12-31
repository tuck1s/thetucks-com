---
title: 'Acer T135 with and new SATA drives - a snag'
date: Sun, 07 Dec 2008 00:29:00 +0000
draft: false
tags: ['computer-hardware']
---

Just came across a nasty little snag with SATA disk drives. I thought I'd write it up, in case you run into something similar.

One of our family computers is an Acer Aspire T135, around 3 years old now. The 80GB hard drive is getting rather full and slow. Time for an upgrade.

Step 1 - check what the existing PC can do.
It has a spare SATA connection on the motherboard, so that seems like the best way to go.
There is plenty of space in the case for a new drive.
There is a spare SATA-compatible power connector (different to the old-style 4-pin Molex's) dangling in the case.
I'll need to get a new SATA data cable for the drive.

Step 2 - plan how to use the drive
My initial thought is to use the new drive for the OS and programs, and the old drive as scratch area and - if it makes things run faster - the windows swap file.

Step 3 - choose a drive
A quick look at drive specs etc. suggests a 500GB drive as the best price point at the moment (given that the machine does not have to handle huge amounts of data usually).

A quick check on the Acer website suggests no particular compatibility requirements for second drives (then again, it doesn't really say a lot about their machines anyway).

Western Digital do a nice looking "Green Power" model which claims to save power. As family PCs get left on all the time, that seems like a good idea.

I got mine from Novatech as they are local (www.novatech.co.uk).

Step 4 - install the drive
Short work with a screwdriver and the new drive was ready to run.
Boot up, check in the BIOS that the drive was recognised ... nothing, only the existing drive is found.

I tried various things, including:

*   Swap the existing and new drive data cables
*   Swap the existing and new drive power cables
*   Try with another SATA drive I had available (an older 160GB Maxtor). Ah-ha! That works!

So why does the new drive not work?

Step 5 - fault-finding .. it's NOT the BIOS ...

*   Maybe the BIOS is too old to read 500GB disks. So I flashed the BIOS with a newer version, directly from Gigabyte's website, rather than from Acer's website.
*   The motherboard is a K8VM800. But it's labelled as the K8VM800MAE on the motherboard itself. Gigabyte list two versions - Rev. 1 and Rev. 2 - but nothing about the MAE. According to the chips fitted it's closest to Rev. 1. (looking at the Ethernet chip type).

*   Incidentally, flashing BIOSs on computers like the Acer T135 that lack a floppy drive is somewhat time-consuming and troublesome. It hasa nice integrated flash card reader though ... so ... After a bit of Googling, and making a DOS-bootable disk on a Flash SD card, which worked fine on one of my other machines, I gave up, yanked a floppy drive out of another machine, and used that instead ...
*   It flashed OK (but with a warning message), and it ran the computer OK, but it kept giving checksum errors on startup. Not such a good idea. Revert back to the original BIOS (which I had kept).

Step 6 - try another 500GB drive

As it happens, I have a 500GB SATA drive (different brand - Samsung) in another machine. So - let's try that.
It works! So there isn't a capacity issue. Must be something weird with the drive.

Step 7 - has anyone else hit this before?
After some tricky Googling to avoid the hundreds of old review articles on T135's, many that have a different spec, I found a forum that suggested that:


*   SATA has different speeds - 150MB/sec and 300 MB/sec
*   While everything is supposed to auto-adjust, it might not


A quick check on the drive manufacturer's website pulls up a datasheet, and -- guess what -- there are some jumpers to force the drive down to the slower speed.

[http://wdc.custhelp.com/cgi-bin/wdc.cfg/php/enduser/std\_adp.php?p\_faqid=1679](http://wdc.custhelp.com/cgi-bin/wdc.cfg/php/enduser/std_adp.php?p_faqid=1679)

It's that OPT1 jumper (pins 5-6 on the drive family I'm using).

Now the BIOS sees the drive. But the original drive seems to be taking ages to boot up in to Windows (despite not being changed at all). More on this tomorrow.