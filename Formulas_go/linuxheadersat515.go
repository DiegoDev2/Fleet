package main

// LinuxHeadersAT515 - Header files of the Linux kernel
// Homepage: https://kernel.org/

import (
	"fmt"
	
	"os/exec"
)

func installLinuxHeadersAT515() {
	// Método 1: Descargar y extraer .tar.gz
	linuxheadersat515_tar_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.166.tar.gz"
	linuxheadersat515_cmd_tar := exec.Command("curl", "-L", linuxheadersat515_tar_url, "-o", "package.tar.gz")
	err := linuxheadersat515_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	linuxheadersat515_zip_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.166.zip"
	linuxheadersat515_cmd_zip := exec.Command("curl", "-L", linuxheadersat515_zip_url, "-o", "package.zip")
	err = linuxheadersat515_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	linuxheadersat515_bin_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.166.bin"
	linuxheadersat515_cmd_bin := exec.Command("curl", "-L", linuxheadersat515_bin_url, "-o", "binary.bin")
	err = linuxheadersat515_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	linuxheadersat515_src_url := "https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.166.src.tar.gz"
	linuxheadersat515_cmd_src := exec.Command("curl", "-L", linuxheadersat515_src_url, "-o", "source.tar.gz")
	err = linuxheadersat515_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	linuxheadersat515_cmd_direct := exec.Command("./binary")
	err = linuxheadersat515_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
