package main

// Cmocka - Unit testing framework for C
// Homepage: https://cmocka.org/

import (
	"fmt"
	
	"os/exec"
)

func installCmocka() {
	// Método 1: Descargar y extraer .tar.gz
	cmocka_tar_url := "https://cmocka.org/files/1.1/cmocka-1.1.7.tar.xz"
	cmocka_cmd_tar := exec.Command("curl", "-L", cmocka_tar_url, "-o", "package.tar.gz")
	err := cmocka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmocka_zip_url := "https://cmocka.org/files/1.1/cmocka-1.1.7.tar.xz"
	cmocka_cmd_zip := exec.Command("curl", "-L", cmocka_zip_url, "-o", "package.zip")
	err = cmocka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmocka_bin_url := "https://cmocka.org/files/1.1/cmocka-1.1.7.tar.xz"
	cmocka_cmd_bin := exec.Command("curl", "-L", cmocka_bin_url, "-o", "binary.bin")
	err = cmocka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmocka_src_url := "https://cmocka.org/files/1.1/cmocka-1.1.7.tar.xz"
	cmocka_cmd_src := exec.Command("curl", "-L", cmocka_src_url, "-o", "source.tar.gz")
	err = cmocka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmocka_cmd_direct := exec.Command("./binary")
	err = cmocka_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
