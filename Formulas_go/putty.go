package main

// Putty - Implementation of Telnet and SSH
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/putty/

import (
	"fmt"
	
	"os/exec"
)

func installPutty() {
	// Método 1: Descargar y extraer .tar.gz
	putty_tar_url := "https://the.earth.li/~sgtatham/putty/0.81/putty-0.81.tar.gz"
	putty_cmd_tar := exec.Command("curl", "-L", putty_tar_url, "-o", "package.tar.gz")
	err := putty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	putty_zip_url := "https://the.earth.li/~sgtatham/putty/0.81/putty-0.81.zip"
	putty_cmd_zip := exec.Command("curl", "-L", putty_zip_url, "-o", "package.zip")
	err = putty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	putty_bin_url := "https://the.earth.li/~sgtatham/putty/0.81/putty-0.81.bin"
	putty_cmd_bin := exec.Command("curl", "-L", putty_bin_url, "-o", "binary.bin")
	err = putty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	putty_src_url := "https://the.earth.li/~sgtatham/putty/0.81/putty-0.81.src.tar.gz"
	putty_cmd_src := exec.Command("curl", "-L", putty_src_url, "-o", "source.tar.gz")
	err = putty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	putty_cmd_direct := exec.Command("./binary")
	err = putty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: halibut")
exec.Command("latte", "install", "halibut")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
