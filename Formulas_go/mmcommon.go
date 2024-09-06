package main

// MmCommon - Build utilities for C++ interfaces of GTK+ and GNOME packages
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installMmCommon() {
	// Método 1: Descargar y extraer .tar.gz
	mmcommon_tar_url := "https://download.gnome.org/sources/mm-common/1.0/mm-common-1.0.6.tar.xz"
	mmcommon_cmd_tar := exec.Command("curl", "-L", mmcommon_tar_url, "-o", "package.tar.gz")
	err := mmcommon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmcommon_zip_url := "https://download.gnome.org/sources/mm-common/1.0/mm-common-1.0.6.tar.xz"
	mmcommon_cmd_zip := exec.Command("curl", "-L", mmcommon_zip_url, "-o", "package.zip")
	err = mmcommon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmcommon_bin_url := "https://download.gnome.org/sources/mm-common/1.0/mm-common-1.0.6.tar.xz"
	mmcommon_cmd_bin := exec.Command("curl", "-L", mmcommon_bin_url, "-o", "binary.bin")
	err = mmcommon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmcommon_src_url := "https://download.gnome.org/sources/mm-common/1.0/mm-common-1.0.6.tar.xz"
	mmcommon_cmd_src := exec.Command("curl", "-L", mmcommon_src_url, "-o", "source.tar.gz")
	err = mmcommon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmcommon_cmd_direct := exec.Command("./binary")
	err = mmcommon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
