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
)

func main() {

	// ===================================================================
	// 								Pre-installation
	// ===================================================================
	// Setup keyboard map.

	// Verify the boot mode
	_, err := os.Stat(efiVerifyPath)
	if os.IsNotExist(err) {
		gocolorful.Err("DON'T HAVE ", efiVerifyPath, ", MAY BE NOT ENABLE UEFI MODE.")
	}

	// Connect to the internet
	gocolorful.Info("Checking Network Connection...")
	ping := exec.Command("ping", networkVerifySite, "-t", "4")
	o, err := ping.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(o))

	// Update the system clock
	gocolorful.Info("Setting ntp service...")
	setNtp := exec.Command("timedatectl", "set-ntp")
	setNtp.Run()

	gocolorful.Info("Check ntp status...")
	ntpStatus := exec.Command("timedatectl", "status")
	ntpStatus.Run()

	// Partition the disks
	// TODO: Complete it.
	gocolorful.Info("Checking disk information...")
	partStatus := exec.Command("gdisk")

	// Format the partitions

	// Mount the file systems

	// ===================================================================
	// 								Installation
	// ===================================================================
	// Select the mirrors
	// * -> echo "Server = https://mirrors.ustc.edu.cn/$repo/os/$arch" > /etc/pacman.d/mirrorlist

	// Install the base packages
	// * -> pacstrap /mnt base

	// ===================================================================
	// 								Configuration
	// ===================================================================
	// Fstab
	// * -> genfstab -U /mnt >> /mnt/etc/fstab

	// Chroot
	// * -> arch-chroot /mnt

	// Time zone
	// * -> ln -sf /usr/share/zoneinfo/Asia/ShangHai /etc/localtime
	// * -> hwclock --systohc

	// Localization
	// * -> locale-gen
	// * -> echo "LANG=en_US.UTF-8" >> /etc/locale.conf

	// Network configuration
	// * -> echo "$HOSTNAME" >> /etc/hostname
	// * -> "127.0.0.1	localhost\n::1		localhost\n127.0.1.1	myhostname.localdomain	myhostname" > /etc/hosts

	// Initramfs
	// * -> mkinitcpio -p linux

	// Root password
	// * -> passwd

	// Boot loader
	// * -> pacman -S grub

	// ===================================================================
	// 								Reboot
	// ===================================================================
	// Reboot system
	// * -> reboot
}
