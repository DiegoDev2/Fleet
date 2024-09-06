package main

// GnomeCommon - Core files for GNOME
// Homepage: https://gitlab.gnome.org/GNOME/gnome-common

import (
	"fmt"
	
	"os/exec"
)

func installGnomeCommon() {
	// Método 1: Descargar y extraer .tar.gz
	gnomecommon_tar_url := "https://download.gnome.org/sources/gnome-common/3.18/gnome-common-3.18.0.tar.xz"
	gnomecommon_cmd_tar := exec.Command("curl", "-L", gnomecommon_tar_url, "-o", "package.tar.gz")
	err := gnomecommon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnomecommon_zip_url := "https://download.gnome.org/sources/gnome-common/3.18/gnome-common-3.18.0.tar.xz"
	gnomecommon_cmd_zip := exec.Command("curl", "-L", gnomecommon_zip_url, "-o", "package.zip")
	err = gnomecommon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnomecommon_bin_url := "https://download.gnome.org/sources/gnome-common/3.18/gnome-common-3.18.0.tar.xz"
	gnomecommon_cmd_bin := exec.Command("curl", "-L", gnomecommon_bin_url, "-o", "binary.bin")
	err = gnomecommon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnomecommon_src_url := "https://download.gnome.org/sources/gnome-common/3.18/gnome-common-3.18.0.tar.xz"
	gnomecommon_cmd_src := exec.Command("curl", "-L", gnomecommon_src_url, "-o", "source.tar.gz")
	err = gnomecommon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnomecommon_cmd_direct := exec.Command("./binary")
	err = gnomecommon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
