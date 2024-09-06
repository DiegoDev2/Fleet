package main

// Dpcmd - Linux software for DediProg SF100/SF600
// Homepage: https://github.com/DediProgSW/SF100Linux

import (
	"fmt"
	
	"os/exec"
)

func installDpcmd() {
	// Método 1: Descargar y extraer .tar.gz
	dpcmd_tar_url := "https://github.com/DediProgSW/SF100Linux/archive/refs/tags/V1.14.20.x.tar.gz"
	dpcmd_cmd_tar := exec.Command("curl", "-L", dpcmd_tar_url, "-o", "package.tar.gz")
	err := dpcmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dpcmd_zip_url := "https://github.com/DediProgSW/SF100Linux/archive/refs/tags/V1.14.20.x.zip"
	dpcmd_cmd_zip := exec.Command("curl", "-L", dpcmd_zip_url, "-o", "package.zip")
	err = dpcmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dpcmd_bin_url := "https://github.com/DediProgSW/SF100Linux/archive/refs/tags/V1.14.20.x.bin"
	dpcmd_cmd_bin := exec.Command("curl", "-L", dpcmd_bin_url, "-o", "binary.bin")
	err = dpcmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dpcmd_src_url := "https://github.com/DediProgSW/SF100Linux/archive/refs/tags/V1.14.20.x.src.tar.gz"
	dpcmd_cmd_src := exec.Command("curl", "-L", dpcmd_src_url, "-o", "source.tar.gz")
	err = dpcmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dpcmd_cmd_direct := exec.Command("./binary")
	err = dpcmd_cmd_direct.Run()
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
