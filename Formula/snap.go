package main

// Snap - Tool to work with .snap files
// Homepage: https://snapcraft.io/

import (
	"fmt"
	
	"os/exec"
)

func installSnap() {
	// Método 1: Descargar y extraer .tar.gz
	snap_tar_url := "https://github.com/snapcore/snapd/releases/download/2.65.1/snapd_2.65.1.vendor.tar.xz"
	snap_cmd_tar := exec.Command("curl", "-L", snap_tar_url, "-o", "package.tar.gz")
	err := snap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snap_zip_url := "https://github.com/snapcore/snapd/releases/download/2.65.1/snapd_2.65.1.vendor.tar.xz"
	snap_cmd_zip := exec.Command("curl", "-L", snap_zip_url, "-o", "package.zip")
	err = snap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snap_bin_url := "https://github.com/snapcore/snapd/releases/download/2.65.1/snapd_2.65.1.vendor.tar.xz"
	snap_cmd_bin := exec.Command("curl", "-L", snap_bin_url, "-o", "binary.bin")
	err = snap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snap_src_url := "https://github.com/snapcore/snapd/releases/download/2.65.1/snapd_2.65.1.vendor.tar.xz"
	snap_cmd_src := exec.Command("curl", "-L", snap_src_url, "-o", "source.tar.gz")
	err = snap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snap_cmd_direct := exec.Command("./binary")
	err = snap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: squashfs")
	exec.Command("latte", "install", "squashfs").Run()
}
