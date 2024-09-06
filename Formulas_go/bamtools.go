package main

// Bamtools - C++ API and command-line toolkit for BAM data
// Homepage: https://github.com/pezmaster31/bamtools

import (
	"fmt"
	
	"os/exec"
)

func installBamtools() {
	// Método 1: Descargar y extraer .tar.gz
	bamtools_tar_url := "https://github.com/pezmaster31/bamtools/archive/refs/tags/v2.5.2.tar.gz"
	bamtools_cmd_tar := exec.Command("curl", "-L", bamtools_tar_url, "-o", "package.tar.gz")
	err := bamtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bamtools_zip_url := "https://github.com/pezmaster31/bamtools/archive/refs/tags/v2.5.2.zip"
	bamtools_cmd_zip := exec.Command("curl", "-L", bamtools_zip_url, "-o", "package.zip")
	err = bamtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bamtools_bin_url := "https://github.com/pezmaster31/bamtools/archive/refs/tags/v2.5.2.bin"
	bamtools_cmd_bin := exec.Command("curl", "-L", bamtools_bin_url, "-o", "binary.bin")
	err = bamtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bamtools_src_url := "https://github.com/pezmaster31/bamtools/archive/refs/tags/v2.5.2.src.tar.gz"
	bamtools_cmd_src := exec.Command("curl", "-L", bamtools_src_url, "-o", "source.tar.gz")
	err = bamtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bamtools_cmd_direct := exec.Command("./binary")
	err = bamtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
}
