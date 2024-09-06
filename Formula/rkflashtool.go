package main

// Rkflashtool - Tools for flashing Rockchip devices
// Homepage: https://sourceforge.net/projects/rkflashtool/

import (
	"fmt"
	
	"os/exec"
)

func installRkflashtool() {
	// Método 1: Descargar y extraer .tar.gz
	rkflashtool_tar_url := "https://downloads.sourceforge.net/project/rkflashtool/rkflashtool-6.1/rkflashtool-6.1-src.tar.bz2"
	rkflashtool_cmd_tar := exec.Command("curl", "-L", rkflashtool_tar_url, "-o", "package.tar.gz")
	err := rkflashtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rkflashtool_zip_url := "https://downloads.sourceforge.net/project/rkflashtool/rkflashtool-6.1/rkflashtool-6.1-src.tar.bz2"
	rkflashtool_cmd_zip := exec.Command("curl", "-L", rkflashtool_zip_url, "-o", "package.zip")
	err = rkflashtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rkflashtool_bin_url := "https://downloads.sourceforge.net/project/rkflashtool/rkflashtool-6.1/rkflashtool-6.1-src.tar.bz2"
	rkflashtool_cmd_bin := exec.Command("curl", "-L", rkflashtool_bin_url, "-o", "binary.bin")
	err = rkflashtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rkflashtool_src_url := "https://downloads.sourceforge.net/project/rkflashtool/rkflashtool-6.1/rkflashtool-6.1-src.tar.bz2"
	rkflashtool_cmd_src := exec.Command("curl", "-L", rkflashtool_src_url, "-o", "source.tar.gz")
	err = rkflashtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rkflashtool_cmd_direct := exec.Command("./binary")
	err = rkflashtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
