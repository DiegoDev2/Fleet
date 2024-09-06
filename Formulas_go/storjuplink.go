package main

// StorjUplink - Uplink CLI for the Storj network
// Homepage: https://storj.io

import (
	"fmt"
	
	"os/exec"
)

func installStorjUplink() {
	// Método 1: Descargar y extraer .tar.gz
	storjuplink_tar_url := "https://github.com/storj/storj/archive/refs/tags/v1.110.3.tar.gz"
	storjuplink_cmd_tar := exec.Command("curl", "-L", storjuplink_tar_url, "-o", "package.tar.gz")
	err := storjuplink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	storjuplink_zip_url := "https://github.com/storj/storj/archive/refs/tags/v1.110.3.zip"
	storjuplink_cmd_zip := exec.Command("curl", "-L", storjuplink_zip_url, "-o", "package.zip")
	err = storjuplink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	storjuplink_bin_url := "https://github.com/storj/storj/archive/refs/tags/v1.110.3.bin"
	storjuplink_cmd_bin := exec.Command("curl", "-L", storjuplink_bin_url, "-o", "binary.bin")
	err = storjuplink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	storjuplink_src_url := "https://github.com/storj/storj/archive/refs/tags/v1.110.3.src.tar.gz"
	storjuplink_cmd_src := exec.Command("curl", "-L", storjuplink_src_url, "-o", "source.tar.gz")
	err = storjuplink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	storjuplink_cmd_direct := exec.Command("./binary")
	err = storjuplink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
