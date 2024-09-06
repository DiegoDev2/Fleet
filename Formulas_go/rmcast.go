package main

// Rmcast - IP Multicast library
// Homepage: https://www.land.ufrj.br/tools/rmcast/rmcast.html

import (
	"fmt"
	
	"os/exec"
)

func installRmcast() {
	// Método 1: Descargar y extraer .tar.gz
	rmcast_tar_url := "https://www.land.ufrj.br/tools/rmcast/download/rmcast-2.0.0.tar.gz"
	rmcast_cmd_tar := exec.Command("curl", "-L", rmcast_tar_url, "-o", "package.tar.gz")
	err := rmcast_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rmcast_zip_url := "https://www.land.ufrj.br/tools/rmcast/download/rmcast-2.0.0.zip"
	rmcast_cmd_zip := exec.Command("curl", "-L", rmcast_zip_url, "-o", "package.zip")
	err = rmcast_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rmcast_bin_url := "https://www.land.ufrj.br/tools/rmcast/download/rmcast-2.0.0.bin"
	rmcast_cmd_bin := exec.Command("curl", "-L", rmcast_bin_url, "-o", "binary.bin")
	err = rmcast_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rmcast_src_url := "https://www.land.ufrj.br/tools/rmcast/download/rmcast-2.0.0.src.tar.gz"
	rmcast_cmd_src := exec.Command("curl", "-L", rmcast_src_url, "-o", "source.tar.gz")
	err = rmcast_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rmcast_cmd_direct := exec.Command("./binary")
	err = rmcast_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
