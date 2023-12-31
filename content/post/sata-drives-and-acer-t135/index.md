---
title: 'SATA drives and Acer T135'
date: Sun, 14 Dec 2008 16:46:00 +0000
draft: false
tags: ['computer-hardware']
---

More on the SATA / Acer T135 saga.

Having got the PC to "see" the drive in the BIOS, it should be a simple matter of restoring the system volume onto this drive (or cloning it from the old drive), making it a primary, active partition (so it's bootable into Windows), and off we go, yes?

So I created 1 large partition for the system and restored onto it. Did this work? No. Not in these Acer machines.

According to boot.ini (which the wonderful StorageCraft ShadowProtect software can read easily), partition 2 was set as the boot volume. And I only had 1 partition - so better change it.
So I edited it to be partition 1, and tried again. No luck.

I recall that the machine was supplied with the drive having two partitions -

*   Partition 1, around 2.8 Gbytes, for automatic 'recovery' of the machine
*   Parition 2, the usual Windows system volume.

The recovery process is an Acer thing. Basically on boot, you hold down a key (Alt+F10) and the machine goes into auto-recovery mode. This basically re-initialises the system partition (2) from a factory-created archive held on partition 1. That's what the 2.8 Gbytes is used for. It provides a degree of safety for uses who trash up their machines so they won't boot, and I suppose it's cheaper than supplying Windows installation media.

I'm not a great fan of this kind of auto-recovery (after all, it will trash all your programs and data, for sake of getting Windows re-installed). I prefer to use StorageCraft ShadowProtect with an external USB drive, as then I can get the machine back exactly to "how it was" the last time backed up.

Out of a mix of desperation and curiosity, I tried holding down Alt+F10 as the machine booted (off the new disk, which did not have any recovery software on it, only a clean system image). Voila! it booted into Windows, off my new partition 1. So this told me two things:

1.  I could restore the machine, somehow.
2.  The Acer BIOS seems to be ignoring boot.ini, and doing its own thing, based on this Alt+F10 key (or lack of it).
3.  The Alt+F10 key forces a boot from partition 1. It works, but it's a bit inconvenient for everyday use.

## Another attempt

Now knowing this about the Acer machine, I tried again, but this time

*   Small initial partition (around 2.8 GB), which I put a copy of the Acer restore stuff onto from the original disk
*   Large system partition (the rest of the disk) for Windows etc.


The first was copied from the original Acer-supplied hard disk. The second came from my ShadowProtect backup.

Still no luck. It would not boot into Windows. But with Alt+F10 held down, it would boot into the Acer Recovery software on Partition 1.

## Yet Another Attempt

Finally, in desperation - I thought - why not run this Acer recovery software, and at least see if it can make the machine (normally) bootable, off a freshly restored partition 2. So I let it restore, effectively over-writing partition 2 with its virgin install of Windows + the Acer utilities etc. that came with the machine. (No CD-ROMs are supplied as standard).

This resulted in a machine that actually boots up (although missing a lot of stuff, obviously). So - here's the trick - that Acer Restore feature is actually necessary for a machine rebuild - in a totally non-obvious way. boot.ini looked just the same, before and after the restore.

I wonder what the Acer Restore does to the disk. It does some magic that is not obvious.

So ... having proved that was possible, I then restored the ShadowProtect image back over partition 2. And that finally worked! I had a machine that booted normally, and had everything back where we wanted it.

I still don't know why the machine is so fickle with its boot-up, and why it ignores Windows' own boot.ini.

## Conclusions

*   If you have an Acer T135, that 2.8G of "stuff" on the Restore partition is essential to the continued use of the machine, if you replace hard disks. Don't just think that having valid Windows install media and a valid key will be enough to get you out of trouble.

*   Acer T135 has its own boot control which ignores Windows.
*   Don't forget to jumper new SATA drives down to the older standard, if they can be detected.

and finally

*   Don't buy machines with proprietary motherboards and BIOSs if you can help it! Sooner or later they will drive you mad.

Recommended tools and parts for this kind of work:

*   Patience
*   Anti-static wrist strap
*   Screws for holding the new drive in
*   SATA cable (usually not supplied with 'OEM' drive packages)
*   Jumper (again, usually not supplied)
*   Tweezers for fitting the jumper
*   Partitioning / formatting / backup / restore software, such as Norton Ghost, Acronis TrueImage, Partition Magic. My personal favourite at the moment is [Shadowcraft StorageProtect 3.2](http://www.storagecraft.com/) which works fast and well in my experience. It comes with a bootable CD-ROM image that means you can work with machines even when the hard disks are unbootable. Review of the software [here](http://www.pcmag.com/article2/0,2817,2254465,00.asp).