package main

// LinuxHeadersAT516 - Header files of the Linux kernel
// Homepage: https://kernel.org/

import (
	"fmt"
	
	"os/exec"
)

func installLinuxHeadersAT516() {
	// Método 1: Descargar y extraer .tar.gz
	linuxheadersat516_tar_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.16.20.tar.gz"
	linuxheadersat516_cmd_tar := exec.Command("curl", "-L", linuxheadersat516_tar_url, "-o", "package.tar.gz")
	err := linuxheadersat516_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linuxheadersat516_zip_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.16.20.zip"
	linuxheadersat516_cmd_zip := exec.Command("curl", "-L", linuxheadersat516_zip_url, "-o", "package.zip")
	err = linuxheadersat516_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linuxheadersat516_bin_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.16.20.bin"
	linuxheadersat516_cmd_bin := exec.Command("curl", "-L", linuxheadersat516_bin_url, "-o", "binary.bin")
	err = linuxheadersat516_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linuxheadersat516_src_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.16.20.src.tar.gz"
	linuxheadersat516_cmd_src := exec.Command("curl", "-L", linuxheadersat516_src_url, "-o", "source.tar.gz")
	err = linuxheadersat516_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linuxheadersat516_cmd_direct := exec.Command("./binary")
	err = linuxheadersat516_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rsync")
exec.Command("latte", "install", "rsync")
}
