package main

import (
	"fmt"
	"os"
	"os/exec"
)

var gettext = &Formula{
	Name:        "gettext",
	Description: "GNU internationalization (i18n) and localization (l10n) library",
	Homepage:    "https://www.gnu.org/software/gettext/",
	URL:         "https://ftp.gnu.org/gnu/gettext/gettext-0.21.tar.gz",
	Sha256:      "d20fcbb537e02dcf1383197ba05bd0734ef7bf5db06bdb241eb69b7d16b73192",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading gettext...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/gettext/gettext-0.21.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting gettext...")
		cmd = exec.Command("tar", "-xzf", "gettext-0.21.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing gettext...")
		cmd = exec.Command("./configure")
		if err := cmd.Run(); err != nil {
			return err
		}

		cmd = exec.Command("make")
		if err := cmd.Run(); err != nil {
			return err
		}

		cmd = exec.Command("sudo", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing gettext installation...")
		cmd := exec.Command("gettext", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := gettext.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := gettext.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("gettext installed and tested successfully!")
}
