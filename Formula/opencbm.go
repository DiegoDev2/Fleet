package main

// Opencbm - Provides access to various floppy drive formats
// Homepage: https://spiro.trikaliotis.net/opencbm

import (
	"fmt"
	
	"os/exec"
)

func installOpencbm() {
	// Método 1: Descargar y extraer .tar.gz
	opencbm_tar_url := "https://github.com/OpenCBM/OpenCBM/archive/refs/tags/v0.4.99.104.tar.gz"
	opencbm_cmd_tar := exec.Command("curl", "-L", opencbm_tar_url, "-o", "package.tar.gz")
	err := opencbm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencbm_zip_url := "https://github.com/OpenCBM/OpenCBM/archive/refs/tags/v0.4.99.104.zip"
	opencbm_cmd_zip := exec.Command("curl", "-L", opencbm_zip_url, "-o", "package.zip")
	err = opencbm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencbm_bin_url := "https://github.com/OpenCBM/OpenCBM/archive/refs/tags/v0.4.99.104.bin"
	opencbm_cmd_bin := exec.Command("curl", "-L", opencbm_bin_url, "-o", "binary.bin")
	err = opencbm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencbm_src_url := "https://github.com/OpenCBM/OpenCBM/archive/refs/tags/v0.4.99.104.src.tar.gz"
	opencbm_cmd_src := exec.Command("curl", "-L", opencbm_src_url, "-o", "source.tar.gz")
	err = opencbm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencbm_cmd_direct := exec.Command("./binary")
	err = opencbm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cc65")
	exec.Command("latte", "install", "cc65").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
}
