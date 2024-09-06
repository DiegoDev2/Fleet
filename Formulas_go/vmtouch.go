package main

// Vmtouch - Portable file system cache diagnostics and control
// Homepage: https://hoytech.com/vmtouch/

import (
	"fmt"
	
	"os/exec"
)

func installVmtouch() {
	// Método 1: Descargar y extraer .tar.gz
	vmtouch_tar_url := "https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.tar.gz"
	vmtouch_cmd_tar := exec.Command("curl", "-L", vmtouch_tar_url, "-o", "package.tar.gz")
	err := vmtouch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vmtouch_zip_url := "https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.zip"
	vmtouch_cmd_zip := exec.Command("curl", "-L", vmtouch_zip_url, "-o", "package.zip")
	err = vmtouch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vmtouch_bin_url := "https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.bin"
	vmtouch_cmd_bin := exec.Command("curl", "-L", vmtouch_bin_url, "-o", "binary.bin")
	err = vmtouch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vmtouch_src_url := "https://github.com/hoytech/vmtouch/archive/refs/tags/v1.3.1.src.tar.gz"
	vmtouch_cmd_src := exec.Command("curl", "-L", vmtouch_src_url, "-o", "source.tar.gz")
	err = vmtouch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vmtouch_cmd_direct := exec.Command("./binary")
	err = vmtouch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
