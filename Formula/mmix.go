package main

// Mmix - 64-bit RISC architecture designed by Donald Knuth
// Homepage: https://mmix.cs.hm.edu/

import (
	"fmt"
	
	"os/exec"
)

func installMmix() {
	// Método 1: Descargar y extraer .tar.gz
	mmix_tar_url := "https://mmix.cs.hm.edu/src/mmix-20160804.tgz"
	mmix_cmd_tar := exec.Command("curl", "-L", mmix_tar_url, "-o", "package.tar.gz")
	err := mmix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmix_zip_url := "https://mmix.cs.hm.edu/src/mmix-20160804.tgz"
	mmix_cmd_zip := exec.Command("curl", "-L", mmix_zip_url, "-o", "package.zip")
	err = mmix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmix_bin_url := "https://mmix.cs.hm.edu/src/mmix-20160804.tgz"
	mmix_cmd_bin := exec.Command("curl", "-L", mmix_bin_url, "-o", "binary.bin")
	err = mmix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmix_src_url := "https://mmix.cs.hm.edu/src/mmix-20160804.tgz"
	mmix_cmd_src := exec.Command("curl", "-L", mmix_src_url, "-o", "source.tar.gz")
	err = mmix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmix_cmd_direct := exec.Command("./binary")
	err = mmix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cweb")
	exec.Command("latte", "install", "cweb").Run()
}
