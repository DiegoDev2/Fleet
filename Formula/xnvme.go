package main

// Xnvme - Cross-platform libraries and tools for efficient I/O and low-level control
// Homepage: https://xnvme.io/

import (
	"fmt"
	
	"os/exec"
)

func installXnvme() {
	// Método 1: Descargar y extraer .tar.gz
	xnvme_tar_url := "https://github.com/OpenMPDK/xNVMe/releases/download/v0.7.4/xnvme-src-0.7.4.tar.gz"
	xnvme_cmd_tar := exec.Command("curl", "-L", xnvme_tar_url, "-o", "package.tar.gz")
	err := xnvme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xnvme_zip_url := "https://github.com/OpenMPDK/xNVMe/releases/download/v0.7.4/xnvme-src-0.7.4.zip"
	xnvme_cmd_zip := exec.Command("curl", "-L", xnvme_zip_url, "-o", "package.zip")
	err = xnvme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xnvme_bin_url := "https://github.com/OpenMPDK/xNVMe/releases/download/v0.7.4/xnvme-src-0.7.4.bin"
	xnvme_cmd_bin := exec.Command("curl", "-L", xnvme_bin_url, "-o", "binary.bin")
	err = xnvme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xnvme_src_url := "https://github.com/OpenMPDK/xNVMe/releases/download/v0.7.4/xnvme-src-0.7.4.src.tar.gz"
	xnvme_cmd_src := exec.Command("curl", "-L", xnvme_src_url, "-o", "source.tar.gz")
	err = xnvme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xnvme_cmd_direct := exec.Command("./binary")
	err = xnvme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
