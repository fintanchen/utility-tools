package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fintanchen/gocolorful"
)

const (
	// Boot
	efiVerifyPath = "/sys/firmware/efi/efivars"

	// Network Connection
	networkVerifySite = "www.baidu.com"

	// mirrors
	mirrorListPath = "/etc/pacman.d/mirrorlist"
	defaultMirror  = "https://mirrors.ustc.edu.cn/$repo/os/$arch"

	// hostname
	hostname = "CMCC-EDU"
)

func main() {

	// ===================================================================
	// 								Pre-installation
	// ===================================================================
	// Setup keyboard map.

	// Verify the boot mode
	// * -> ls /sys/firmware/efi/efivars
	gocolorful.Info("Checking wether UEFI mode...")
	_, err := os.Stat(efiVerifyPath)
	if os.IsNotExist(err) {
		gocolorful.Err("DON'T HAVE ", efiVerifyPath, ", MAY BE NOT ENABLE UEFI MODE.")
	}

	// Connect to the internet
	// * -> ping $website
	gocolorful.Info("Checking Network Connection...")
	ping := exec.Command("ping", networkVerifySite, "-t", "4")
	o, err := ping.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(o))

	// Update the system clock
	// * -> timedatectl set-ntp
	gocolorful.Info("Setting ntp service...")
	setNtp := exec.Command("timedatectl", "set-ntp")
	setNtp.Run()

	// * -> timedatectl status
	gocolorful.Info("Check ntp status...")
	ntpStatus := exec.Command("timedatectl", "status")
	ntpStatus.Run()

	// Partition the disks
	// TODO: Complete it.
	gocolorful.Info("Checking disk information...")
	// partStatus := exec.Command("gdisk")

	// Format the partitions
	// * -> mkfs.ext4 /dev/sda2
	gocolorful.Info("Formating the disk...")
	mkfs := exec.Command("mkfs.ext4", "/dev/sda2")
	mkfsStatus := mkfs.Run()
	if err != nil {
		gocolorful.Err(mkfsStatus)
	}

	// * -> mkswap /dev/sda3
	gocolorful.Info("Making swap partition...")
	mkswap := exec.Command("mkswap", "/dev/sda3")
	mkswapStatus := mkswap.Run()
	if err != nil {
		gocolorful.Err(mkswapStatus)
	}

	// * -> swapon /dev/sda3
	gocolorful.Info("Enable swap partition...")
	swapon := exec.Command("swapon", "/dev/sda3")
	swaponStatus := swapon.Run()
	if swaponStatus != nil {
		gocolorful.Err(swaponStatus)
	}

	// * -> free -m
	gocolorful.Info("Checking swap partition information...")
	swapSpace := exec.Command("free", "-m")
	b, err := swapSpace.Output()
	if err != nil {
		gocolorful.Err(err)
	}
	fmt.Println(string(b))

	// Mount the file systems
	// * -> mount /dev/sda2 /mnt
	gocolorful.Info("Mount the partition...")
	mount := exec.Command("mount", "/dev/sda2", "/mnt")
	mountStatus := mount.Run()
	if mountStatus != nil {
		gocolorful.Err(mountStatus)
	}

	// ===================================================================
	// 								Installation
	// ===================================================================
	// Select the mirrors
	// * -> echo "Server = https://mirrors.ustc.edu.cn/$repo/os/$arch" > /etc/pacman.d/mirrorlist
	writeMorrorsList := exec.Command("echo", "\"Server = ", defaultMirror, "\"", ">", mirrorListPath)
	writeMorrorsListStatus := writeMorrorsList.Run()
	if writeMorrorsListStatus != nil {
		gocolorful.Err(writeMorrorsListStatus)
	}

	// Install the base packages
	// * -> pacstrap /mnt base
	pacstrap := exec.Command("pacstrap", "/mnt", "base")
	pacstrapStatus := pacstrap.Run()
	if pacstrapStatus != nil {
		gocolorful.Err(pacstrapStatus)
	}

	// ===================================================================
	// 								Configuration
	// ===================================================================
	// Fstab
	// * -> genfstab -U /mnt >> /mnt/etc/fstab
	genfstab := exec.Command("genfstab", "-U", "/mnt", ">>", "/mnt/etc/fstab")
	genfstabStatus := genfstab.Run()
	if genfstabStatus != nil {
		gocolorful.Err(genfstabStatus)
	}

	// Chroot
	// * -> arch-chroot /mnt
	archChroot := exec.Command("arch-chroot", "/mnt")
	archChrootStatus := archChroot.Run()
	if archChrootStatus != nil {
		gocolorful.Err(archChrootStatus)
	}

	// Time zone
	// * -> ln -sf /usr/share/zoneinfo/Asia/ShangHai /etc/localtime
	ln := exec.Command("ln", "-sf", "/usr/share/zoneinfo/Asia/ShangHai", "/etc/localtime")
	lnStatus := ln.Run()
	if lnStatus != nil {
		gocolorful.Err(lnStatus)
	}

	// * -> hwclock --systohc
	hwclockSync := exec.Command("hwclock", "--systohc")
	hwclockSyncStatus := hwclockSync.Run()
	if hwclockSyncStatus != nil {
		gocolorful.Err(hwclockSyncStatus)
	}

	// Localization
	// * -> locale-gen
	localeGen := exec.Command("locale-gen")
	localeGenStatus := localeGen.Run()
	if localeGenStatus != nil {
		gocolorful.Err(localeGenStatus)
	}

	// * -> echo "LANG=en_US.UTF-8" >> /etc/locale.conf
	writeLocale := exec.Command("echo", "\""+"LANG=en_US.UTF-8\"", ">>", "/etc/locale.conf")
	writeLocaleStatus := writeLocale.Run()
	if writeLocaleStatus != nil {
		gocolorful.Err(writeLocaleStatus)
	}

	// Network configuration
	// * -> echo "$HOSTNAME" >> /etc/hostname
	setHostname := exec.Command("echo", hostname, ">>", "/etc/hostname")
	setHostnameStatus := setHostname.Run()
	if setHostnameStatus != nil {
		gocolorful.Err(setHostnameStatus)
	}

	// Hosts config
	// * -> "127.0.0.1\tlocalhost\n::1\tlocalhost\n127.0.1.1\tmyhostname.localdomain\tmyhostname" > /etc/hosts

	// Initramfs
	// * -> mkinitcpio -p linux
	mkinitcpio := exec.Command("mkinitcpio", "-p", "linux")
	mkinitcpioStatus := mkinitcpio.Run()
	if mkinitcpioStatus != nil {
		gocolorful.Err(mkinitcpioStatus)
	}

	// Root password
	// TODO: passwd
	// * -> passwd

	// Boot loader
	// * -> pacman -S grub

	// ===================================================================
	// 								Reboot
	// ===================================================================
	// Reboot system
	// * -> reboot
	reboot := exec.Command("reboot")
	rebootStatus := reboot.Run()
	if reboot != nil {
		gocolorful.Err(rebootStatus)
	}
}
