package main

// Dtools - D programming language tools
// Homepage: https://dlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installDtools() {
	// Método 1: Descargar y extraer .tar.gz
	dtools_tar_url := "https://github.com/dlang/tools/archive/refs/tags/v2.109.1.tar.gz"
	dtools_cmd_tar := exec.Command("curl", "-L", dtools_tar_url, "-o", "package.tar.gz")
	err := dtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dtools_zip_url := "https://github.com/dlang/tools/archive/refs/tags/v2.109.1.zip"
	dtools_cmd_zip := exec.Command("curl", "-L", dtools_zip_url, "-o", "package.zip")
	err = dtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dtools_bin_url := "https://github.com/dlang/tools/archive/refs/tags/v2.109.1.bin"
	dtools_cmd_bin := exec.Command("curl", "-L", dtools_bin_url, "-o", "binary.bin")
	err = dtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dtools_src_url := "https://github.com/dlang/tools/archive/refs/tags/v2.109.1.src.tar.gz"
	dtools_cmd_src := exec.Command("curl", "-L", dtools_src_url, "-o", "source.tar.gz")
	err = dtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dtools_cmd_direct := exec.Command("./binary")
	err = dtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dub")
	exec.Command("latte", "install", "dub").Run()
	fmt.Println("Instalando dependencia: ldc")
	exec.Command("latte", "install", "ldc").Run()
}
