package main

// Zzz - Command-line tool to put Macs to sleep
// Homepage: https://github.com/Orc/Zzz

import (
	"fmt"
	
	"os/exec"
)

func installZzz() {
	// Método 1: Descargar y extraer .tar.gz
	zzz_tar_url := "https://github.com/Orc/Zzz/archive/refs/tags/v1.tar.gz"
	zzz_cmd_tar := exec.Command("curl", "-L", zzz_tar_url, "-o", "package.tar.gz")
	err := zzz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zzz_zip_url := "https://github.com/Orc/Zzz/archive/refs/tags/v1.zip"
	zzz_cmd_zip := exec.Command("curl", "-L", zzz_zip_url, "-o", "package.zip")
	err = zzz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zzz_bin_url := "https://github.com/Orc/Zzz/archive/refs/tags/v1.bin"
	zzz_cmd_bin := exec.Command("curl", "-L", zzz_bin_url, "-o", "binary.bin")
	err = zzz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zzz_src_url := "https://github.com/Orc/Zzz/archive/refs/tags/v1.src.tar.gz"
	zzz_cmd_src := exec.Command("curl", "-L", zzz_src_url, "-o", "source.tar.gz")
	err = zzz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zzz_cmd_direct := exec.Command("./binary")
	err = zzz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
