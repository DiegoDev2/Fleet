package main

// Mdds - Multi-dimensional data structure and indexing algorithm
// Homepage: https://gitlab.com/mdds/mdds

import (
	"fmt"
	
	"os/exec"
)

func installMdds() {
	// Método 1: Descargar y extraer .tar.gz
	mdds_tar_url := "https://kohei.us/files/mdds/src/mdds-2.1.1.tar.bz2"
	mdds_cmd_tar := exec.Command("curl", "-L", mdds_tar_url, "-o", "package.tar.gz")
	err := mdds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdds_zip_url := "https://kohei.us/files/mdds/src/mdds-2.1.1.tar.bz2"
	mdds_cmd_zip := exec.Command("curl", "-L", mdds_zip_url, "-o", "package.zip")
	err = mdds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdds_bin_url := "https://kohei.us/files/mdds/src/mdds-2.1.1.tar.bz2"
	mdds_cmd_bin := exec.Command("curl", "-L", mdds_bin_url, "-o", "binary.bin")
	err = mdds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdds_src_url := "https://kohei.us/files/mdds/src/mdds-2.1.1.tar.bz2"
	mdds_cmd_src := exec.Command("curl", "-L", mdds_src_url, "-o", "source.tar.gz")
	err = mdds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdds_cmd_direct := exec.Command("./binary")
	err = mdds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
