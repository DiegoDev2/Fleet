package main

// Cdecl - Turn English phrases to C or C++ declarations
// Homepage: https://cdecl.org/

import (
	"fmt"
	
	"os/exec"
)

func installCdecl() {
	// Método 1: Descargar y extraer .tar.gz
	cdecl_tar_url := "https://cdecl.org/files/cdecl-blocks-2.5.tar.gz"
	cdecl_cmd_tar := exec.Command("curl", "-L", cdecl_tar_url, "-o", "package.tar.gz")
	err := cdecl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdecl_zip_url := "https://cdecl.org/files/cdecl-blocks-2.5.zip"
	cdecl_cmd_zip := exec.Command("curl", "-L", cdecl_zip_url, "-o", "package.zip")
	err = cdecl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdecl_bin_url := "https://cdecl.org/files/cdecl-blocks-2.5.bin"
	cdecl_cmd_bin := exec.Command("curl", "-L", cdecl_bin_url, "-o", "binary.bin")
	err = cdecl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdecl_src_url := "https://cdecl.org/files/cdecl-blocks-2.5.src.tar.gz"
	cdecl_cmd_src := exec.Command("curl", "-L", cdecl_src_url, "-o", "source.tar.gz")
	err = cdecl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdecl_cmd_direct := exec.Command("./binary")
	err = cdecl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
