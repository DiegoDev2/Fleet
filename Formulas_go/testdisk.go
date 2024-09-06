package main

// Testdisk - Powerful free data recovery utility
// Homepage: https://www.cgsecurity.org/wiki/TestDisk

import (
	"fmt"
	
	"os/exec"
)

func installTestdisk() {
	// Método 1: Descargar y extraer .tar.gz
	testdisk_tar_url := "https://www.cgsecurity.org/testdisk-7.2.tar.bz2"
	testdisk_cmd_tar := exec.Command("curl", "-L", testdisk_tar_url, "-o", "package.tar.gz")
	err := testdisk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	testdisk_zip_url := "https://www.cgsecurity.org/testdisk-7.2.tar.bz2"
	testdisk_cmd_zip := exec.Command("curl", "-L", testdisk_zip_url, "-o", "package.zip")
	err = testdisk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	testdisk_bin_url := "https://www.cgsecurity.org/testdisk-7.2.tar.bz2"
	testdisk_cmd_bin := exec.Command("curl", "-L", testdisk_bin_url, "-o", "binary.bin")
	err = testdisk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	testdisk_src_url := "https://www.cgsecurity.org/testdisk-7.2.tar.bz2"
	testdisk_cmd_src := exec.Command("curl", "-L", testdisk_src_url, "-o", "source.tar.gz")
	err = testdisk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	testdisk_cmd_direct := exec.Command("./binary")
	err = testdisk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
