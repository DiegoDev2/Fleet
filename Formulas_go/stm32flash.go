package main

// Stm32flash - Open source flash program for STM32 using the ST serial bootloader
// Homepage: https://sourceforge.net/projects/stm32flash/

import (
	"fmt"
	
	"os/exec"
)

func installStm32flash() {
	// Método 1: Descargar y extraer .tar.gz
	stm32flash_tar_url := "https://downloads.sourceforge.net/project/stm32flash/stm32flash-0.7.tar.gz"
	stm32flash_cmd_tar := exec.Command("curl", "-L", stm32flash_tar_url, "-o", "package.tar.gz")
	err := stm32flash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stm32flash_zip_url := "https://downloads.sourceforge.net/project/stm32flash/stm32flash-0.7.zip"
	stm32flash_cmd_zip := exec.Command("curl", "-L", stm32flash_zip_url, "-o", "package.zip")
	err = stm32flash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stm32flash_bin_url := "https://downloads.sourceforge.net/project/stm32flash/stm32flash-0.7.bin"
	stm32flash_cmd_bin := exec.Command("curl", "-L", stm32flash_bin_url, "-o", "binary.bin")
	err = stm32flash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stm32flash_src_url := "https://downloads.sourceforge.net/project/stm32flash/stm32flash-0.7.src.tar.gz"
	stm32flash_cmd_src := exec.Command("curl", "-L", stm32flash_src_url, "-o", "source.tar.gz")
	err = stm32flash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stm32flash_cmd_direct := exec.Command("./binary")
	err = stm32flash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
