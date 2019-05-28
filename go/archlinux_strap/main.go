package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fintanchen/gocolorful"
)

const (
	efiVerifyPath     = "/sys/firmware/efi/efivars"
	networkVerifySite = "www.baidu.com"
)

func main() {

	// ===================================================================
	// 								Pre-installation
	// ===================================================================
	// Setup keyboard map.

	// Verify the boot mode
	_, err := os.Stat(efiVerifyPath)
	if os.IsNotExist(err) {
		// gocolorful.Err("DON'T HAVE ", efiVerifyPath, ", MAY BE NOT ENABLE UEFI MODE.")
	}

	// Connect to the internet
	gocolorful.Info("Check Network Connection...")
	ping := exec.Command("ping", networkVerifySite, "-t", "4")
	o, err := ping.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(o))

	// Update the system clock
	gocolorful.Info("Check Network Connection...")
	setNtp := exec.Command("timedatectl", "set-ntp")
	setNtp.Run()

	gocolorful.Info("Check ntp status...")
	ntpStatus := exec.Command("timedatectl", "status")
	ntpStatus.Run()

	// Partition the disks

	// Format the partitions

	// Mount the file systems

	// ===================================================================
	// 								Installation
	// ===================================================================
	// Select the mirrors

	// Install the base packages

	// ===================================================================
	// 								Configuration
	// ===================================================================
	// Fstab

	// Chroot

	// Time zone

	// Localization

	// Network configuration

	// Initramfs

	// Root password

	// Boot loader

	// ===================================================================
	// 								Reboot
	// ===================================================================
	// Reboot system
}
