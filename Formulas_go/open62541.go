package main

// Open62541 - Open source implementation of OPC UA
// Homepage: https://open62541.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpen62541() {
	// Método 1: Descargar y extraer .tar.gz
	open62541_tar_url := "https://github.com/open62541/open62541/archive/refs/tags/v1.4.4.tar.gz"
	open62541_cmd_tar := exec.Command("curl", "-L", open62541_tar_url, "-o", "package.tar.gz")
	err := open62541_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	open62541_zip_url := "https://github.com/open62541/open62541/archive/refs/tags/v1.4.4.zip"
	open62541_cmd_zip := exec.Command("curl", "-L", open62541_zip_url, "-o", "package.zip")
	err = open62541_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	open62541_bin_url := "https://github.com/open62541/open62541/archive/refs/tags/v1.4.4.bin"
	open62541_cmd_bin := exec.Command("curl", "-L", open62541_bin_url, "-o", "binary.bin")
	err = open62541_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	open62541_src_url := "https://github.com/open62541/open62541/archive/refs/tags/v1.4.4.src.tar.gz"
	open62541_cmd_src := exec.Command("curl", "-L", open62541_src_url, "-o", "source.tar.gz")
	err = open62541_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	open62541_cmd_direct := exec.Command("./binary")
	err = open62541_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
