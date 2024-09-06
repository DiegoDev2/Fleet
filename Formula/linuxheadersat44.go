package main

// LinuxHeadersAT44 - Header files of the Linux kernel
// Homepage: https://kernel.org/

import (
	"fmt"
	
	"os/exec"
)

func installLinuxHeadersAT44() {
	// Método 1: Descargar y extraer .tar.gz
	linuxheadersat44_tar_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.4.302.tar.gz"
	linuxheadersat44_cmd_tar := exec.Command("curl", "-L", linuxheadersat44_tar_url, "-o", "package.tar.gz")
	err := linuxheadersat44_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linuxheadersat44_zip_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.4.302.zip"
	linuxheadersat44_cmd_zip := exec.Command("curl", "-L", linuxheadersat44_zip_url, "-o", "package.zip")
	err = linuxheadersat44_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linuxheadersat44_bin_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.4.302.bin"
	linuxheadersat44_cmd_bin := exec.Command("curl", "-L", linuxheadersat44_bin_url, "-o", "binary.bin")
	err = linuxheadersat44_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linuxheadersat44_src_url := "https://cdn.kernel.org/pub/linux/kernel/v4.x/linux-4.4.302.src.tar.gz"
	linuxheadersat44_cmd_src := exec.Command("curl", "-L", linuxheadersat44_src_url, "-o", "source.tar.gz")
	err = linuxheadersat44_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linuxheadersat44_cmd_direct := exec.Command("./binary")
	err = linuxheadersat44_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
