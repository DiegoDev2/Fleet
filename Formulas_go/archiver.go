package main

// Archiver - Cross-platform, multi-format archive utility
// Homepage: https://github.com/mholt/archiver

import (
	"fmt"
	
	"os/exec"
)

func installArchiver() {
	// Método 1: Descargar y extraer .tar.gz
	archiver_tar_url := "https://github.com/mholt/archiver/archive/refs/tags/v3.5.1.tar.gz"
	archiver_cmd_tar := exec.Command("curl", "-L", archiver_tar_url, "-o", "package.tar.gz")
	err := archiver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	archiver_zip_url := "https://github.com/mholt/archiver/archive/refs/tags/v3.5.1.zip"
	archiver_cmd_zip := exec.Command("curl", "-L", archiver_zip_url, "-o", "package.zip")
	err = archiver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	archiver_bin_url := "https://github.com/mholt/archiver/archive/refs/tags/v3.5.1.bin"
	archiver_cmd_bin := exec.Command("curl", "-L", archiver_bin_url, "-o", "binary.bin")
	err = archiver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	archiver_src_url := "https://github.com/mholt/archiver/archive/refs/tags/v3.5.1.src.tar.gz"
	archiver_cmd_src := exec.Command("curl", "-L", archiver_src_url, "-o", "source.tar.gz")
	err = archiver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	archiver_cmd_direct := exec.Command("./binary")
	err = archiver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
