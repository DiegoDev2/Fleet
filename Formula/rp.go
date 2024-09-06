package main

// Rp - Tool to find ROP sequences in PE/Elf/Mach-O x86/x64 binaries
// Homepage: https://github.com/0vercl0k/rp

import (
	"fmt"
	
	"os/exec"
)

func installRp() {
	// Método 1: Descargar y extraer .tar.gz
	rp_tar_url := "https://github.com/0vercl0k/rp/archive/refs/tags/v2.1.3.tar.gz"
	rp_cmd_tar := exec.Command("curl", "-L", rp_tar_url, "-o", "package.tar.gz")
	err := rp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rp_zip_url := "https://github.com/0vercl0k/rp/archive/refs/tags/v2.1.3.zip"
	rp_cmd_zip := exec.Command("curl", "-L", rp_zip_url, "-o", "package.zip")
	err = rp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rp_bin_url := "https://github.com/0vercl0k/rp/archive/refs/tags/v2.1.3.bin"
	rp_cmd_bin := exec.Command("curl", "-L", rp_bin_url, "-o", "binary.bin")
	err = rp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rp_src_url := "https://github.com/0vercl0k/rp/archive/refs/tags/v2.1.3.src.tar.gz"
	rp_cmd_src := exec.Command("curl", "-L", rp_src_url, "-o", "source.tar.gz")
	err = rp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rp_cmd_direct := exec.Command("./binary")
	err = rp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
