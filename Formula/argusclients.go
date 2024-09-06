package main

// ArgusClients - Audit Record Generation and Utilization System clients
// Homepage: https://openargus.org

import (
	"fmt"
	
	"os/exec"
)

func installArgusClients() {
	// Método 1: Descargar y extraer .tar.gz
	argusclients_tar_url := "https://github.com/openargus/clients/archive/refs/tags/v5.0.0.tar.gz"
	argusclients_cmd_tar := exec.Command("curl", "-L", argusclients_tar_url, "-o", "package.tar.gz")
	err := argusclients_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argusclients_zip_url := "https://github.com/openargus/clients/archive/refs/tags/v5.0.0.zip"
	argusclients_cmd_zip := exec.Command("curl", "-L", argusclients_zip_url, "-o", "package.zip")
	err = argusclients_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argusclients_bin_url := "https://github.com/openargus/clients/archive/refs/tags/v5.0.0.bin"
	argusclients_cmd_bin := exec.Command("curl", "-L", argusclients_bin_url, "-o", "binary.bin")
	err = argusclients_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argusclients_src_url := "https://github.com/openargus/clients/archive/refs/tags/v5.0.0.src.tar.gz"
	argusclients_cmd_src := exec.Command("curl", "-L", argusclients_src_url, "-o", "source.tar.gz")
	err = argusclients_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argusclients_cmd_direct := exec.Command("./binary")
	err = argusclients_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: rrdtool")
	exec.Command("latte", "install", "rrdtool").Run()
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
