package main

// Deheader - Analyze C/C++ files for unnecessary headers
// Homepage: http://www.catb.org/~esr/deheader/

import (
	"fmt"
	
	"os/exec"
)

func installDeheader() {
	// Método 1: Descargar y extraer .tar.gz
	deheader_tar_url := "http://www.catb.org/~esr/deheader/deheader-1.10.tar.gz"
	deheader_cmd_tar := exec.Command("curl", "-L", deheader_tar_url, "-o", "package.tar.gz")
	err := deheader_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	deheader_zip_url := "http://www.catb.org/~esr/deheader/deheader-1.10.zip"
	deheader_cmd_zip := exec.Command("curl", "-L", deheader_zip_url, "-o", "package.zip")
	err = deheader_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	deheader_bin_url := "http://www.catb.org/~esr/deheader/deheader-1.10.bin"
	deheader_cmd_bin := exec.Command("curl", "-L", deheader_bin_url, "-o", "binary.bin")
	err = deheader_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	deheader_src_url := "http://www.catb.org/~esr/deheader/deheader-1.10.src.tar.gz"
	deheader_cmd_src := exec.Command("curl", "-L", deheader_src_url, "-o", "source.tar.gz")
	err = deheader_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	deheader_cmd_direct := exec.Command("./binary")
	err = deheader_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
